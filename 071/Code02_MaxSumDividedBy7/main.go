package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 子序列累加和必须被7整除的最大累加和
// 给定一个非负数组nums，
// 可以任意选择数字组成子序列，但是子序列的累加和必须被7整除
// 返回最大累加和
// 对数器验证
func maxSum1(nums []int) int {
	return f(nums, 0, 0)
}
func f(nums []int, index, sum int) int {
	if len(nums) == index {
		if sum%7 == 0 {
			return sum
		}
		return 0
	}
	return max(f(nums, index+1, sum), f(nums, index+1, sum+nums[index]))
}

// dp[i][j], 考虑前i个数字, 取模等于j，最大子序列的和是多少
func maxSum2(nums []int) int {
	n := len(nums)
	var dp [][]int = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 7)
	}
	dp[0][0] = 0
	for j := 1; j < 7; j++ {
		dp[0][j] = -1
	}
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		cur := nums[i-1] % 7
		for j := range dp[i] {
			need := (j - cur + 7) % 7
			if dp[i-1][need] != -1 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][need]+x)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][0]
}

func randomArray(n, v int) []int {
	// 创建一个新的随机数生成器，使用当前时间作为种子
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, n)
	for i := range n {
		arr[i] = rng.Intn(v) // 生成一个 [0, v) 范围内的随机整数
	}
	return arr
}
func main() {
	nums := randomArray(10, 100)
	fmt.Print(nums)
	fmt.Print(maxSum1(nums), maxSum2(nums))
}
