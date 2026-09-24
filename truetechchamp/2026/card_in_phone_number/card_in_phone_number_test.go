package main

import "testing"

func TestCountNumbersExample1(t *testing.T) {
	got := countNumbers("916")
	if got != 0 {
		t.Errorf("countNumbers(\"916\") = %d, want 0", got)
	}
}
