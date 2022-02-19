package goforleetcode

func maxNumberOfBalloons(text string) int {
	// balloon
	m := make(map[rune]int)
	for _, v := range text {
		m[v]++
	}
	m['l'] /= 2
	m['o'] /= 2

	return min(m['n'], min(m['b'], min(m['a'], min(m['l'], m['o']))))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
