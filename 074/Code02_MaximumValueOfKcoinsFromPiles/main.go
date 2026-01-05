package main

// 从栈中取出K个硬币的最大面值和
// 一张桌子上总共有 n 个硬币 栈 。每个栈有 正整数 个带面值的硬币
// 每一次操作中，你可以从任意一个栈的 顶部 取出 1 个硬币，从栈中移除它，并放入你的钱包里
// 给你一个列表 piles ，其中 piles[i] 是一个整数数组
// 分别表示第 i 个栈里 从顶到底 的硬币面值。同时给你一个正整数 k
// 请你返回在 恰好 进行 k 次操作的前提下，你钱包里硬币面值之和 最大为多少
// 测试链接 : https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/
func maxValueOfCoins(piles [][]int, k int) int {
	teams := len(piles)
	dp := make([][]int, teams+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= teams; i++ {
		// 计算前缀和，方便进入背包
		preSum := make([]int, len(piles[i-1])+1)
		for idx := 1; idx <= len(piles[i-1]); idx++ {
			preSum[idx] = preSum[idx-1] + piles[i-1][idx-1]
		}
		for j := 0; j <= k; j++ {
			// 只看第i组，情况一不选
			dp[i][j] = dp[i-1][j]
			// 若选择本组第一个,也可尝试选择2
			for idx := 1; idx <= len(piles[i-1]); idx++ {
				if j-idx >= 0 {
					dp[i][j] = max(dp[i][j], dp[i-1][j-idx]+preSum[idx])
				}
			}
		}
	}
	return dp[teams][k]
}
