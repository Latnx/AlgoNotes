package main

// 一和零(多维费用背包)
// 给你一个二进制字符串数组 strs 和两个整数 m 和 n
// 请你找出并返回 strs 的最大子集的长度
// 该子集中 最多 有 m 个 0 和 n 个 1
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集
// 测试链接 : https://leetcode.cn/problems/ones-and-zeroes/
var dp [101][101]int

func findMaxForm(strs []string, m int, n int) int {
	for i := range 101 {
		for j := range 101 {
			dp[i][j] = 0
		}
	}
	return f3(strs, m, n)
}
func f3(strs []string, m int, n int) int {
	// 遍历每个字符串
	for k := 1; k <= len(strs); k++ {
		count0, count1 := 0, 0
		for _, ch := range strs[k-1] {
			if ch == '0' {
				count0++
			} else {
				count1++
			}
		}
		for i := m; i >= count0; i-- {
			for j := n; j >= count1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-count0][j-count1]+1)
			}
		}
	}
	return dp[m][n]
}

//	func dfs(strs []string, m int, n int, index int) int {
//		if index == len(strs) {
//			return 0
//		}
//		if dp[m][n][index] != -1 {
//			return dp[m][n][index]
//		}
//		count0, count1 := 0, 0
//		for _, ch := range strs[index] {
//			if ch == '0' {
//				count0++
//			} else {
//				count1++
//			}
//		}
//		if m-count0 >= 0 && n-count1 >= 0 {
//			dp[m][n][index] = max(dfs(strs, m-count0, n-count1, index+1) + 1)
//		}
//		dp[m][n][index] = max(dp[m][n][index], dfs(strs, m, n, index+1))
//		return dp[m][n][index]
//	}
//
//	func f2(strs []string, m int, n int) int {
//		// 遍历每个字符串
//		for k := 1; k <= len(strs); k++ {
//			count0, count1 := 0, 0
//			for _, ch := range strs[k-1] {
//				if ch == '0' {
//					count0++
//				} else {
//					count1++
//				}
//			}
//			for i := 0; i <= m; i++ {
//				for j := 0; j <= n; j++ {
//					dp[i][j][k] = dp[i][j][k-1]
//					if i-count0 >= 0 && j-count1 >= 0 {
//						dp[i][j][k] = max(dp[i][j][k], dp[i-count0][j-count1][k-1]+1)
//					}
//				}
//			}
//		}
//		return dp[m][n][len(strs)]
//	}
