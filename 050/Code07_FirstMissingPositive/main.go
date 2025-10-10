package main

// 缺失的第一个正数
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
// 测试链接 : https://leetcode.cn/problems/first-missing-positive/

// 双指针
// l左边都是在位置的, l是待找出的位置，最后返回 l+1 就是找不到的最小正整数
// r右边都是垃圾，垃圾的意思就是，小于已经排序好的，或者是不可能抵达的，或者是
// arr[l] == l+1 l++
// arr[l] <=l 垃圾
// arr[l] > r 垃圾
// arr[arr[l]-1] == arr[l] 重复，垃圾
// 交换
func firstMissingPositive(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == l+1 {
			l++
		} else if nums[l] <= l || nums[l] > r || nums[nums[l]-1] == nums[l] {
			r--
			nums[l], nums[r] = nums[r], nums[l]
		} else {
			nums[l], nums[nums[l]-1] = nums[nums[l]-1], nums[l]
		}
	}
	return l + 1
}
