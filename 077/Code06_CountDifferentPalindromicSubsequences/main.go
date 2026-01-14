package main

// 统计不同回文子序列
// 给你一个字符串s，返回s中不同的非空回文子序列个数
// 由于答案可能很大，答案对 1000000007 取模
// 测试链接 : https://leetcode.cn/problems/count-different-palindromic-subsequences/
func countPalindromicSubsequences(s string) int {
	n := len(s)
	dp = make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}
	return f1([]byte(s))
}

// [i,j]上有多少不同的回文子序列
var dp [][]int

const MOD = 1000000007

func f1(s []byte) int {
	n := len(s)
	for i := range n {
		dp[i][i] = 1
	}
	for i := 0; i+1 < n; i++ {
		dp[i][i+1] = 2
	}
	// 创建一个结构，用来找，元素相同左边和右边
	var m [26]int
	for i := range m {
		m[i] = -1
	}
	leftIndex := make([]int, n)
	for i := range n {
		leftIndex[i] = m[s[i]-'a']
		m[s[i]-'a'] = i
	}
	// 找右边
	for i := range m {
		m[i] = -1
	}
	rightIndex := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		rightIndex[i] = m[s[i]-'a']
		m[s[i]-'a'] = i
	}
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			if s[i] != s[j] {
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1] + MOD) % MOD
			} else if s[i] == s[j] {
				l := rightIndex[i]
				r := leftIndex[j]
				if l == j {
					dp[i][j] = (dp[i+1][j-1]*2 + 2) % MOD
				} else if l == r {
					dp[i][j] = (dp[i+1][j-1]*2 + 1) % MOD
				} else {
					dp[i][j] = (dp[i+1][j-1]*2 - dp[l+1][r-1] + MOD) % MOD
				}
			}
		}
	}
	return dp[0][n-1]
}
