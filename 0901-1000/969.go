package goforleetcode

func pancakeSort(arr []int) []int {
	var res []int
	for n := len(arr); n > 1; n-- {
		index := 0

		// fine max
		for i, v := range arr[:n] {
			if v > arr[index] {
				index = i
			}
		}

		if index == n-1 {
			continue
		}
		arr = reverseArray(arr, 0, index)
		arr = reverseArray(arr, 0, len(arr)-1)
		res = append(res, index)
		res = append(res, n)

	}
	return res
}

func reverseArray(arr []int, begin int, end int) []int {
	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
