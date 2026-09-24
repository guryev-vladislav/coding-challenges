package calldetailsrecovery

import "testing"

func TestRecoverDurationExamples(t *testing.T) {
	tests := []struct {
		name    string
		records []int
		want    int
	}{
		{
			name:    "example 1",
			records: []int{11, 4, -6, 8, -9, -16},
			want:    8,
		},
		{
			name:    "example 2 first case",
			records: []int{9, -3, -16, 18},
			want:    -1,
		},
		{
			name:    "example 2 second case",
			records: []int{8, 8, 8, -8, -8, -8},
			want:    0,
		},
		{
			name:    "example 2 third case",
			records: []int{1, -2, 3, -4, 5},
			want:    -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := recoverDuration(test.records); got != test.want {
				t.Errorf("recoverDuration(%v) = %d, want %d", test.records, got, test.want)
			}
		})
	}
}
