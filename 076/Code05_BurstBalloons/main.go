package main

// 戳气球
// 有 n 个气球，编号为0到n-1，每个气球上都标有一个数字，这些数字存在数组nums中
// 现在要求你戳破所有的气球。戳破第 i 个气球
// 你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币
// 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号
// 如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球
// 求所能获得硬币的最大数量
// 测试链接 : https://leetcode.cn/problems/burst-balloons/
func maxCoins(nums []int) int {
	temp := make([]int, 0, len(nums)+2)
	temp = append(temp, 1)
	temp = append(temp, nums...)
	temp = append(temp, 1)
	dp = make([][]int, len(temp))
	for i := range dp {
		dp[i] = make([]int, len(temp))
	}
	return f2(temp)
}
func f1(nums []int, l, r int) int {
	if l+1 == r || l >= r {
		return 0
	}
	res := 0
	for k := l + 1; k < r; k++ {
		res = max(res, f1(nums, l, k)+f1(nums, k, r)+nums[l]*nums[k]*nums[r])
	}
	return res
}

var dp [][]int

func f2(nums []int) int {
	n := len(nums)
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			cost := 0
			for k := i + 1; k < j; k++ {
				cost = max(cost, dp[i][k]+dp[k][j]+(nums[i]*nums[k]*nums[j]))
			}
			dp[i][j] = cost
		}
	}
	return dp[0][n-1]
}
