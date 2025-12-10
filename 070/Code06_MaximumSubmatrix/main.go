package main

// 子矩阵最大累加和问题
// 给定一个二维数组grid，找到其中子矩阵的最大累加和
// 返回拥有最大累加和的子矩阵左上角和右下角坐标
// 如果有多个子矩阵都有最大累加和，返回哪一个都可以
// 测试链接 : https://leetcode.cn/problems/max-submatrix-lcci/

// 1. 先压缩，计算0到每一行的和
// 2. 求每一行的最大子数组，即为0到这一行的最大子矩阵
// 3. 重复计算，1到每一行，2到每一行，再求最大累加和
func getMaxMatrix(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	res := make([]int, 4)
	maxRes := matrix[0][0]
	for start := range len(matrix) {
		for j := range len(matrix[0]) {
			dp[start][j] = matrix[start][j]
		}
		for i := start + 1; i < n; i++ {
			for j := 0; j < m; j++ {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			}
		}
		for i := start; i < n; i++ {
			maxL := 0
			for j := 1; j < m; j++ {
				if dp[i][j-1] > 0 {
					dp[i][j] += dp[i][j-1]
				} else {
					// dp[i][j] = dp[i][j-1]
					maxL = j
				}
				if dp[i][j] > maxRes {
					maxRes = dp[i][j]
					res[0], res[1], res[2], res[3] = start, maxL, i, j
				}
			}
		}
	}
	return res
}
