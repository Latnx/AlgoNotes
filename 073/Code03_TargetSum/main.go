package main

// 目标和
// 给你一个非负整数数组 nums 和一个整数 target 。
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数
// 可以构造一个表达式
// 例如nums=[2, 1]，可以在2之前添加'+' ，在1之前添加'-'
// 然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的，运算结果等于 target 的不同表达式的数目
// 测试链接 : https://leetcode.cn/problems/target-sum/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if target < 0 {
		target = -target
	}
	if sum < target || target%2 != sum%2 {
		return 0
	}
	return findTarget(nums, (target+sum)/2)
}
func findTarget(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := target; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
