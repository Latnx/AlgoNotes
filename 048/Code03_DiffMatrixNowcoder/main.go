package main

import (
	"fmt"
)

// 二维差分模版(牛客)
// 测试链接 : https://www.nowcoder.com/practice/50e1a93989df42efb0b1dec386fb4ccc

var grid [][]int
var varGrid [][]int

func set(a, b, c, d, num int) {
	varGrid[a][b] += num
	varGrid[c+1][b] -= num
	varGrid[a][d+1] -= num
	varGrid[c+1][d+1] += num
}
func build() {
	for i := 1; i < len(varGrid); i++ {
		for j := 1; j < len(varGrid[0]); j++ {
			varGrid[i][j] += varGrid[i-1][j] + varGrid[i][j-1] - varGrid[i-1][j-1]
		}
	}
	for i := range grid {
		for j := range grid[0] {
			grid[i][j] += varGrid[i+1][j+1]
		}
	}
}

func main() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	grid = make([][]int, n)
	varGrid = make([][]int, n+2)
	varGrid[0] = make([]int, m+2)
	varGrid[n+1] = make([]int, m+2)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		varGrid[i+1] = make([]int, m+2)
		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	for i := 0; i < q; i++ {
		var a, b, c, d, num int
		fmt.Scan(&a, &b, &c, &d, &num)
		set(a, b, c, d, num)
	}
	build()
	for i := range grid {
		for j := range grid[0] {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}
