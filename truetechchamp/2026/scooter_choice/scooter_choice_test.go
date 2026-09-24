package main

import "testing"

func TestChooseOption(t *testing.T) {
	got := chooseOption(4, 3, 10)
	if got != 2 {
		t.Errorf("chooseOption(4, 3, 10) = %d, want 2", got)
	}
}
