package leetcode

import "testing"

func Test_maxArrayValue(t *testing.T) {
	nums := []int{2, 3, 7, 9, 3}
	got := maxArrayValue(nums)
	want := int64(21)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
