package leetcode

import "strings"

func maximumOddBinaryNumber(s string) string {
	numOfOne := strings.Count(s, "1")
	numOfZero := len(s) - numOfOne
	return strings.Repeat("1", numOfOne-1) + strings.Repeat("0", numOfZero) + "1"
}
