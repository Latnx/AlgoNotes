package main

import (
	"container/list"
	"math"
)

// 到达角落需要移除障碍物的最小数目
// 测试链接 : https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/
func minimumObstacles(grid [][]int) int {
	type point [2]int
	n, m := len(grid), len(grid[0])
	deque := list.New()
	distance := make([]int, n*m)
	for i := 1; i < n*m; i++ {
		distance[i] = math.MaxInt
	}
	move := [4]point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	deque.PushFront(point{0, 0})
	for deque.Len() > 0 {
		u := deque.Front().Value.(point)
		deque.Remove(deque.Front())
		for _, to := range move {
			x, y := u[0]+to[0], u[1]+to[1]
			if x >= 0 && x < n && y >= 0 && y < m {
				if distance[x*m+y] > distance[u[0]*m+u[1]]+grid[x][y] {
					distance[x*m+y] = distance[u[0]*m+u[1]] + grid[x][y]
					if grid[x][y] == 0 {
						deque.PushFront(point{x, y})
					} else {
						deque.PushBack(point{x, y})
					}
				}
			}
		}
	}
	return distance[n*m-1]
}
