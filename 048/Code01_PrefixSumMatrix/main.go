package main

// 利用二维前缀和信息迅速得到二维区域和
// 测试链接 : https://leetcode.cn/problems/range-sum-query-2d-immutable/
// 计算前缀: 自己 + 左 + 上 - 左上
type NumMatrix struct {
	preMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
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
	return NumMatrix{preMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row2++
	col2++
	return this.preMatrix[row2][col2] - this.preMatrix[row1][col2] - this.preMatrix[row2][col1] + this.preMatrix[row1][col1]
}
