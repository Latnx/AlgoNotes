package main

// 滑动窗口最大值（单调队列经典用法模版）
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧
// 你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回 滑动窗口中的最大值 。
// 测试链接 : https://leetcode.cn/problems/sliding-window-maximum/

// 单调队列通常与滑动窗口连用, 用来求滑动窗口的最大和最小值
// 流程分为加数字和减数字两部分
// 1. 加数字从队尾进入，若是队尾大于当前数字的元素，就从队尾弹出
// 2. 减数字时，若队头过期则循环出队列

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, len(nums))
	l, r := 0, 0
	// 处理排序
	add := func(index int) {
		for l != r && nums[queue[r-1]] <= nums[index] {
			r--
		}
		queue[r] = index
		r++
	}
	// 处理过期
	del := func(index int) {
		for l != r && queue[l] <= index {
			l++
		}
	}
	for i := range k {
		add(i)
	}
	res := make([]int, len(nums)-k+1)
	res[0] = nums[queue[l]]
	for i := 0; i+k < len(nums); i++ {
		del(i)
		add(i + k)
		res[i+1] = nums[queue[l]]
	}
	return res
}

func main() {
	maxSlidingWindow([]int{1, -1}, 1)
}
