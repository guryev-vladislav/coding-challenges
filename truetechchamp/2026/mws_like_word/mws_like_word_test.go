package main

import "testing"

func TestBuildWordExample1(t *testing.T) {
	got := buildWord(3)
	if got != "MSW" {
		t.Errorf("buildWord(3) = %q, want %q", got, "MSW")
	}
}
