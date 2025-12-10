package main

import "math"

// 打家劫舍 IV
// 沿街有一排连续的房屋。每间房屋内都藏有一定的现金
// 现在有一位小偷计划从这些房屋中窃取现金
// 由于相邻的房屋装有相互连通的防盗系统，所以小偷不会窃取相邻的房屋
// 小偷的 窃取能力 定义为他在窃取过程中能从单间房屋中窃取的 最大金额
// 给你一个整数数组 nums 表示每间房屋存放的现金金额
// 第i间房屋中放有nums[i]的钱数
// 另给你一个整数k，表示小偷需要窃取至少 k 间房屋
// 返回小偷需要的最小窃取能力值
// 测试链接 : https://leetcode.cn/problems/house-robber-iv/

// 一个数组，至少选择不连续的k的数字，返回所选数字中的最大值，求最小为多少
// 二分答案法，当能力值是k的时候可以选择不相邻的数字
func minCapability(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	minCap, maxCap := math.MaxInt, math.MinInt
	for i := range nums {
		minCap = min(minCap, nums[i])
		maxCap = max(maxCap, nums[i])
	}
	res := 0
	for minCap <= maxCap {
		midCap := (minCap + maxCap) / 2
		if f1(nums, midCap) < k {
			minCap = midCap + 1
		} else if f1(nums, midCap) >= k {
			res = midCap
			maxCap = midCap - 1
		}
	}
	return res
}

func f1(nums []int, capability int) (k int) {
	dp := make([]int, len(nums))
	if nums[0] <= capability {
		dp[0] = 1
	}
	if nums[1] <= capability {
		dp[1] = 1
	}
	dp[1] = max(dp[0], dp[1])

	for i := 2; i < len(dp); i++ {
		if capability >= nums[i] {
			dp[i] = max(dp[i-1], dp[i-2]+1)
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
