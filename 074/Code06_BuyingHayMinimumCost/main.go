package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// 购买足量干草的最小花费
// 有n个提供干草的公司，每个公司都有两个信息
// cost[i]代表购买1次产品需要花的钱
// val[i]代表购买1次产品所获得的干草数量
// 每个公司的产品都可以购买任意次
// 你一定要至少购买h数量的干草，返回最少要花多少钱
// 测试链接 : https://www.luogu.com.cn/problem/P2918
// 干草 = 完全背包 + 最小化 + 至少装满
// 完全采草药 = 完全背包 + 最大化 + 不超过容量
// 跟模板的本质区别就是这个j允许超过，所以dp比原来的大
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	H, _ := strconv.Atoi(scanner.Text())
	costs := make([]int, N+1)
	vals := make([]int, N+1)
	maxVal := 0
	for i := 1; i <= N; i++ {
		scanner.Scan()
		vals[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		costs[i], _ = strconv.Atoi(scanner.Text())
		maxVal = max(maxVal, vals[i])
	}
	writer.WriteString(strconv.Itoa(compute(N, H+maxVal, costs, vals, maxVal)))
}
func compute(n, h int, costs, vals []int, maxVal int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, h+1)

	}
	for j := 1; j <= h; j++ {
		dp[0][j] = math.MaxInt
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= h; j++ {
			dp[i][j] = dp[i-1][j]
			if j-vals[i] >= 0 && dp[i][j-vals[i]] != math.MaxInt {
				dp[i][j] = min(dp[i][j], dp[i][j-vals[i]]+costs[i])
			}
		}
	}
	res := math.MaxInt
	for j := h; maxVal >= 0; maxVal-- {
		res = min(dp[n][j-maxVal], res)
	}
	return res
}
