package main

import "testing"

func TestMaxRevenue(t *testing.T) {
	tests := []struct {
		name       string
		upgrade    int
		expansions []expansion
		want       int64
	}{
		{
			name:    "sample",
			upgrade: 10,
			expansions: []expansion{
				{capacityLoss: 51, revenue: 2026},
				{capacityLoss: 52, revenue: 2027},
				{capacityLoss: 50, revenue: 2},
			},
			want: 2028,
		},
		{
			name:    "upgrade is unavailable",
			upgrade: 0,
			expansions: []expansion{
				{capacityLoss: 60, revenue: 10},
				{capacityLoss: 60, revenue: 20},
			},
			want: 20,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maxRevenue(test.upgrade, test.expansions); got != test.want {
				t.Errorf("maxRevenue() = %d, want %d", got, test.want)
			}
		})
	}
}
