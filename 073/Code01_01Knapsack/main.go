package main

import (
	"bufio"
	"os"
	"strconv"
)

// 01背包(模版)
// 给定一个正数t，表示背包的容量
// 有m个货物，每个货物可以选择一次
// 每个货物有自己的体积costs[i]和价值values[i]
// 返回在不超过总容量的情况下，怎么挑选货物能达到价值最大
// 返回最大的价值
// 测试链接 : https://www.luogu.com.cn/problem/P1048

// 一号物品有一个(w, v),重量和价值
// dp[i][j]: 前i个物品，总重量<=j，能获得的最大价值
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	w := make([]int, M+1)
	v := make([]int, M+1)
	for i := range M {
		scanner.Scan()
		w[i+1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		v[i+1], _ = strconv.Atoi(scanner.Text())
	}
	dp := make([][]int, M+1)
	for i := range dp {
		dp[i] = make([]int, T+1)
	}
	for i := 1; i <= M; i++ {
		for j := 1; j <= T; j++ {
			if j >= w[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	writer.WriteString(strconv.Itoa(dp[M][T]))
}
