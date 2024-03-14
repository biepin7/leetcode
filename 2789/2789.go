package leetcode

func maxArrayValue(nums []int) int64 {
	size := len(nums)
	num := int64(nums[size-1])

	for i := size - 2; i >= 0; i-- {
		if int64(nums[i]) <= num {
			num = int64(nums[i]) + num
		} else {
			num = int64(nums[i])
		}
	}

	return num
}
