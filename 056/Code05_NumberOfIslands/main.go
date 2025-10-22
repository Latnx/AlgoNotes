package main

import "fmt"

// 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成
// 此外，你可以假设该网格的四条边均被水包围
// 测试链接 : https://leetcode.cn/problems/number-of-islands/

// 将二维坐标转换为一维坐标
var arr []int
var sets int

func find(i int) int {
	if arr[i] != i {
		arr[i] = find(arr[i])
	}
	return arr[i]
}

func union(x, y int) {
	findX, findY := find(x), find(y)
	if findX != findY {
		sets--
	}
	arr[findX] = findY
}

func numIslands(grid [][]byte) int {
	arr = make([]int, len(grid)*len(grid[0]))
	for i := range arr {
		arr[i] = i
	}
	sets = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 如果左面或者上面有1
			if grid[i][j] == '1' {
				sets++
				if j > 0 && grid[i][j-1] == '1' {
					union(i*len(grid[0])+j, i*len(grid[0])+j-1)
				}
				if i > 0 && grid[i-1][j] == '1' {
					union(i*len(grid[0])+j, (i-1)*len(grid[0])+j)
				}
			}
		}
	}
	return sets
}

func main() {
	grid := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}

	result := numIslands(grid)
	fmt.Println("岛屿数量:", result)
}
