package leetcode

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums { // range return index and value
		// if 的神奇语法
		if res, ok := m[target-num]; ok { //如果存在 target-num ，就 return int[]
			return []int{res, i}
		}
		m[num] = i // i 存入好，好 return 回来
	}
	return nil
}
