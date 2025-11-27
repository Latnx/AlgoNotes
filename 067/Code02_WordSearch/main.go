package main

// 单词搜索（无法改成动态规划）
// 测试链接 : https://leetcode.cn/problems/word-search/

// 决定返回值的可变参数往往都比较简单，不会比int更复杂
// var visited [][]int
var move = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j int, index int) bool {
	m, n := len(board), len(board[0])
	if index == len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[index] {
		return false
	}
	temp := board[i][j]
	board[i][j] = 0
	for _, to := range move {
		if dfs(board, word, i+to[0], j+to[1], index+1) {
			return true
		}
	}
	board[i][j] = temp
	return false
}
