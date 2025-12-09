package main

// 盈利计划(多维费用背包)
// 集团里有 n 名员工，他们可以完成各种各样的工作创造利润
// 第 i 种工作会产生 profit[i] 的利润，它要求 group[i] 名成员共同参与
// 如果成员参与了其中一项工作，就不能参与另一项工作
// 工作的任何至少产生 minProfit 利润的子集称为 盈利计划
// 并且工作的成员总数最多为 n
// 有多少种计划可以选择？因为答案很大，答案对 1000000007 取模
// 测试链接 : https://leetcode.cn/problems/profitable-schemes/
const MOD = 1000000007

var dp [101][101]int

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	for i := range 101 {
		for j := range 101 {
			dp[i][j] = 0
		}
	}
	return f2(group, profit, n, minProfit)
}

// // 普通dfs
// func dfs(group []int, profit []int, n int, minProfit int, index int) int {
// 	if n < 0 {
// 		return 0
// 	}
// 	if index == len(group) || n == 0 {
// 		if minProfit <= 0 {
// 			return 1
// 		}
// 		return 0
// 	}
// 	res := 0
// 	res += dfs(group, profit, n-group[index], minProfit-profit[index], index+1)
// 	res += dfs(group, profit, n, minProfit, index+1)
// 	return res
// }

// // 加缓存, 需要处理minprofit 《 0 的情况，因为在这种情况下，和等于0相同，为了简化处理，所有小于零就等于零
// func dfsdp(group []int, profit []int, n int, minProfit int, index int) int {
// 	if n < 0 {
// 		return 0
// 	}
// 	if index == len(group) || n == 0 {
// 		if minProfit <= 0 {
// 			return 1
// 		}
// 		return 0
// 	}
// 	if dp[index][n][minProfit] != -1 {
// 		return dp[index][n][minProfit]
// 	}
// 	res := 0
// 	res += dfsdp(group, profit, n-group[index], max(0, minProfit-profit[index]), index+1) % MOD
// 	res += dfsdp(group, profit, n, minProfit, index+1) % MOD
// 	dp[index][n][minProfit] = res % MOD
// 	return res % MOD
// }

//	func f1(group []int, profit []int, n int, minProfit int) int {
//		for i := len(group) n + 1 {
//			dp[0][i][0] = 1
//		}
//		for l := range len(group) + 1 {
//			dp[l][0][0] = 1
//		}
//		for l := len(group)-1; l >= 1; l-- {
//			for i := 0; i <= n; i++ {
//				for j := 0; j <= minProfit; j++ {
//					dp[l][i][j] += dp[l-1][i][j]
//					if i-group[l-1] >= 0 {
//						dp[l][i][j] += dp[l-1][i-group[l-1]][max(j-profit[l-1], 0)]
//					}
//				}
//			}
//		}
//		return dp[len(group)][n][minProfit]
//	}
func f2(group []int, profit []int, n int, minProfit int) int {
	for i := range n + 1 {
		dp[i][0] = 1
	}
	for l := len(group); l >= 1; l-- {
		for i := n; i >= 0; i-- {
			for j := minProfit; j >= 0; j-- {
				if i-group[l-1] >= 0 {
					dp[i][j] = (dp[i][j] + dp[i-group[l-1]][max(j-profit[l-1], 0)]) % MOD
				}
			}
		}
	}
	return dp[n][minProfit]
}
