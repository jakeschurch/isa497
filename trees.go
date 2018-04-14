package isa497

import "fmt"

// Tree represents a Tree data structure.
type Tree struct {
	Val   int
	root  *Tree
	left  *Tree
	right *Tree
}

// NewTree returns a new Tree struct.
func NewTree(v int, p *Tree) *Tree {
	return &Tree{v, p, nil, nil}
}

func inorderTreeWalk(t *Tree) {
	if t != nil {
		inorderTreeWalk(t.left)
		fmt.Println(t.Val)
		inorderTreeWalk(t.right)
	}
}

func recTreeSearch(x *Tree, Val int) *Tree {
	if x == nil || x.Val == Val {
		return x
	}
	if Val < x.Val {
		return recTreeSearch(x.left, Val)
	}
	return recTreeSearch(x.right, Val)
}

func iterTreeSearch(x *Tree, Val int) *Tree {
	for x != nil && Val != x.Val {
		if Val < x.Val {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func treeMin(x *Tree) *Tree {
	for x.left != nil {
		x = x.left
	}
	return x
}

func treeMax(x *Tree) *Tree {
	for x.right != nil {
		x = x.right
	}
	return x
}

func treeSuccessor(x *Tree) *Tree {
	if x.right != nil {
		return treeMin(x.right)
	}
	y := x.root
	for y != nil && x == y.right {
		x = y
		y = y.root
	}
	return y
}
