package goforleetcode

// 题目保证了是个有效的星形图，所以只需要随便选两条边就行，其中重复的点必然是所求点
func findCenter(edges [][]int) int {
	// 选择 edges[0][0] 和 edges[0][1]
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
