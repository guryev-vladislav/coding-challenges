package mapsynonyms

import (
	"testing"
)

const exampleSynonymLine = "hello world"
const secondSynonymLine = "world foo"

func TestMapSynonyms(t *testing.T) {
	cases := []struct {
		lines  []string
		target string
		want   string
		wantOk bool
	}{
		{
			lines:  []string{exampleSynonymLine, secondSynonymLine},
			target: "hello",
			want:   "world",
			wantOk: true,
		},
		{
			lines:  []string{exampleSynonymLine, secondSynonymLine},
			target: "foo",
			want:   "world",
			wantOk: true,
		},
		{
			lines:  []string{exampleSynonymLine, secondSynonymLine},
			target: "bar",
			want:   "",
			wantOk: false,
		},
	}

	for _, testCase := range cases {
		got, ok := MapSynonyms(testCase.lines, testCase.target)
		if got != testCase.want {
			t.Errorf("MapSynonyms(%v, %q) = %q, want %q", testCase.lines, testCase.target, got, testCase.want)
		}

		if ok != testCase.wantOk {
			t.Errorf("MapSynonyms(%v, %q) = (_, %v), want (_, %v)", testCase.lines, testCase.target, ok, testCase.wantOk)
		}
	}
}
