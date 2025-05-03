package main

// 返回峰值（左右小于本身）
// https://leetcode.cn/problems/find-peak-element/
func findPeakElement(nums []int) int {
	n := len(nums)
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[n-1] > nums[n-2] {
		return n - 1
	}
	var l, r = 1, n - 2
	for l <= r {
		middle := (l + r) / 2
		if nums[middle] < nums[middle-1] {
			r = middle - 1
		} else if nums[middle] < nums[middle+1] {
			l = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
