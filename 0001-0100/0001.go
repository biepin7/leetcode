package goforleetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)           // key:num[i], value:i 数，数对应下标
	for i := 0; i < len(nums); i++ { // 遍历数组
		temp := target - nums[i]  // 需要的另一个值
		if _, ok := m[temp]; ok { // 若是map已经有需要的值	返回结果
			return []int{m[temp], i}
		}
		m[nums[i]] = i //若是没有需要的值，则添加到map中
	}
	return nil
}
