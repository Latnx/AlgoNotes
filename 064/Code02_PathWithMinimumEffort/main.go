package main

import (
	"container/heap"
	"math"
)

// 最小体力消耗路径
// 测试链接 ：https://leetcode.cn/problems/path-with-minimum-effort/
type hp [][3]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.([3]int)) }
func (h *hp) Pop() any          { l := len(*h) - 1; x := (*h)[l]; *h = (*h)[:l]; return x }

var distance [][]int
var visited [][]bool

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	distance = make([][]int, len(heights))
	visited = make([][]bool, len(heights))
	for i := range heights {
		distance[i] = make([]int, len(heights[0]))
		visited[i] = make([]bool, len(heights[0]))
	}
	for i := range distance {
		for j := range distance[i] {
			distance[i][j] = math.MaxInt
		}
	}
	move := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	distance[0][0] = 0
	h := hp{{0, 0, 0}}
	for len(h) > 0 {
		node := heap.Pop(&h).([3]int)
		x, y, _ := node[0], node[1], node[2]
		if visited[x][y] {
			continue
		}
		visited[x][y] = true
		for _, to := range move {
			toX, toY := x+to[0], y+to[1]
			if toX >= 0 && toY >= 0 && toX < n && toY < m && !visited[toX][toY] {
				distance[toX][toY] = min(distance[toX][toY], max(abs(heights[x][y]-heights[toX][toY]), distance[x][y]))
				heap.Push(&h, [3]int{toX, toY, distance[toX][toY]})
			}
		}
	}
	return distance[n-1][m-1]
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
