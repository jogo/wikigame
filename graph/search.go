package graph

import (
	"github.com/boltdb/bolt"
	"strings"
)

//Linked returns true if link to dest on the source page
func (g *Graph) Linked(source, dest string) bool {
	source = strings.ToLower(source)
	dest = strings.ToLower(dest)
	for _, p := range g.Get(source) {
		if p == dest {
			return true
		}
	}
	return false
}

//Path search using A* algorithm, returns path as a human readable string
func (g *Graph) Path(source, dest string) string {
	source = strings.ToLower(source)
	dest = strings.ToLower(dest)
	closed := make(map[string]bool)
	openmap := make(map[string]bool)
	from := make(map[string]string)
	open := NewQueue()
	open.Push(source)
	var curr string
	var ok bool

	var value []string
	var path string

	//do search inside of a single boltdb transaction for performance
	g.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(g.bucket)

		for {
			curr, ok = open.Pop()
			if !ok {
				break
			}
			delete(openmap, curr)
			closed[curr] = true

			if curr == dest {
				path = g.printPath(from, dest)
				return nil
			}

			raw := b.Get([]byte(curr))
			value = g.fromBytes(raw)

			for _, p := range value {
				// if never seen
				if closed[p] == false && openmap[p] == false {
					from[p] = curr
					open.Push(p)
					openmap[p] = true
				}
			}

		}
		return nil

	})
	if path != "" {
		return path
	}
	return "No Path Found: " + source + "!=>" + dest
}

func (g *Graph) printPath(from map[string]string, curr string) string {
	path := curr
	for _, ok := from[curr]; ok; _, ok = from[curr] {
		curr = from[curr]
		path = curr + " => " + path
	}
	return "Path: " + path
}
