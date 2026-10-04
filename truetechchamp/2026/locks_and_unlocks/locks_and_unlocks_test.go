package main

import "testing"

func TestSolveKey(t *testing.T) {
	tests := []struct {
		name string
		code string
		want int
	}{
		{name: "no overlap", code: "3791", want: 0},
		{name: "invalid key", code: "3795", want: -1},
		{name: "overlap", code: "5469", want: 1},
		{name: "all nodes", code: "123456789", want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := solveKey(test.code); got != test.want {
				t.Errorf("solveKey(%q) = %d, want %d", test.code, got, test.want)
			}
		})
	}
}
