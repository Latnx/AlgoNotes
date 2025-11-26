package main

import "math"

// 最低票价
// 测试链接 : https://leetcode.cn/problems/minimum-cost-for-tickets/
var duration = []int{1, 7, 30}
var dp []int

// days 一年中要旅行的日子
// costs 1, 7, 30 天通行证的花费
func mincostTickets(days []int, costs []int) int {
	return f2(days, costs, 0)
}

// 从下标i开始到最后花费多少元，O(3^n)
func f2(days []int, costs []int, i int) int {
	if i == len(days) {
		return 0
	}
	ans := math.MaxInt
	// dfs
	for k, j := 0, i; k < len(costs); k++ {
		for j < len(days) && days[j] < days[i]+duration[k] {
			j++
		}
		ans = min(ans, costs[k]+f1(days, costs, j))
	}
	return ans
}

// 根据可变参数，创建缓存表，查询（记忆化搜索）
func f1(days []int, costs []int, i int) int {
	if i == len(days) {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}
	ans := math.MaxInt
	// dfs
	for k, j := 0, i; k < len(costs); k++ {
		for j < len(days) && days[j] < days[i]+duration[k] {
			j++
		}
		ans = min(ans, costs[k]+f1(days, costs, j))
	}
	dp[i] = ans
	return ans
}

// 从底到顶
func f0(days []int, costs []int) int {
	n := len(days)
	dp = make([]int, 366)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[n] = 0

	// dfs
	for i := n - 1; i >= 0; i-- {
		for k, j := 0, i; k < len(costs); k++ {
			for j < len(days) && days[j] < days[i]+duration[k] {
				j++
			}
			dp[i] = min(dp[i], costs[k]+dp[j])
		}
	}
	return dp[0]
}
