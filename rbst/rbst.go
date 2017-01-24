package rbst

import (
	"math/rand"
)

type Node struct {
	Value int
	Size  int
	Left  *Node
	Right *Node
}

func Find(n *Node, k int) *Node {
	if n == nil {
		return nil
	}
	if n.Size <= k {
		return nil
	}
	rank := 0
	if n.Left != nil {
		rank += n.Left.Size
	}
	if k == rank {
		return n
	}
	if k < rank {
		return Find(n.Left, k)
	}
	return Find(n.Right, k-rank-1)
}

func Merge(a *Node, b *Node) *Node {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if a.Size <= rand.Intn(a.Size+b.Size) {
		// a is new root.
		a.Size += b.Size
		if a.Value < b.Value {
			a.Right = Merge(a.Right, b)
		} else {
			a.Left = Merge(a.Left, b)
		}
		return a
	} else {
		// b is new root.
		b.Size += a.Size
		if b.Value < a.Value {
			b.Right = Merge(b.Right, a)
		} else {
			b.Left = Merge(b.Left, a)
		}
		return b
	}
}

func NewNode(v int) *Node {
	return &Node{v, 1, nil, nil}
}
