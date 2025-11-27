package main

// 最长公共子序列
// 测试链接 : https://leetcode.cn/problems/longest-common-subsequence/
var dp [1000][1000]int

func longestCommonSubsequence(text1 string, text2 string) int {
	for i := range text1 {
		for j := range text2 {
			dp[i+1][j+1] = -1
		}
	}
	return f1([]byte(text1), []byte(text2), len(text1)-1, len(text2)-1)
}
func f1(s1, s2 []byte, index1, index2 int) int {
	if index1 < 0 || index2 < 0 {
		return 0
	}
	if dp[index1][index2] != -1 {
		return dp[index1][index2]
	}
	p1 := f1(s1, s2, index1-1, index2-1)
	p2 := f1(s1, s2, index1-1, index2)
	p3 := f1(s1, s2, index1, index2-1)
	p4 := 0
	if s1[index1] == s2[index2] {
		p4 = p1 + 1
	}
	dp[index1][index2] = max(p1, p2, p3, p4)
	return max(p1, p2, p3, p4)
}
func f2(s1, s2 []byte, index1, index2 int) int {
	if index1 < 0 || index2 < 0 {
		return 0
	}
	if dp[index1][index2] != -1 {
		return dp[index1][index2]
	}
	ans := 0
	if s1[index1] == s2[index2] {
		ans = f2(s1, s2, index1-1, index2-1) + 1
	} else {
		ans = max(f2(s1, s2, index1-1, index2), f2(s1, s2, index1, index2-1))
	}
	dp[index1][index2] = ans
	return ans
}
func f3(s1, s2 []byte) int {
	n, m := len(s1), len(s2)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
