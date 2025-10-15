package main

//	i < j 且 A[i] <= A[j], 求j-i 最大值
// 测试链接 : https://leetcode.cn/problems/maximum-width-ramp/

// 单调栈维持一个可能性
// 本题 上升或者相等就没有必要进栈，(在nums[1]>=nums[0]的情况下, 若(1, ?) 符合答案, (0, ?)也符合答案)
func maxWidthRamp(nums []int) int {
	stack := make([]int, len(nums))
	r := 0
	for i := range nums {
		// 先不弹出
		if r <= 0 || nums[stack[r-1]] > nums[i] {
			stack[r] = i
			r++
		}
	}
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		for r > 0 && nums[stack[r-1]] <= nums[i] {
			r--
			res = max(res, i-stack[r])
		}
	}
	return res
}
