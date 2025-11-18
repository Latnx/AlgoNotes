package main

import (
	"container/heap"
	"math"
)

// Dijkstra算法模版
// 网络延迟时间
// 测试链接 : https://leetcode.cn/problems/network-delay-time
type hp [][2]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp) Pop() any          { l := len(*h) - 1; x := (*h)[l]; *h = (*h)[:l]; return x }

var distance []int
var visited []bool

func networkDelayTime(times [][]int, n int, k int) int {
	visited = make([]bool, n+1)
	distance = make([]int, n+1)
	for i := 1; i < len(distance); i++ {
		distance[i] = math.MaxInt
	}
	// 建图
	adj := make([][][2]int, n+1)
	for i := range adj {
		adj[i] = make([][2]int, 0)
	}
	for _, time := range times {
		adj[time[0]] = append(adj[time[0]], [2]int{time[1], time[2]})
	}
	// 建堆
	// 弹出过-忽略
	// 没弹出过， 处理所有边
	// 让其他没弹出的节点，距离变小的记录加入堆
	h := hp{{k, 0}}
	distance[k] = 0

	for len(h) > 0 {
		node := heap.Pop(&h).([2]int)[0]
		if visited[node] {
			continue
		}
		visited[node] = true
		for _, neighbor := range adj[node] {
			w, v := neighbor[0], neighbor[1]
			if !visited[w] && distance[node]+v < distance[w] {
				distance[w] = distance[node] + v
				heap.Push(&h, [2]int{w, distance[w]})
			}
		}
	}
	res := 0
	for _, dis := range distance {
		res = max(dis, res)
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
