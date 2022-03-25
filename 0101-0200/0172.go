package goforleetcode

func trailingZeroes(n int) int {
	res := 0
	div := 5
	for div <= n {
		res += n / div
		div *= 5
	}
	return res
}
