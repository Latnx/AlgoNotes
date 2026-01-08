package main

// 让字符串成为回文串的最少插入次数
// 给你一个字符串 s
// 每一次操作你都可以在字符串的任意位置插入任意字符
// 请你返回让s成为回文串的最少操作次数
// 测试链接 : https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/
func minInsertions(s string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	return f2([]byte(s), dp)
}
func f1(s []byte, l, r int) int {
	if l == r {
		return 0
	}
	if l+1 == r {
		if s[l] == s[r] {
			return 0
		} else {
			return 1
		}
	}
	if s[l] == s[r] {
		return f1(s, l+1, r-1)
	} else {
		return min(f1(s, l+1, r), f1(s, l, r-1)) + 1
	}
}
func f1_dp(s []byte, l, r int, dp [][]int) int {
	if l == r {
		return 0
	}
	if l+1 == r {
		if s[l] == s[r] {
			return 0
		} else {
			return 1
		}
	}
	if dp[l][r] != 0 {
		return dp[l][r]
	}
	res := 0
	if s[l] == s[r] {
		res = f1_dp(s, l+1, r-1, dp)
	} else {
		res = min(f1_dp(s, l+1, r, dp), f1_dp(s, l, r-1, dp)) + 1
	}
	dp[l][r] = res
	return res
}
func f2(s []byte, dp [][]int) int {
	n := len(s)
	for l := range n {
		dp[l][l] = 0
	}
	for l := 0; l+1 < n; l++ {
		if s[l] == s[l+1] {
			dp[l][l+1] = 0
		} else {
			dp[l][l+1] = 1
		}
	}
	for l := n - 3; l >= 0; l++ {
		for r := l + 2; r < n; r++ {
			if s[l] == s[r] {
				dp[l][r] = dp[l+1][r-1]
			} else {
				dp[l][r] = min(dp[l+1][r], dp[l][r-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}
