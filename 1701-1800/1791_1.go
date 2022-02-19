package goforleetcode

//类似邻接矩阵存储无向图 求度数

func findCenter1(edges [][]int) int {
	n := len(edges) + 1         // 为什么+1，因为 edges 代表的是边，而点是要比边多1的
	degrees := make([]int, n+1) // map来存储每个顶点的度数,为啥又+1，因为点是从1开始数的
	for _, e := range edges {
		degrees[e[0]]++
		degrees[e[1]]++
	}
	for i, d := range degrees {
		if d == n-1 {
			return i
		}
	}
	return -1
}
