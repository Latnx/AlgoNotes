package main

import (
	"fmt"
)

// N皇后问题
// 测试链接 : https://leetcode.cn/problems/n-queens-ii/
var qipan [][]bool
var total int

func check(row, col int) bool {
	n := len(qipan)
	for k := range row {
		if qipan[k][col] {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if qipan[i][j] {
			return false
		}
	}
	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if qipan[i][j] {
			return false
		}
	}
	return true
}

func dfs(level int, index int) {
	n := len(qipan)
	if level == n-1 {
		total++
		return
	}
	qipan[level][index] = true
	for j := range n {
		if check(level+1, j) {
			dfs(level+1, j)
		}
	}
	qipan[level][index] = false
}
func totalNQueens(n int) int {
	total = 0
	qipan = make([][]bool, n)
	for k := range qipan {
		qipan[k] = make([]bool, n)
	}
	for k := range n {
		dfs(0, k)
	}
	return total
}
func main() {
	fmt.Print(totalNQueens(4))
}
