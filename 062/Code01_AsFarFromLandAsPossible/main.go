package main

// 地图分析
// 测试链接 : https://leetcode.cn/problems/as-far-from-land-as-possible/
func maxDistance(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	type point struct{ x, y int }
	queue := make([]point, n*m)
	l, r := 0, 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				queue[r] = point{i, j}
				r++
			}
		}
	}
	if r == 0 || r == n*m {
		return -1
	}
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dist := -1
	for l < r {
		levelCnt := r - l
		for levelCnt > 0 {
			for _, d := range dirs {
				p := queue[l]
				x, y := p.x+d.x, p.y+d.y
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 0 {
					queue[r] = point{x, y}
					grid[x][y] = 1
					r++
				}
			}
			levelCnt--
			l++
		}
		dist++
	}
	return dist - 1
}
