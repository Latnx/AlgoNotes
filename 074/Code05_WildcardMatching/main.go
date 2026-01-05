package main

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	return f2([]byte(s), []byte(p), dp)
}

// p[j]不等于*
// 情况1:i没了，
//
//	1.j也没了-》true
//	2.j还有-》判断是否是x*
//
// 情况2:j没了
//
//	false
//
// 情况3:都还有,j+1位置
//
//  1. 不是*
//  2. 是*
func f1(s, p []byte, i, j int, dp [][]int) bool {
	if dp[i][j] != 0 {
		return dp[i][j] == 1
	}
	res := false
	if i == len(s) {
		if j == len(p) {
			res = true
		} else {
			res = p[j] == '*' && f1(s, p, i, j+1, dp)
		}
	} else if j == len(p) {
		res = false
	} else {
		if p[j] != '*' {
			res = (s[i] == p[j] || p[j] == '?') && f1(s, p, i+1, j+1, dp)
		} else {
			p1 := f1(s, p, i, j+1, dp)
			p2 := f1(s, p, i+1, j, dp)
			res = p1 || p2
		}
	}
	if res {
		dp[i][j] = 1
	} else {
		dp[i][j] = 0
	}
	return res
}
func f2(s, p []byte, dp [][]bool) bool {
	n, m := len(s), len(p)
	for i := 0; i <= n; i++ {
		dp[i][m] = false
	}
	dp[n][m] = true
	for j := m - 1; j >= 0; j-- {
		dp[n][j] = p[j] == '*' && dp[n][j+1]
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if p[j] != '*' {
				dp[i][j] = (s[i] == p[j] || p[j] == '?') && dp[i+1][j+1]
			} else {
				dp[i][j] = dp[i+1][j] || dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
