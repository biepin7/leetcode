package _303

type NumArray struct {
	nums []int //NumArray(int[] nums) 使用数组 nums 初始化对象
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1) // 前缀和数组
	for i, v := range nums {         // 计算前缀和数组
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums} //
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
