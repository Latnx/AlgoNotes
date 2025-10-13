package main

// 柱状图中最大的矩形
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度
// 每个柱子彼此相邻，且宽度为 1 。求在该柱状图中，能够勾勒出来的矩形的最大面积
// 测试链接：https://leetcode.cn/problems/largest-rectangle-in-histogram

func largestRectangleArea(heights []int) int {
	stack := make([]int, len(heights))
	r := 0
	lMin := make([]int, len(heights))
	rMin := make([]int, len(heights))
	for i, val := range heights {
		for r > 0 && val <= heights[stack[r-1]] {
			r--
			if r == 0 {
				lMin[stack[r]] = -1
			} else {
				lMin[stack[r]] = stack[r-1]
			}
			rMin[stack[r]] = i
		}
		stack[r] = i
		r++
	}
	for r > 0 {
		r--
		if r == 0 {
			lMin[stack[r]] = -1
		} else {
			lMin[stack[r]] = stack[r-1]
		}
		rMin[stack[r]] = len(heights)
	}
	for i := len(stack) - 2; i >= 0; i-- {
		if stack[i] != -1 && heights[i] == heights[stack[i]] {
			stack[i] = stack[stack[i]]
		}
	}
	res := 0
	for i := range heights {
		res = max(res, heights[i]*(rMin[i]-lMin[i]-1))
	}
	return res
}
