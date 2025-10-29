package main

import "strings"

// 火星词典
// 测试链接 : https://leetcode.cn/problems/alien-dictionary/
// 测试链接(不需要会员) : https://leetcode.cn/problems/Jf1JuT/
func alienOrder(words []string) string {
	indegree := make([]int, 26)
	total := 0
	// 初始化入度，以及字符有多少种
	for i := range indegree {
		indegree[i] = -1
	}
	for i := range words {
		for j := range words[i] {
			indegree[words[i][j]-'a'] = 0
		}
	}
	for _, val := range indegree {
		if val == 0 {
			total++
		}
	}

	// 建图，邻接矩阵
	adj := make([][]int, 26)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	for i := 1; i < len(words); i++ {
		j := 0
		lenth := min(len(words[i-1]), len(words[i]))
		for ; j < lenth && words[i][j] == words[i-1][j]; j++ {
		}
		if j < lenth {
			adj[words[i-1][j]-'a'] = append(adj[words[i-1][j]-'a'], int(words[i][j]-'a'))
			indegree[words[i][j]-'a']++
		}
		// 排序出错，特殊情况
		if j < len(words[i-1]) && j == len(words[i]) {
			return ""
		}
	}
	queue := make([]int, 26)
	l, r := 0, 0
	for i := range adj {
		if indegree[i] == 0 {
			queue[r] = i
			r++
		}
	}

	for l < r {
		for j := range adj[queue[l]] {
			indegree[adj[queue[l]][j]]--
			if indegree[adj[queue[l]][j]] == 0 {
				queue[r] = adj[queue[l]][j]
				r++
			}
		}
		l++
	}
	var res strings.Builder
	if r != total {
		return res.String()
	}

	for i := range r {
		res.WriteByte(byte(queue[i] + 'a'))
	}
	return res.String()
}
