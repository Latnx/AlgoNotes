package main

// 丑数 II
// 测试链接 : https://leetcode.cn/problems/ugly-number-ii/

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i, i2, i3, i5, a, b, c, cur := 2, 1, 1, 1, 0, 0, 0, 0; i <= n; i++ {
		a = dp[i2] * 2
		b = dp[i3] * 3
		c = dp[i5] * 5
		cur = min(a, b, c)
		if cur == a {
			i2++
		}
		if cur == b {
			i3++
		}
		if cur == c {
			i5++
		}
		dp[i] = cur
	}
	return dp[n]
}
