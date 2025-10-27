package main

// 最大人工岛
// 给你一个大小为 n * n 二进制矩阵 grid 。最多 只能将一格 0 变成 1 。
// 返回执行此操作后，grid 中最大的岛屿面积是多少？
// 岛屿 由一组上、下、左、右四个方向相连的 1 形成
// 测试链接 : https://leetcode.cn/problems/making-a-large-island/
func dfs(grid [][]int, i, j int, id int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = id
	return dfs(grid, i-1, j, id) + dfs(grid, i, j-1, id) + dfs(grid, i+1, j, id) + dfs(grid, i, j+1, id) + 1
}

func largestIsland(grid [][]int) int {
	id := 2
	var m [20000]int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				m[id] = dfs(grid, i, j, id)
				id++
			}
		}
	}
	res := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != 0 {
				res = max(res, m[grid[i][j]])
			} else {
				now := make(map[int]struct{})
				if i > 0 {
					now[grid[i-1][j]] = struct{}{}
				}
				if i < len(grid)-1 {
					now[grid[i+1][j]] = struct{}{}
				}
				if j > 0 {
					now[grid[i][j-1]] = struct{}{}
				}
				if j < len(grid[0])-1 {
					now[grid[i][j+1]] = struct{}{}
				}
				totalVal := 1
				for index := range now {
					totalVal += m[index]
				}
				res = max(res, totalVal)
			}
		}
	}
	return res
}
