package main

// 环形数组的子数组最大累加和
// 给定一个数组nums，长度为n
// nums是一个环形数组，下标0和下标n-1是连在一起的
// 返回环形数组中，子数组最大累加和
// 测试链接 : https://leetcode.cn/problems/maximum-sum-circular-subarray/
func maxSubarraySumCircular(nums []int) int {
	dpMax, dpMin, resMax, resMin, all := nums[0], nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		all += nums[i]
		dpMax = max(nums[i]+dpMax, nums[i])
		resMax = max(resMax, dpMax)
		dpMin = min(nums[i]+dpMin, nums[i])
		resMin = min(resMin, dpMin)
	}
	if all != resMin {
		return max(resMax, all-resMin)
	}
	return resMax
}
