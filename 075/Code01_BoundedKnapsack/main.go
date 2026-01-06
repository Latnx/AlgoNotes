package main

import (
	"bufio"
	"os"
	"strconv"
)

// 多重背包不进行枚举优化
// 宝物筛选
// 一共有n种货物, 背包容量为t
// 每种货物的价值(v[i])、重量(w[i])、数量(c[i])都给出
// 请返回选择货物不超过背包容量的情况下，能得到的最大的价值
// 测试链接 : https://www.luogu.com.cn/problem/P1776
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	v, w, c := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		v[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		w[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
	}
	writer.WriteString(strconv.Itoa(compute(n, t, v, w, c)))
}

func compute(n, t int, v, w, c []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= t; j++ {
			dp[i][j] = dp[i-1][j]
			// 遍历要1...c[i]个
			for k := 1; k <= c[i] && j-k*w[i] >= 0; k++ {
				dp[i][j] = max(dp[i][j], dp[i-1][j-k*w[i]]+k*v[i])
			}
		}
	}
	return dp[n][t]
}
