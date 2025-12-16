package main

// 乘积最大子数组
// 给你一个整数数组 nums
// 请你找出数组中乘积最大的非空连续子数组
// 并返回该子数组所对应的乘积
// 测试链接 : https://leetcode.cn/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	res := nums[0]
	preP, preN := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curP := max(preP*nums[i], nums[i], preN*nums[i])
		curN := min(preP*nums[i], nums[i], preN*nums[i])
		preP, preN = curP, curN
		res = max(res, preP)
	}
	return res
}
