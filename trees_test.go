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
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	parent := &Tree{0, nil, nil, nil, nil}
	parent.root = parent

	type args struct {
		v int
		p *Tree
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{"base case", args{0, nil}, parent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Key(t *testing.T) {
	var parent = NewTree(10)

	tests := []struct {
		name string
		t    *Tree
		want int
	}{
		{"base case", parent, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Key(); got != tt.want {
				t.Errorf("Tree.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recTreeSearch(t *testing.T) {
	var parent, child *Tree
	parent = NewTree(2)
	child = NewTree(20)
	treeInsert(parent, child)

	type args struct {
		x   *Tree
		key int
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{"base case", args{parent, 20}, child},
		{"not found", args{parent, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recTreeSearch(tt.args.x, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recTreeSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iterTreeSearch(t *testing.T) {
	var parent, child *Tree
	parent = NewTree(2)
	child = NewTree(20)
	treeInsert(parent, child)

	type args struct {
		x   *Tree
		key int
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{"base case", args{parent, 20}, child},
		{"not found", args{parent, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterTreeSearch(tt.args.x, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iterTreeSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeMin(t *testing.T) {
	var parent, child *Tree
	parent = NewTree(2)
	child = NewTree(20)
	treeInsert(parent, child)

	type args struct {
		x *Tree
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{"base case", args{parent}, parent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeMin(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeMax(t *testing.T) {
	var parent, child *Tree
	parent = NewTree(2)
	child = NewTree(20)
	treeInsert(parent, child)

	type args struct {
		x *Tree
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{"base case", args{parent}, child},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeMax(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeSuccessor(t *testing.T) {
	var parent, child *Tree

	parent = NewTree(2)
	child = NewTree(1)
	treeInsert(parent, child)

	type args struct {
		x *Tree
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{"base case", args{parent.left}, parent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeSuccessor(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeInsert(t *testing.T) {
	var parent, child *Tree

	parent = NewTree(2)
	child = NewTree(1)

	type args struct {
		t *Tree
		z *Tree
	}
	tests := []struct {
		name string
		args args
	}{
		{"base case", args{parent, child}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treeInsert(tt.args.t, tt.args.z)
		})
		if tt.name = "base case"; tt.args.t.left != tt.args.z {
			t.Errorf("treeInsert() failed test: %v", tt.name)
		}
	}
}

func Test_transplant(t *testing.T) {
	var parent, child, toInsert *Tree

	parent = NewTree(2)
	child = NewTree(0)
	treeInsert(parent, child)
	toInsert = NewTree(1)

	type args struct {
		t *Tree
		u *Tree
		v *Tree
	}
	tests := []struct {
		name string
		args args
	}{
		{"base case", args{parent, child, toInsert}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transplant(tt.args.t, tt.args.u, tt.args.v)
		})
		if tt.name == "base case" && tt.args.t.left != tt.args.v {
			t.Errorf("transplant() failed test: %v", tt.name)
		}
	}
}

func Test_treeDelete(t *testing.T) {
	var parent, child *Tree

	parent = NewTree(2)
	child = NewTree(0)
	treeInsert(parent, child)

	type args struct {
		t *Tree
		z *Tree
	}
	tests := []struct {
		name string
		args args
	}{
		{"base case", args{parent, child}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treeDelete(tt.args.t, tt.args.z)
		})
		if tt.name == "base case" && tt.args.t.left != nil {
			t.Errorf("treeDelete() failed test: %v", tt.name)
		}
	}
}

func Test_inorderTreeWalk(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
	}{
		{"base case", args{NewTree(2)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inorderTreeWalk(tt.args.t)
		})
	}
}
