package main

// 最大矩形
// 给定一个仅包含 0 和 1 、大小为 rows * cols 的二维二进制矩阵
// 找出只包含 1 的最大矩形，并返回其面积
// 测试链接：https://leetcode.cn/problems/maximal-rectangle/
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

func maximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0]))
	res := 0
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		res = max(largestRectangleArea(heights), res)
	}
	return res
}
