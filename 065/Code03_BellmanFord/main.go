package main

import (
	"math"
	"slices"
)

// Bellman-Ford算法应用（不是模版）
// 测试链接 : https://leetcode.cn/problems/cheapest-flights-within-k-stops/
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cur := make([]int, n)
	for i := range n {
		cur[i] = math.MaxInt
	}
	cur[src] = 0
	for range k + 1 {
		next := slices.Clone(cur)
		for _, flight := range flights {
			if cur[flight[0]] != math.MaxInt {
				next[flight[1]] = min(next[flight[1]], cur[flight[0]]+flight[2])
			}
		}
		cur = next
	}
	if cur[dst] == math.MaxInt {
		return -1
	}
	return cur[dst]
}
