package main

// 被围绕的区域
// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域
// 并将这些区域里所有的 'O' 用 'X' 填充。
// 测试链接 : https://leetcode.cn/problems/surrounded-regions/
func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'F'
	dfs(board, i, j-1)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i+1, j)
}

func solve(board [][]byte) {
	// 先遍历四周，沿着四周生长，将O标记成F，再将剩余O标记为X, 再将F标记回O
	for j := range len(board[0]) {
		dfs(board, 0, j)
	}
	for j := range len(board[0]) {
		dfs(board, len(board)-1, j)
	}
	for i := range len(board) {
		dfs(board, i, 0)
	}
	for i := range len(board) {
		dfs(board, i, len(board[0])-1)
	}
	for i := range board {
		for j := range board[0] {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'F':
				board[i][j] = 'O'
			}
		}
	}

}
