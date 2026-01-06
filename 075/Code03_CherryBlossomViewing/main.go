package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// 观赏樱花
// 给定一个背包的容量t，一共有n种货物，并且给定每种货物的信息
// 花费(cost)、价值(val)、数量(cnt)
// 如果cnt == 0，代表这种货物可以无限选择
// 如果cnt > 0，那么cnt代表这种货物的数量
// 挑选货物的总容量在不超过t的情况下，返回能得到的最大价值
// 背包容量不超过1000，每一种货物的花费都>0
// 测试链接 : https://www.luogu.com.cn/problem/P1833
func toMinute(s string) int {
	parts := strings.Split(s, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return h*60 + m
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	start := toMinute(scanner.Text())
	scanner.Scan()
	end := toMinute(scanner.Text())
	t := end - start
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	v, w, c := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		v[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		w[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
	}
	writer.WriteString(strconv.Itoa(compute(n, t, w, v, c)))
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
			if c[i] == 0 && j-w[i] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-w[i]]+v[i])
			}
			for k := 1; k <= c[i] && j-k*w[i] >= 0; k++ {
				dp[i][j] = max(dp[i][j], dp[i-1][j-k*w[i]]+k*v[i])
			}
		}
	}
	return dp[n][t]
}
