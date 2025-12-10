package main

// 子数组最大累加和
// 给你一个整数数组 nums
// 返回非空子数组的最大累加和
// 测试链接 : https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	dp := nums[0]
	res := dp
	for i := 1; i < len(nums); i++ {
		if dp < 0 {
			dp = nums[i]
		} else if dp >= 0 {
			dp += nums[i]
		}
		res = max(res, dp)
	}
	return res
}
