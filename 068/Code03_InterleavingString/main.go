package main

// 交错字符串
// 给定三个字符串 s1、s2、s3
// 请帮忙验证s3是否由s1和s2交错组成
// 测试链接 : https://leetcode.cn/problems/interleaving-string/
var dp [101][101]bool

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	m, n := len(s1), len(s2)
	for i := range m + 1 {
		for j := range n + 1 {
			dp[i][j] = false
		}
	}
	dp[0][0] = true
	for i := range m + 1 {
		for j := range n + 1 {
			if i > 0 && s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if j > 0 && s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
			if i+j == len(s3) && dp[i][j] {
				return true
			}
		}
	}
	return false
}
