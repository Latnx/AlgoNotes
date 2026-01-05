package main

import (
	"bufio"
	"os"
	"strconv"
)

// 完全背包(模版)
// 给定一个正数t，表示背包的容量
// 有m种货物，每种货物可以选择任意个
// 每种货物都有体积costs[i]和价值values[i]
// 返回在不超过总容量的情况下，怎么挑选货物能达到价值最大
// 返回最大的价值
// 测试链接 : https://www.luogu.com.cn/problem/P1616

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	costs := make([]int, m+1)
	values := make([]int, m+1)
	for i := 1; i <= m; i++ {
		scanner.Scan()
		costs[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		values[i], _ = strconv.Atoi(scanner.Text())
	}
	writer.WriteString(strconv.Itoa(compute1(t, m, costs, values)))
}
func compute1(t, m int, costs, values []int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= t; j++ {
			dp[i][j] = dp[i-1][j]
			if j-costs[i] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-costs[i]]+values[i])
			}
		}
	}
	return dp[m][t]
}
func compute2(t, m int, costs, values []int) int {
	dp := make([]int, t+1)
	for i := 1; i <= m; i++ {
		for j := costs[i]; j <= t; j++ {
			dp[j] = max(dp[j], dp[j-costs[i]]+values[i])
		}
	}
	return dp[t]
}
