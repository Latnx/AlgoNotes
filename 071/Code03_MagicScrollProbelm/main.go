package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 魔法卷轴
// 给定一个数组nums，其中可能有正、负、0
// 每个魔法卷轴可以把nums中连续的一段全变成0
// 你希望数组整体的累加和尽可能大
// 卷轴使不使用、使用多少随意，但一共只有2个魔法卷轴
// 请返回数组尽可能大的累加和
func maxSum1(nums []int) int {
	p1, p2, p3 := 0, 0, math.MinInt
	// 不用
	for _, num := range nums {
		p1 += num
	}
	// 用一次
	p2 = mustOneScroll(nums, 0, len(nums)-1)
	// 用两次
	for i := 1; i < len(nums); i++ {
		p3 = max(p3, mustOneScroll(nums, 0, i-1), mustOneScroll(nums, i, len(nums)-1))
	}
	return max(p1, p2, p3)
}

func mustOneScroll(nums []int, l, r int) int {
	res := math.MinInt
	for i := l; i <= r; i++ {
		for j := i; j <= r; j++ {
			temp := 0
			for index := l; index <= r; index++ {
				if index >= i && index <= j {
					continue
				}
				temp += nums[index]
			}
			res = max(res, temp)
		}
	}
	return res
}

func maxSum2(nums []int) int {
	n := len(nums)
	p1, p2, p3 := 0, 0, math.MinInt
	// 不用
	for _, num := range nums {
		p1 += num
	}
	// 用一次卷轴, 0-i 一定用一次, 前缀和,最大前缀和
	dp := make([]int, n)
	dp[0] = 0
	sum := nums[0]
	sumMax := max(0, sum)
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], sumMax)
		sum += nums[i]
		sumMax = max(sumMax, sum)
	}
	p2 = dp[n-1]
	// 用两次卷轴,只需要再求n-1 到 i上只是用一次卷轴的最大后缀和
	// 重复上述步骤
	dp2 := make([]int, n)
	dp2[n-1] = 0
	sum2 := nums[n-1]
	sumMax2 := max(0, sum2)
	for i := n - 2; i >= 0; i-- {
		dp2[i] = max(dp2[i+1]+nums[i], sumMax2)
		sum2 += nums[i]
		sumMax2 = max(sumMax2, sum2)
	}
	for i := 1; i < n; i++ {
		p3 = max(p3, dp[i-1]+dp2[i])
	}
	return max(p1, p2, p3)
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
	nums := randomArray(5, 100)
	fmt.Print(nums, maxSum1(nums), maxSum2(nums))
}
