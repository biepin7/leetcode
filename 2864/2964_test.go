package leetcode

import "testing"

func TestMaximumOddBinaryNumber(t *testing.T) {
	got := maximumOddBinaryNumber("010")
	want := "001"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
