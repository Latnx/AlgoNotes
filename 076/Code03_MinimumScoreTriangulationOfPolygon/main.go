package main

import "math"

// 多边形三角剖分的最低得分
// 你有一个凸的 n 边形，其每个顶点都有一个整数值
// 给定一个整数数组values，其中values[i]是第i个顶点的值(顺时针顺序)
// 假设将多边形 剖分 为 n - 2 个三角形
// 对于每个三角形，该三角形的值是顶点标记的乘积
// 三角剖分的分数是进行三角剖分后所有 n - 2 个三角形的值之和
// 返回 多边形进行三角剖分后可以得到的最低分
// 测试链接 : https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/
func minScoreTriangulation(values []int) int {
	return f2(values)
}
func f1(values []int, l, r int) int {
	if l == r || l+1 == r {
		return 0
	}
	res := math.MaxInt
	for k := l + 1; k < r; k++ {
		res = min(res, values[l]*values[r]*values[k]+f1(values, l, k)+f1(values, k, r))
	}
	return res
}
func f2(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}
	for l := n - 3; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			res := math.MaxInt
			for k := l + 1; k < r; k++ {
				res = min(res, values[l]*values[r]*values[k]+dp[l][k]+dp[k][r])
			}
			dp[l][r] = res
		}
	}
	return dp[0][n-1]
}
