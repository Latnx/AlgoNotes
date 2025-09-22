package main

// 边框为1的最大正方形
// 给你一个由若干 0 和 1 组成的二维网格 grid
// 请你找出边界全部由 1 组成的最大 正方形 子网格
// 并返回该子网格中的元素数量。如果不存在，则返回 0。
// 测试链接 : https://leetcode.cn/problems/largest-1-bordered-square/
// 枚举所有正方形,没法优化
// 只能优化边框都为1
func Constructor(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	preMatrix := make([][]int, m+1)
	preMatrix[0] = make([]int, n+1)
	for i := range m {
		preMatrix[i+1] = make([]int, n+1)
		for j := range n {
			preMatrix[i+1][j+1] = matrix[i][j]
		}
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			preMatrix[i][j] = preMatrix[i-1][j] + preMatrix[i][j-1] - preMatrix[i-1][j-1] + preMatrix[i][j]
		}
	}
	return preMatrix
}

func SumRegion(preMatrix [][]int, row1 int, col1 int, row2 int, col2 int) int {
	row2++
	col2++
	return preMatrix[row2][col2] - preMatrix[row1][col2] - preMatrix[row2][col1] + preMatrix[row1][col1]
}
func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	preMatrix := Constructor(grid)
	res := 0
	for i := range m {
		for j := range n {
			for k := range min(i, j) + 1 {
				// k是向左上的长度
				if k == 0 || i == 0 || j == 0 {
					if grid[i][j] == 1 {
						res = max(res, 1)
					}
					continue
				}
				if SumRegion(preMatrix, i-k, j-k, i, j)-SumRegion(preMatrix, i-k+1, j-k+1, i-1, j-1) == k*4 {
					res = max(res, (k+1)*(k+1))
				}
			}
		}
	}
	return res
}
