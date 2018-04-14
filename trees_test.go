package isa497

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	mockTree := NewTree(1, nil)
	parentTree = *mockTree
	parentTree.left = NewTree(2, nil) 
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{"base", args{mockTree}, mockTree},
		{"baseErr", args{0, nil}, &Tree{0, nil, nil, nil}},
		{"detect child", }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree(tt.args.v, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inorderTreeWalk(tt.args.t)
		})
	}
}
