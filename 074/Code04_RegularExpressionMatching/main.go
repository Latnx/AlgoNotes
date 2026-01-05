package main

// 正则表达式匹配
// 给你字符串s、字符串p
// s中一定不含有'.'、'*'字符，p中可能含有'.'、'*'字符
// '.' 表示可以变成任意字符，数量1个
// '*' 表示可以让 '*' 前面那个字符数量任意(甚至可以是0个)
// p中即便有'*'，一定不会出现以'*'开头的情况，也一定不会出现多个'*'相邻的情况(无意义)
// 请实现一个支持 '.' 和 '*' 的正则表达式匹配
// 返回p的整个字符串能不能匹配出s的整个字符串
// 测试链接 : https://leetcode.cn/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	sCh, pCh := []byte(s), []byte(p)
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(p))
	}
	return f2(sCh, pCh, 0, 0, dp)
}

// p[j]不等于*
// 情况1:p没了，
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
func f1(s, p []byte, i, j int) bool {
	if i == len(s) {
		if j == len(p) {
			return true
		} else {
			return j+1 < len(p) && p[j+1] == '*' && f1(s, p, i, j+2)
		}
	} else if j == len(p) {
		return false
	} else {
		if j+1 == len(p) || p[j+1] != '*' {
			return (s[i] == p[j] || p[j] == '.') && f1(s, p, i+1, j+1)
		} else {
			p1 := f1(s, p, i, j+2)
			p2 := (s[i] == p[j] || p[j] == '.') && f1(s, p, i+1, j)
			return p1 || p2
		}
	}
}
func f2(s, p []byte, i, j int, dp [][]int) bool {
	if dp[i][j] != 0 {
		return dp[i][j] == 1
	}
	res := false
	if i == len(s) {
		if j == len(p) {
			res = true
		} else {
			res = j+1 < len(p) && p[j+1] == '*' && f2(s, p, i, j+2, dp)
		}
	} else if j == len(p) {
		res = false
	} else {
		if j+1 == len(p) || p[j+1] != '*' {
			res = (s[i] == p[j] || p[j] == '.') && f2(s, p, i+1, j+1, dp)
		} else {
			p1 := f2(s, p, i, j+2, dp)
			p2 := (s[i] == p[j] || p[j] == '.') && f2(s, p, i+1, j, dp)
			res = p1 || p2
		}
	}
	if res {
		dp[i][j] = 1
	} else {
		dp[i][j] = 2
	}
	return res
}
