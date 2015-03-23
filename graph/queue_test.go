package graph

import "testing"

func TestPushPop(t *testing.T) {
	q := NewQueue()
	values := []string{"a", "b", "c", "bigbadstring"}
	for _, x := range values {
		q.Push(x)
	}
	var out string
	var ok bool
	for i := range values {
		out, ok = q.Pop()
		if out != values[i] || !ok {
			t.Errorf("expected %s, but got %s", values[i], out)
		}
	}
	out, ok = q.Pop()
	if ok {
		t.Errorf("queue should be empty, but got %s", out)
	}
}
