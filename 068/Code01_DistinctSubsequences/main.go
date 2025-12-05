package main

// 不同的子序列
// 测试链接 : https://leetcode.cn/problems/distinct-subsequences/
var dp [1001][1001]int

func numDistinct(s string, t string) int {
	for i := range len(s) {
		for j := range len(t) {
			dp[i+1][j+1] = -1
		}
	}
	return f1([]byte(s), []byte(t), len(s), len(t))
}
func f1(s, t []byte, sLen, tLen int) int {
	if tLen == 0 {
		return 1
	}
	if sLen == 0 {
		return 0
	}
	if dp[sLen][tLen] != -1 {
		return dp[sLen][tLen]
	}
	ans := 0
	if s[sLen-1] == t[tLen-1] {
		ans += f1(s, t, sLen-1, tLen-1)
	}
	ans += f1(s, t, sLen-1, tLen)
	dp[sLen][tLen] = ans
	return ans
}
func f2(s, t []byte) int {
	sLen, tLen := len(s), len(t)
	for i := range sLen + 1 {
		dp[i][0] = 1
	}
	for j := 1; j <= tLen; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen; j++ {
			dp[i][j] = 0
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
			dp[i][j] += dp[i-1][j]
		}
	}
	return dp[sLen][tLen]
}
