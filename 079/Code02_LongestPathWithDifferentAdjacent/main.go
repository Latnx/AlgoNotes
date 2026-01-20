package main

// 相邻字符不同的最长路径
// 给你一棵 树（即一个连通、无向、无环图），根节点是节点 0
// 这棵树由编号从 0 到 n - 1 的 n 个节点组成
// 用下标从 0 开始、长度为 n 的数组 parent 来表示这棵树
// 其中 parent[i] 是节点 i 的父节点
// 由于节点 0 是根节点，所以 parent[0] == -1
// 另给你一个字符串 s ，长度也是 n ，其中 s[i] 表示分配给节点 i 的字符
// 请你找出路径上任意一对相邻节点都没有分配到相同字符的 最长路径
// 并返回该路径的长度
// 测试链接 : https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/
func longestPath(parent []int, s string) int {
	n := len(parent)
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for i := 1; i < n; i++ {
		graph[parent[i]] = append(graph[parent[i]], i)
	}
	res, _ := dfs(graph, []byte(s), 0)
	return res
}

// 最大，单边
func dfs(graph [][]int, s []byte, now int) (int, int) {
	if len(graph[now]) == 0 {
		return 1, 1
	}
	// 第一大单边，第二大单边
	max1, max2 := 0, 0
	maxPath := 1
	for _, v := range graph[now] {
		a, b := dfs(graph, s, v)
		maxPath = max(maxPath, a)
		// 记录最大两个个不等子节点
		if s[now] != s[v] {
			if b > max1 {
				max2 = max1
				max1 = b
			} else if b > max2 {
				max2 = b
			}
		}
	}
	// 假设已经遍历完了，开始讨论
	// maxPath =  max(不选当前点 遍历所有子树的最大值， 选当前点 找到不相邻的两个最大边+1)
	// maxPathS = max(找到最大的和当前不同的+1)
	maxPath = max(maxPath, max1+max2+1)
	return maxPath, max1 + 1
}
