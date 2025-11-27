package main

import (
	"bufio"
	"os"
	"strconv"
)

// 节点数为n高度不大于m的二叉树个数
// https://www.nowcoder.com/practice/aaefe5896cce4204b276e213e725f3ea
const MOD = 1000000007

var dp [50][50]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	writer.WriteString(strconv.Itoa(compute1_p(n, m)))
}

func compute1(n, m int) int {
	if n == 0 {
		return 1
	}
	if m == 0 {
		return 0
	}
	ans := 0
	for k := range m + 1 {
		l := compute1(n-1, k)
		r := compute1(n-1, m-k)
		ans += l * r
	}
	return ans
}
func compute1_p(n, m int) int {
	if n == 0 {
		return 1
	}
	if m == 0 {
		return 0
	}
	ans := 0
	for k := 0; k < n; k++ {
		ans = (ans + compute1_p(n-k-1, m-1)%MOD*compute1_p(k, m-1)%MOD) % MOD
	}
	return ans
}
func compute1_dp(n, m int) int {
	if n == 0 {
		return 1
	}
	if m == 0 {
		return 0
	}
	if dp[n][m] != -1 {
		return dp[n][m]
	}
	ans := 0
	for k := 0; k < n; k++ {
		ans = (ans + compute1_dp(n-k-1, m-1)%MOD*compute1_dp(k, m-1)%MOD) % MOD
	}
	dp[n][m] = ans
	return ans
}

// dp表设置为m行，n列，方便压缩（其实没压缩）
func compute2(n, m int) int {
	for i := 0; i <= m; i++ {
		dp[i][0] = 1
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = 0
			for k := 0; k < j; k++ {
				dp[i][j] = (dp[i][j] + (dp[i-1][j-k-1]*dp[i-1][k])%MOD) % MOD
			}
		}
	}
	return dp[m][n]
}
