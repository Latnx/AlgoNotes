package main

// 矩阵中的最长递增路径
// 测试链接 : https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/
var dp [201][201]int
var mv = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func longestIncreasingPath(matrix [][]int) int {
	ans := 0
	for i := range len(matrix) + 1 {
		for j := range len(matrix[i]) + 1 {
			dp[i][j] = -1
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			ans = max(ans, f2(matrix, i, j))
		}
	}
	return ans
}

func f1(matrix [][]int, i, j int) int {
	n, m := len(matrix), len(matrix[0])
	ans := 1
	for _, to := range mv {
		x, y := i+to[0], j+to[1]
		if x >= 0 && y >= 0 && x < n && y < m && matrix[i][j] < matrix[x][y] {
			temp := matrix[i][j]
			matrix[i][j] = -1
			ans = max(ans, f1(matrix, x, y)+1)
			matrix[i][j] = temp
		}
	}
	return ans
}
func f2(matrix [][]int, i, j int) int {
	n, m := len(matrix), len(matrix[0])
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	ans := 1
	for _, to := range mv {
		x, y := i+to[0], j+to[1]
		if x >= 0 && y >= 0 && x < n && y < m && matrix[i][j] < matrix[x][y] {
			ans = max(ans, f2(matrix, x, y)+1)
		}
	}
	dp[i][j] = ans
	return ans
}
