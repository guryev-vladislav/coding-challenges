package polyglots

import (
	"reflect"
	"testing"
)

const (
	english = "English"
	french  = "French"
	russian = "Russian"
	spanish = "Spanish"
	german  = "German"
)

func TestPolyglots(t *testing.T) {
	tests := []struct {
		name      string
		students  [][]string
		common    []string
		allUnique []string
	}{
		{
			name: "valid input",
			students: [][]string{
				{english, russian, french},
				{english, russian, spanish},
				{english, russian, german},
			},
			common:    []string{english, russian},
			allUnique: []string{english, french, german, russian, spanish},
		},
		{
			name: "nil common languages input",
			students: [][]string{
				{english, french},
				{russian, spanish},
				{"Korean", german},
			},
			common:    []string{},
			allUnique: []string{english, french, german, "Korean", russian, spanish},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			polyglots := Polyglots(tt.students)
			if !reflect.DeepEqual(polyglots.commonLanguages, tt.common) {
				t.Errorf("Polyglots().commonLanguages = %v, want %v", polyglots.commonLanguages, tt.common)
			}

			if !reflect.DeepEqual(polyglots.allUniqueLanguages, tt.allUnique) {
				t.Errorf("Polyglots().allUniqueLanguages = %v, want %v", polyglots.allUniqueLanguages, tt.allUnique)
			}
		})
	}
}
