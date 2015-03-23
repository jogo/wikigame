package graph

import (
	"testing"
)

func TestLinked(t *testing.T) {
	graph := setup()
	defer teardown(graph)
	linked := graph.Linked("t1", "t2")
	if linked != true {
		t.Errorf("t1 is linked to t2, but not Linked() failed")
	}
	linked = graph.Linked("t1", "t3")
	if linked != false {
		t.Errorf("t1 is linked to t3, but not Linked() failed")
	}

}

func TestLinkedCaseSensitivity(t *testing.T) {
	graph := setup()
	defer teardown(graph)
	linked := graph.Linked("T1", "t2")
	if linked != true {
		t.Errorf("t1 is linked to t3, but not Linked() failed")
	}
}

func TestPath(t *testing.T) {
	graph := setup()
	defer teardown(graph)
	linked := graph.Path("t1", "t3")
	if linked != "Path: t1 => t2 => t3" {
		t.Errorf("expected 'Path: t1 => t2 => t3', but got '%s'", linked)
	}
}
