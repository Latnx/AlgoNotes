package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// 分组背包(模版)
// 给定一个正数m表示背包的容量，有n个货物可供挑选
// 每个货物有自己的体积(容量消耗)、价值(获得收益)、组号(分组)
// 同一个组的物品只能挑选1件，所有挑选物品的体积总和不能超过背包容量
// 怎么挑选货物能达到价值最大，返回最大的价值
// 测试链接 : https://www.luogu.com.cn/problem/P1757
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	objs := make([][3]int, n+1)

	// 读取物品信息
	for i := 1; i <= n; i++ {
		scanner.Scan()
		objs[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		objs[i][1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		objs[i][2], _ = strconv.Atoi(scanner.Text())
	}
	sort.Slice(objs, func(i, j int) bool { return objs[i][2] < objs[j][2] })
	fmt.Print(compute1(m, n, objs))
}
func compute1(m, n int, objs [][3]int) int {
	// 记录分组数
	teams := 1
	for i := 2; i <= n; i++ {
		if objs[i-1][2] != objs[i][2] {
			teams++
		}
	}
	// 按照分组创建dp
	dp := make([][]int, teams+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	// 每个分组依次尝试
	start, end := 1, 1
	for i := 1; i <= teams; i++ {
		// 记录分组开始和结束的位置，不包含end
		for end <= n && objs[end][2] == objs[start][2] {
			end++
		}
		for j := 0; j <= m; j++ {
			// 一个分组不选物品
			dp[i][j] = dp[i-1][j]
			// 逐个逐个尝试一个分组中的每一个物品
			for idx := start; idx < end; idx++ {
				if j-objs[idx][0] >= 0 {
					dp[i][j] = max(dp[i][j], dp[i-1][j-objs[idx][0]]+objs[idx][1])
				}
			}
		}
		start = end
	}
	return dp[teams][m]
}
