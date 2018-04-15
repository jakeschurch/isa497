// Copyright (c) 2017 Jake Schurch
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package isa497

import (
	"fmt"
)

// Tree represents a Tree data structure.
type Tree struct {
	key    int
	root   *Tree
	parent *Tree
	left   *Tree
	right  *Tree
}

// Key returns the key-value of a tree
func (t *Tree) Key() int {
	return t.key
}

// NewTree returns a new Tree struct.
func NewTree(k int) (t *Tree) {
	t = &Tree{
		key:    k,
		parent: nil,
		left:   nil,
		right:  nil,
	}
	t.root = t
	return
}

func inorderTreeWalk(t *Tree) {
	if t != nil {
		inorderTreeWalk(t.left)
		fmt.Println(t.key)
		inorderTreeWalk(t.right)
	}
}

func recTreeSearch(x *Tree, key int) *Tree {
	if x == nil || x.key == key {
		return x
	}
	if key < x.key {
		return recTreeSearch(x.left, key)
	}
	return recTreeSearch(x.right, key)
}

func iterTreeSearch(x *Tree, key int) *Tree {
	for x != nil && key != x.key {
		if key < x.key {
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

func treeInsert(t, z *Tree) {
	var x, y *Tree

	for x = t.root; x != nil; {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	z.root = y.root
	switch {
	case y == nil:
		t.root = z
	case z.key < y.key:
		y.left = z
	default:
		y.right = z
	}
}

func transplant(t, u, v *Tree) {
	switch {
	case u.parent == nil:
		t.root = v
	case u == u.parent.left:
		u.parent.left = v
	default:
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func treeDelete(t, z *Tree) {
	var y *Tree

	switch {
	case z.left == nil:
		transplant(t, z, z.right)
	case z.right == nil:
		transplant(t, z, z.left)
	default:
		if y = treeMin(z.right); y.parent != z {
			transplant(t, y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		transplant(t, z, y)
		y.left = z.left
		y.left.parent = y
	}
}
