package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name           string
		p1, q1, p2, q2 int64
		want           int64
	}{
		{name: "sample one", p1: 12, q1: 10, p2: 13, q2: 12, want: 8},
		{name: "sample two", p1: 3, q1: 10, p2: 20, q2: 26, want: 1592},
		{name: "terminating fraction", p1: 1, q1: 2, p2: 1, q2: 2, want: 1},
		{name: "period of seven", p1: 1, q1: 7, p2: 1, q2: 7, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := solve(test.p1, test.q1, test.p2, test.q2); got != test.want {
				t.Errorf("solve() = %d, want %d", got, test.want)
			}
		})
	}
}
