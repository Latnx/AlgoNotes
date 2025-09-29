package main

import "math"

// 累加和大于等于target的最短子数组长度
// 给定一个含有 n 个正整数的数组和一个正整数 target
// 找到累加和 >= target 的长度最小的子数组并返回其长度
// 如果不存在符合条件的子数组返回0
// 测试链接 : https://leetcode.cn/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	sum := 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for sum-nums[l] >= target {
			sum -= nums[l]
			l++
		}
		if sum >= target {
			res = min(res, r-l+1)
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}
