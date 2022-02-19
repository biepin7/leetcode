package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	arr = reverseArray(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func reverseArray(arr []int, begin int, end int) []int {
	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
