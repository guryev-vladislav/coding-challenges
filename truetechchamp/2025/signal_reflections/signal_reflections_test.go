package signalreflections

import "testing"

func TestReflectionsExamples(t *testing.T) {
	tests := []struct {
		p, q int64
		want int64
	}{
		{p: 1, q: 2, want: 359},
		{p: 30, q: 1, want: 5},
		{p: 8, q: 9, want: 404},
		{p: 45, q: 2, want: 7},
		{p: 3, q: 14, want: 839},
		{p: 15, q: 1, want: 11},
	}

	for _, test := range tests {
		if got := reflections(test.p, test.q); got != test.want {
			t.Errorf("reflections(%d, %d) = %d, want %d", test.p, test.q, got, test.want)
		}
	}
}
