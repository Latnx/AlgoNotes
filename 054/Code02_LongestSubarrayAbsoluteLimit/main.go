package main

// 绝对差不超过限制的最长连续子数组
// 给你一个整数数组 nums ，和一个表示限制的整数 limit
// 请你返回最长连续子数组的长度
// 该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit
// 如果不存在满足条件的子数组，则返回 0
// 测试链接 : https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

// 求一个长度给定长度的最小绝对差

func longestSubarray(nums []int, limit int) int {
	queueMin := make([]int, len(nums))
	queueMax := make([]int, len(nums))
	lMin, rMin := 0, 0
	lMax, rMax := 0, 0
	addMax := func(index int) {
		for lMax < rMax && nums[queueMax[rMax-1]] <= nums[index] {
			rMax--
		}
		queueMax[rMax] = index
		rMax++
	}
	delMax := func(index int) {
		for lMax < rMax && queueMax[lMax] <= index {
			lMax++
		}
	}
	addMin := func(index int) {
		for lMin < rMin && nums[queueMin[rMin-1]] >= nums[index] {
			rMin--
		}
		queueMin[rMin] = index
		rMin++
	}
	delMin := func(index int) {
		for lMin < rMin && queueMin[lMin] <= index {
			lMin++
		}
	}
	l, res := 0, 0
	for r := range nums {
		addMax(r)
		addMin(r)
		for nums[queueMax[lMax]]-nums[queueMin[lMin]] > limit {
			delMax(l)
			delMin(l)
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
func main() {
	longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5)
}
