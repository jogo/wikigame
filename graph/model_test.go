package graph

import (
	"os"
	"testing"
)

var filename = "temp.bolt"
var value = []string{"a", "b"}

func setup() *Graph {
	return NewGraph(filename)
}
func teardown(graph *Graph) {
	graph.Close()
	os.Remove(filename)
}

func TestToFromBytes(t *testing.T) {
	graph := setup()
	defer teardown(graph)
	actual := graph.fromBytes(graph.toBytes(value))
	if actual[0] != "a" || actual[1] != actual[1] {
		t.Errorf("something went wrong in converting to and from a byte array")
	}
}

func TestNewGraph(t *testing.T) {
	graph := setup()
	defer teardown(graph)
	graph.Put("key", value)
	actual := graph.Get("key")
	if actual[0] != "a" || actual[1] != actual[1] {
		t.Errorf("Expected %s, but got %s", value, actual)
	}
}

func TestFlush(t *testing.T) {
	graph := setup()
	defer teardown(graph)
	graph.Put("key", value)
	if len(graph.buffer) != 1 {
		t.Errorf("expected data in the buffer")
	}
	graph.Flush()
	if len(graph.buffer) != 0 {
		t.Errorf("expected the buffer to be flushed")
	}

}
