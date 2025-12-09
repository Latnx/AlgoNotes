package main

// 矩阵中和能被 K 整除的路径

// 测试链接 : https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/
const MOD = 1000000007

var dp [][][]int

func numberOfPaths(grid [][]int, k int) int {
	dp = make([][][]int, len(grid))
	for i := range dp {
		dp[i] = make([][]int, len(grid[0]))
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
			grid[i][j] %= k
		}
	}
	return f1(grid, k)
}

func dfs(grid [][]int, k, sum int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}
	if i == m-1 && j == n-1 && (sum+grid[i][j])%k == 0 {
		return 1
	}
	if dp[sum][i][j] != -1 {
		return dp[sum][i][j]
	}
	res := 0
	res = (res + dfs(grid, k, (sum+grid[i][j])%k, i+1, j)) % MOD
	res = (res + dfs(grid, k, (sum+grid[i][j])%k, i, j+1)) % MOD
	dp[sum][i][j] = res
	return res
}

// dp表示，从i,j点，总和为k时，
func f1(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp[m-1][n-1][grid[m-1][n-1]%k] = 1

	// 从右下向左上填表
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue // 终点已初始化
			}
			for sum := range k {
				// 当前格子值
				val := grid[i][j]
				// 进入当前格子前的sum，加上当前格子后变成sum
				// 即 (prevSum + val) % k == sum
				// 推导：prevSum = (sum - val%k + k) % k
				prevSum := (sum - val + k) % k
				if j+1 < n {
					dp[i][j][sum] = (dp[i][j][sum] + dp[i][j+1][prevSum]) % MOD
				}
				if i+1 < m {
					dp[i][j][sum] = (dp[i][j][sum] + dp[i+1][j][prevSum]) % MOD
				}
			}
		}
	}
	return dp[0][0][0]
}
func main() {
	numberOfPaths([][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3)
}
