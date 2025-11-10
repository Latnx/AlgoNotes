package main

import (
	"container/list"
	"math"
)

// 使网格图至少有一条有效路径的最小代价
// 测试链接 : https://leetcode.cn/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/
// 01dfs
func minCost(grid [][]int) int {
	type point [2]int
	n, m := len(grid), len(grid[0])
	deque := list.New()
	move := [4]point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	distance := make([]int, m*n)
	for i := 1; i < len(distance); i++ {
		distance[i] = math.MaxInt
	}
	deque.PushBack(point{0, 0})
	for deque.Len() > 0 {
		front := deque.Front()
		u := front.Value.(point)
		deque.Remove(front)
		for po, to := range move {
			x, y := u[0]+to[0], u[1]+to[1]
			if x >= 0 && y >= 0 && x < n && y < m {
				diff := 0
				if po == grid[u[0]][u[1]]-1 {
					diff = 1
				}
				if distance[u[0]*m+u[1]]+diff < distance[x*m+y] {
					distance[x*m+y] = distance[u[0]*m+u[1]] + diff
					if diff == 0 {
						deque.PushFront(point{x, y})
					} else {
						deque.PushBack(point{x, y})
					}
				}
			}
		}
	}
	return distance[m*n-1]
}
