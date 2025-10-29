package main

// 喧闹和富有
// 测试链接 : https://leetcode.cn/problems/loud-and-rich/
func loudAndRich(richer [][]int, quiet []int) []int {
	adj := make([][]int, len(quiet))
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	indegree := make([]int, len(quiet))
	for i := range richer {
		adj[richer[i][0]] = append(adj[richer[i][0]], richer[i][1])
		indegree[richer[i][1]]++
	}
	queue := make([]int, len(quiet))
	l, r := 0, 0
	res := make([]int, len(quiet))

	for i := range indegree {
		if indegree[i] == 0 {
			queue[r] = i
			r++
		}
		res[i] = i
	}
	for l < r {
		for j := range adj[queue[l]] {
			indegree[adj[queue[l]][j]]--
			if quiet[res[queue[l]]] < quiet[res[adj[queue[l]][j]]] {
				res[adj[queue[l]][j]] = res[queue[l]]
			}
			if indegree[adj[queue[l]][j]] == 0 {
				queue[r] = adj[queue[l]][j]
				r++
			}
		}
		l++
	}
	return res
}
