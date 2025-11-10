package main

import "container/heap"

// 二维接雨水
// 请计算图中形状最多能接多少体积的雨水。
// 测试链接 : https://leetcode.cn/problems/trapping-rain-water-ii/

// 先把边缘放进小根堆，弹出堆顶，添加相邻
// 堆中节点（x, y, 水线高度(max(自己的高度，添加我节点的水线)））

type hp [][3]int

func (a hp) Len() int           { return len(a) }
func (a hp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hp) Less(i, j int) bool { return a[i][2] < a[j][2] }
func (a *hp) Push(x any)        { *a = append(*a, x.([3]int)) }
func (a *hp) Pop() any          { x := (*a)[len(*a)-1]; *a = (*a)[:len(*a)-1]; return x }

func trapRainWater(heightMap [][]int) int {
	h := make(hp, 0)
	m, n := len(heightMap), len(heightMap[0])
	visited := make([][]bool, m)
	move := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := range m {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		h = append(h, [3]int{0, i, heightMap[0][i]})
		h = append(h, [3]int{m - 1, i, heightMap[m-1][i]})
		visited[0][i] = true
		visited[m-1][i] = true
	}
	for i := 1; i < m-1; i++ {
		h = append(h, [3]int{i, 0, heightMap[i][0]})
		h = append(h, [3]int{i, n - 1, heightMap[i][n-1]})
		visited[i][0] = true
		visited[i][n-1] = true
	}
	heap.Init(&h)
	res := 0
	for len(h) != 0 {
		u := heap.Pop(&h).([3]int)
		res += u[2] - heightMap[u[0]][u[1]]
		for _, to := range move {
			x, y := u[0]+to[0], u[1]+to[1]
			if x >= 0 && y >= 0 && x < m && y < n && !visited[x][y] {
				visited[x][y] = true
				heap.Push(&h, [3]int{x, y, max(heightMap[x][y], u[2])})
			}
		}
	}
	return res
}
