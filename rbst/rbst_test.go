package rbst

import (
	"testing"
)

func TestRBST(t *testing.T) {
	var root *Node
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		n := NewNode(v)
		root = Merge(root, n)
	}

	if root.Size != 5 {
		t.Fatalf("tree size is not 5, got=%d", root.Size)
	}

	for i, v := range values {
		found := Find(root, i)
		if found.Value != v {
			t.Fatalf("%dth number is not %d, got=%d", i, v, found.Value)
		}
	}
}
