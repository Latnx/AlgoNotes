package main

import "math"

// 删掉1个数字后长度为k的子数组最大累加和
// 给定一个数组nums，求必须删除一个数字后的新数组中
// 长度为k的子数组最大累加和，删除哪个数字随意
// 对数器验证
func maxSum(nums []int, k int) int {
	n := len(nums)
	// 单调队列:维持窗口内最小值的更新结构
	window := make([]int, n)
	l, r := 0, 0
	sum := 0
	ans := math.MinInt
	for i := range nums {
		for l < r && nums[window[r-1]] >= nums[i] {
			r--
		}
		r++
		window[r] = i
		sum += nums[i]
		if i >= k {
			ans = max(ans, sum-nums[window[l]])
			if window[l] == i-k {
				l++
			}
			sum -= nums[i-k]
		}
	}
	return ans
}
