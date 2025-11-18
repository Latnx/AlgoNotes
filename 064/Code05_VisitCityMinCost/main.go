package main

import (
	"container/heap"
	"math"
)

// 电动车游城市
// 测试链接 : https://leetcode.cn/problems/DFPeFJ/

// 点(城市，电量)和到达这个点花费的时间
type hp [][3]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.([3]int)) }
func (h *hp) Pop() any          { l := len(*h) - 1; x := (*h)[l]; *h = (*h)[:l]; return x }

var distance [][100]int
var visited [][100]bool

func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	n := len(charge)
	distance = make([][100]int, n)
	visited = make([][100]bool, n)
	for i := range distance {
		for j := range distance[i] {
			distance[i][j] = math.MaxInt
		}
	}
	// 建图
	adj := make([][][2]int, n)
	for i := range adj {
		adj[i] = make([][2]int, 0)
	}
	for _, path := range paths {
		adj[path[0]] = append(adj[path[0]], [2]int{path[1], path[2]})
		adj[path[1]] = append(adj[path[1]], [2]int{path[0], path[2]})
	}
	// dj
	// 添加初始点
	h := hp{{start, 0, 0}}
	distance[start][0] = 0
	// 带状态dj
	for len(h) > 0 {
		node := heap.Pop(&h).([3]int)
		city, battery, time := node[0], node[1], node[2]
		if city == end {
			return distance[city][battery]
		}
		if visited[city][battery] {
			continue
		}
		visited[city][battery] = true
		if battery < cnt {
			nt := time + charge[city]
			if nt < distance[city][battery+1] {
				distance[city][battery+1] = nt
				heap.Push(&h, [3]int{city, battery + 1, nt})
			}
		}
		for _, to := range adj[city] {
			nextCity, cost := to[0], to[1]
			if battery >= cost {
				nt := time + cost
				if nt < distance[nextCity][battery-cost] {
					distance[nextCity][battery-cost] = nt
					heap.Push(&h, [3]int{nextCity, battery - cost, nt})
				}
			}
		}
	}
	// 结果收集
	res := math.MaxInt
	for _, val := range distance[end] {
		res = min(res, val)
	}
	return res
}
