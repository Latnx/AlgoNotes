package main

// 骑士在棋盘上的概率
// 测试链接 : https://leetcode.cn/problems/knight-probability-in-chessboard/
var dp [101][25][25]float64

func knightProbability(n int, k int, row int, column int) float64 {
	for l := range k + 1 {
		for i := range n {
			for j := range n {
				dp[l][i][j] = -1
			}
		}
	}
	return dfs(n, k, row, column)
}

var move = [8][2]int{{-1, -2}, {-2, -1}, {1, -2}, {2, -1}, {-1, 2}, {-2, 1}, {1, 2}, {2, 1}}

func dfs(n, k, row, column int) (yes float64) {
	if row < 0 || column < 0 || row >= n || column >= n {
		return 0
	}
	if k == 0 {
		return 1
	}
	if dp[k][row][column] != -1 {
		return dp[k][row][column]
	}
	for _, to := range move {
		yes += dfs(n, k-1, row+to[0], column+to[1]) / 8
	}
	dp[k][row][column] = yes
	return yes
}
func main() {
	knightProbability(8, 30, 6, 4)
}
