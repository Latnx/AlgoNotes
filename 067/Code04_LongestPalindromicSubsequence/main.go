package main

// 最长回文子序列
// 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度
// 测试链接 : https://leetcode.cn/problems/longest-palindromic-subsequence/
var dp [1000][1000]int

func longestPalindromeSubseq(s string) int {
	for i := range s {
		for j := range s {
			dp[i+1][j+1] = -1
		}
	}
	return f4([]byte(s))
}
func f2(s1 []byte, l, r int) int {
	if l > r {
		return 0
	} else if l == r {
		return 1
	}
	if dp[l][r] != -1 {
		return dp[l][r]
	}
	ans := 0
	if s1[l] == s1[r] {
		ans = f2(s1, l+1, r-1) + 2
	} else {
		ans = max(f2(s1, l+1, r), f2(s1, l, r-1))
	}
	dp[l][r] = ans
	return ans
}

// 错位，让出0，方便计算dp 原先的l = l-1
func f3(s1 []byte, l, r int) int {
	if l > r {
		return 0
	} else if l == r {
		return 1
	}
	if dp[l][r] != -1 {
		return dp[l][r]
	}
	ans := 0
	if s1[l-1] == s1[r-1] {
		ans = f2(s1, l+1, r-1) + 2
	} else {
		ans = max(f2(s1, l+1, r), f2(s1, l, r-1))
	}
	dp[l][r] = ans
	return ans
}

// 根据l+1, r-1可以推断，从下到上，从左到右，填一个上三角区域
// 需要左，下，左下
func f4(s []byte) int {
	n := len(s)
	for i := n; i >= 1; i-- {
		dp[i][i] = 1
		if j := i + 1; j <= n {
			if s[i-1] == s[j-1] {
				dp[i][j] = 2
			} else {
				dp[i][j] = 1
			}
		}
		for j := i + 2; j <= n; j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[1][n]
}
