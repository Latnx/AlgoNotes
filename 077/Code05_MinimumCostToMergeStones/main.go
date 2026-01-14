package main

import "math"

// 合并石头的最低成本
// 有 n 堆石头排成一排，第 i 堆中有 stones[i] 块石头
// 每次 移动 需要将 连续的 k 堆石头合并为一堆，而这次移动的成本为这 k 堆中石头的总数
// 返回把所有石头合并成一堆的最低成本
// 如果无法合并成一堆返回-1
// 测试链接 : https://leetcode.cn/problems/minimum-cost-to-merge-stones/
func mergeStones_best(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	// [l..r]上的累加和为presum[r+1]-presum[l]
	presum := make([]int, n+1)
	for i, val := range stones {
		presum[i+1] = presum[i] + val
	}
	// dp[l][r]:l到r上最终份数，代价是多少
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}
	for l := n - 2; l >= 0; l-- {
		for r := l + 1; r < n; r++ {
			ans := math.MaxInt
			for m := l; m < r; m += k - 1 {
				ans = min(ans, dp[l][m]+dp[m+1][r])
			}
			if (r-l)%(k-1) == 0 {
				ans += presum[r+1] - presum[l]
			}
			dp[l][r] = ans
		}
	}
	return dp[0][n-1]
}
func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	// [l..r]上的累加和为presum[r+1]-presum[l]
	presum = make([]int, n+1)
	for i, val := range stones {
		presum[i+1] = presum[i] + val
	}
	// dp[l][r]:l到r上最终份数，代价是多少
	dp = make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	return f1(stones, k, 0, n-1, 1)
}

var presum []int
var dp [][][]int

// [l...r]合成p堆
func f1(stones []int, k, l, r, p int) int {
	// 退出条件
	if l == r {
		return 0
	}
	if dp[l][r][p] != 0 {
		return dp[l][r][p]
	}
	// 递归体
	if p == 1 {
		dp[l][r][p] = f1(stones, k, l, r, k) + presum[r+1] - presum[l]
		return dp[l][r][p]
	}
	ans := math.MaxInt
	for m := l; m < r; m += k - 1 {
		ans = min(ans, f1(stones, k, l, m, 1)+f1(stones, k, m+1, r, p-1))
	}
	dp[l][r][p] = ans
	return ans
}
