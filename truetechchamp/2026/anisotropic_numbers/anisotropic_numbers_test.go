package main

import "testing"

func TestCountNumbers(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		want   int
	}{
		{
			name:   "invalid prefix",
			prefix: "8916",
			want:   0,
		},
		{
			name:   "valid prefix",
			prefix: "8906",
			want:   24,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := countNumbers(test.prefix); got != test.want {
				t.Errorf("countNumbers(%q) = %d, want %d", test.prefix, got, test.want)
			}
		})
	}
}
