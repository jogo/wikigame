package graph

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
)

/*
Graph object

Store map of page name: list of links on that page

use Put, Get to access pages map
*/
type Graph struct {
	//exported so gob can access this
	buffer     map[string][]string
	db         *bolt.DB
	batchsSize int
	bucket     []byte
}

func prepBolt(filename string) *bolt.DB {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

//NewGraph initializes a graph
func NewGraph(filename string) *Graph {
	g := Graph{
		buffer:     make(map[string][]string),
		db:         prepBolt(filename),
		batchsSize: 10000,
		bucket:     []byte("pages"),
	}

	// make sure bucket exists
	err := g.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(g.bucket)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	g.db.NoSync = true
	return &g
}

//Put saves a list of links on a given page
func (g *Graph) Put(page string, value []string) {
	g.buffer[page] = value
	if len(g.buffer) > g.batchsSize {
		g.Flush()
	}
}

//Get list of links on a given page
func (g *Graph) Get(page string) []string {
	g.Flush()
	var value []string
	g.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(g.bucket)
		raw := b.Get([]byte(page))
		value = g.fromBytes(raw)
		return nil
	})
	return value
}

func (g *Graph) toBytes(value []string) (bytes []byte) {
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (g *Graph) fromBytes(raw []byte) (value []string) {
	if len(raw) == 0 {
		// don't try unmarshalling if empty
		return
	}
	err := json.Unmarshal((raw), &value)
	if err != nil {
		log.Fatal(err)
	}
	return
}

//Close passes through to bolt.Close
func (g *Graph) Close() {
	g.db.Close()
}

//Flush makes sure buffer is empty and all data is stored in boltdb
func (g *Graph) Flush() {
	err := g.db.Update(func(tx *bolt.Tx) error {
		//var err error
		b := tx.Bucket(g.bucket)
		for key, value := range g.buffer {
			bytes := g.toBytes(value)
			err := b.Put([]byte(key), bytes)
			delete(g.buffer, key)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
