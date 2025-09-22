package main

import "fmt"

// 用邮票贴满网格图
// 给你一个 m * n 的二进制矩阵 grid
// 每个格子要么为 0 （空）要么为 1 （被占据）
// 给你邮票的尺寸为 stampHeight * stampWidth
// 我们想将邮票贴进二进制矩阵中，且满足以下 限制 和 要求 ：
// 覆盖所有空格子，不覆盖任何被占据的格子
// 可以放入任意数目的邮票，邮票可以相互有重叠部分
// 邮票不允许旋转，邮票必须完全在矩阵内
// 如果在满足上述要求的前提下，可以放入邮票，请返回 true ，否则返回 false
// 测试链接 : https://leetcode.cn/problems/stamping-the-grid/

// 为了判断区域内有没有空, 使用前缀:若区域内值为零则可贴
// 快速设置一个区域的值, 使用差分
var diff [][]int
var preSum [][]int
var m, n int

// 差分
func set(row1, col1, row2, col2 int) {
	diff[row1][col1] += 1
	diff[row1][col2+1] -= 1
	diff[row2+1][col1] -= 1
	diff[row2+1][col2+1] += 1
}

func build() {
	for i := range m {
		for j := range n {
			diff[i+1][j+1] += -diff[i][j] + diff[i][j+1] + diff[i+1][j]
		}
	}
}

// 前缀和
func sumRegion(row1, col1, row2, col2 int) int {
	return preSum[row2][col2] - preSum[row1-1][col2] - preSum[row2][col1-1] + preSum[row1-1][col1-1]
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n = len(grid), len(grid[0])

	// 前缀和
	preSum = make([][]int, m+2)
	preSum[0] = make([]int, n+2)
	preSum[m+1] = make([]int, n+2)
	for i := range m {
		preSum[i+1] = make([]int, n+2)
		for j := range n {
			preSum[i+1][j+1] = grid[i][j] + preSum[i][j+1] + preSum[i+1][j] - preSum[i][j]
		}
	}
	// 差分
	diff = make([][]int, m+2)
	for i := range m + 2 {
		diff[i] = make([]int, n+2)
	}

	// 主逻辑
	for i := range m {
		for j := range n {
			if grid[i][j] == 0 && i+stampHeight < m+1 && j+stampWidth < n+1 && sumRegion(i+1, j+1, i+stampHeight, j+stampWidth) == 0 {
				set(i+1, j+1, i+stampHeight, j+stampWidth)
			}
		}
	}
	build()
	for i := range m {
		for j := range n {
			if grid[i][j] == 0 && diff[i+1][j+1] == 0 {
				return false
			}
		}
	}
	return true
}
func main() {
	grid := [][]int{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	stampHeight := 4
	stampWidth := 3
	result := possibleToStamp(grid, stampHeight, stampWidth)
	fmt.Println(diff)
	fmt.Println("结果:", result)
}
