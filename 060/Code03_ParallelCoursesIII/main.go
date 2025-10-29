package main

// 并行课程 III
// 测试链接 : https://leetcode.cn/problems/parallel-courses-iii/
func minimumTime(n int, relations [][]int, time []int) int {
	adj := make([][]int, n)
	indegree := make([]int, n)
	preMax := make([]int, n)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	for _, relation := range relations {
		A, B := relation[0]-1, relation[1]-1
		adj[A] = append(adj[A], B)
		indegree[B]++
	}
	queue := make([]int, n)
	l, r := 0, 0
	for i := range indegree {
		if indegree[i] == 0 {
			queue[r] = i
			preMax[i] = time[i]
			r++
		}
	}
	for l < r {
		for _, outNow := range adj[queue[l]] {
			indegree[outNow]--
			preMax[outNow] = max(preMax[outNow], preMax[queue[l]]+time[outNow])
			if indegree[outNow] == 0 {
				queue[r] = outNow
				r++
			}
		}
		l++
	}
	res := 0
	for i := range adj {
		if len(adj[i]) == 0 {
			res = max(preMax[i], res)
		}
	}
	return res
}
