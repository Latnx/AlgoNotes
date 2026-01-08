package main

import (
	"math"
	"sort"
)

// 切棍子的最小成本
// 有一根长度为n个单位的木棍，棍上从0到n标记了若干位置
// 给你一个整数数组cuts，其中cuts[i]表示你需要将棍子切开的位置
// 你可以按顺序完成切割，也可以根据需要更改切割的顺序
// 每次切割的成本都是当前要切割的棍子的长度，切棍子的总成本是历次切割成本的总和
// 对棍子进行切割将会把一根木棍分成两根较小的木棍
// 这两根木棍的长度和就是切割前木棍的长度
// 返回切棍子的最小总成本
// 测试链接 : https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/
func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	tmp := make([]int, 0, len(cuts)+2)
	tmp = append(tmp, 0)
	tmp = append(tmp, cuts...)
	tmp = append(tmp, n)
	cuts = tmp
	dp = make([][]int, len(cuts))
	for i := range dp {
		dp[i] = make([]int, len(cuts))
	}
	return f2(cuts)
}

var dp [][]int

func f1(cuts []int, l, r int) int {
	if l >= r || l+1 == r {
		return 0
	}
	if dp[l][r] != 0 {
		return dp[l][r]
	}
	cost := math.MaxInt
	for i := l + 1; i < r; i++ {
		cost = min(cost, f1(cuts, l, i)+f1(cuts, i, r)+(cuts[r]-cuts[l]))
	}
	if cost == math.MaxInt {
		cost = 0
	}
	dp[l][r] = cost
	return cost
}
func f2(cuts []int) int {
	n := len(cuts)
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			cost := math.MaxInt
			for k := i + 1; k < j; k++ {
				cost = min(cost, dp[i][k]+dp[k][j]+(cuts[j]-cuts[i]))
			}
			if cost == math.MaxInt {
				cost = 0
			}
			dp[i][j] = cost
		}
	}
	return dp[0][n-1]
}
func main() {
	minCost(7, []int{1, 3, 4, 5})
}
