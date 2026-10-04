package main

import "testing"

func TestSegmentTreeQueries(t *testing.T) {
	tree := newSegmentTree("MWSWMMWWSS")

	tests := []struct {
		name        string
		left, right int
		want        int
	}{
		{name: "first three characters", left: 0, right: 3, want: 1},
		{name: "whole string", left: 0, right: 10, want: 2},
		{name: "suffix", left: 1, right: 10, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := tree.query(test.left, test.right); got != test.want {
				t.Errorf("query(%d, %d) = %d, want %d", test.left, test.right, got, test.want)
			}
		})
	}

	tree.update(8, 'M')
	tree.update(9, 'W')

	if got := tree.query(0, 10); got != 1 {
		t.Errorf("query after updates = %d, want 1", got)
	}
}

func TestSegmentTreePartialPattern(t *testing.T) {
	tree := newSegmentTree("MWSM")
	if got := tree.query(0, 4); got != 1 {
		t.Errorf("query() = %d, want 1", got)
	}
}
