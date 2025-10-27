package main

// 打砖块
// 测试链接 : https://leetcode.cn/problems/bricks-falling-when-hit/
// 1. 炮位置--
// 2. 洪水填充天花板
// 3. 时光倒流处理炮弹
func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 2
	return dfs(grid, i-1, j) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i, j+1) + 1
}

func hitBricks(grid [][]int, hits [][]int) []int {
	// 炮弹位置先减少：所有格子<=1
	for _, index := range hits {
		grid[index[0]][index[1]]--
	}
	// 把和天花板相连的全部1都涂成2
	for j := range grid[0] {
		dfs(grid, 0, j)
	}
	// 判断是否和2相连，若相连，就将1涂成2，1的个数是砖块数量
	res := make([]int, len(hits))
	for index := len(hits) - 1; index >= 0; index-- {
		i, j := hits[index][0], hits[index][1]
		grid[i][j]++
		if grid[i][j] == 1 {
			// 遇到2||在天花板
			if i == 0 ||
				(i > 0 && grid[i-1][j] == 2) ||
				(j > 0 && grid[i][j-1] == 2) ||
				(i < len(grid)-1 && grid[i+1][j] == 2) ||
				(j < len(grid[0])-1) && grid[i][j+1] == 2 {
				res[index] = dfs(grid, i, j) - 1
			}
		}
	}
	return res
}
