package main

// 获取所有钥匙的最短路径
// 测试链接：https://leetcode.cn/problems/shortest-path-to-get-all-keys

// 2^6种状态
var visited [64][][]bool
var move [][2]int = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	key := 0
	for k := range visited {
		visited[k] = make([][]bool, m)
		for i := range m {
			visited[k][i] = make([]bool, n)
		}
	}
	queue := make([][3]int, 0)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				queue = append(queue, [3]int{i, j, 0})
				visited[0][i][j] = true
			}
			if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				key |= 1 << (grid[i][j] - 'a')
			}
		}
	}
	step := 0
	for len(queue) > 0 {
		queueL := len(queue)
		for range queueL {
			node := queue[0]
			queue = queue[1:]
			x, y, nowKey := node[0], node[1], node[2]
			for _, to := range move {
				toX, toY := x+to[0], y+to[1]
				if toX >= 0 && toY >= 0 && toX < m && toY < n && !visited[nowKey][toX][toY] {
					visited[nowKey][toX][toY] = true
					newKey := nowKey
					cell := grid[toX][toY]
					if cell >= 'a' && cell <= 'f' {
						newKey |= 1 << (cell - 'a')
					} else if cell >= 'A' && cell <= 'F' && nowKey&(1<<(cell-'A')) == 0 {
						continue
					} else if cell == '#' {
						continue
					}
					queue = append(queue, [3]int{toX, toY, newKey})
					if newKey == key {
						return step + 1
					}
				}
			}
		}
		step++
	}
	return -1
}
