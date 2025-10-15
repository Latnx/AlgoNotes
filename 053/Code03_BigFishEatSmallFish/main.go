package main

// 在一步操作中，移除所有满足 nums[i - 1] > nums[i] 的 nums[i] ，其中 0 < i < nums.length 。
// 重复执行步骤，直到 nums 变为 非递减 数组，返回所需执行的操作数。
// 算不算找到右侧首个不小于自己的数？只需要求出最大值即可
// 其实是找到一个数， 左侧的最大值
// 测试链接 : https://leetcode.cn/problems/steps-to-make-array-non-decreasing/

// 从右往左便利小压大，弹出记录大于信息
func totalSteps(nums []int) int {
	stack := make([]int, len(nums))
	n := make([]int, len(nums)) //轮数
	r := 0
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		for r > 0 && nums[i] > nums[stack[r-1]] {
			r--
			n[i] = max(n[i]+1, n[stack[r]])
		}
		res = max(res, n[i])
		stack[r] = i
		r++
	}
	return res
}
func main() {
	totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11})
}
