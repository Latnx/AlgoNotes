package main

// 按奇偶排序数组II
// 给定一个非负整数数组 nums。nums 中一半整数是奇数 ，一半整数是偶数
// 对数组进行排序，以便当 nums[i] 为奇数时，i也是奇数
// 当 nums[i] 为偶数时， i 也是 偶数
// 你可以返回 任何满足上述条件的数组作为答案
// 测试链接 : https://leetcode.cn/problems/sort-array-by-parity-ii/
func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	last := len(nums) - 1
	for even < len(nums) && odd < len(nums) {
		if nums[last]%2 == 0 {
			nums[even], nums[last] = nums[last], nums[even]
			even += 2
		} else {
			nums[odd], nums[last] = nums[last], nums[odd]
			odd += 2
		}
	}
	return nums
}
