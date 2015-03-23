package graph

import (
	"container/list"
)

//Next iteration should use http://play.golang.org/p/HXjmmRXl6k

// Queue type used in search
type Queue list.List

//NewQueue initializes a new queue object
func NewQueue() *Queue {
	q := list.New()
	return (*Queue)(q)
}

// Push adds a node to the queue.
func (q *Queue) Push(s string) {
	l := (*list.List)(q)
	l.PushBack(s)
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() (string, bool) {
	l := (*list.List)(q)
	if l.Len() == 0 {
		return "", false
	}
	s := l.Remove(l.Front()).(string)
	return s, true
}
