package main

// 编辑距离
// 测试链接 : https://leetcode.cn/problems/edit-distance/
var dp [501][501]int

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	for i := range m + 1 {
		dp[i][0] = i
	}
	for j := range n + 1 {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
