package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 可以翻转1次的情况下子数组最大累加和
// 给定一个数组nums，
// 现在允许你随意选择数组连续一段进行翻转，也就是子数组逆序的调整
// 比如翻转[1,2,3,4,5,6]的[2~4]范围，得到的是[1,2,5,4,3,6]
// 返回必须随意翻转1次之后，子数组的最大累加和
func maxSumReverse1(nums []int) int {
	ans := math.MinInt

	for l := range nums {
		for r := l; r < len(nums); r++ {
			reverse(nums, l, r)
			ans = max(ans, maxSum(nums))
			reverse(nums, l, r)
		}
	}
	return ans
}

// nums[l...r] 范围上的数字进行逆序调整
func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 返回子数组最大累加和（Kadane）
func maxSum(nums []int) int {
	pre := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		pre = max(pre+nums[i], nums[i])
		ans = max(pre, ans)
	}
	return ans
}
func randomArray(n, v int) []int {
	// 创建一个新的随机数生成器，使用当前时间作为种子
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, n)
	for i := range n {
		arr[i] = rng.Intn(v)*2 - v // 生成一个 [0, v) 范围内的随机整数
	}
	return arr
}
func maxSumReverse2(nums []int) int {
	n := len(nums)
	prefix, suffix := make([]int, n), make([]int, n)
	prefix[0] = nums[0]
	suffix[n-1] = nums[n-1]
	// 最大前缀和,后缀和
	for i := 1; i < n; i++ {
		prefix[i] = max(prefix[i-1]+nums[i], nums[i])
	}
	for i := n - 2; i >= 0; i-- {
		suffix[i] = max(suffix[i+1]+nums[i], nums[i])
	}
	// 最大的最大前缀和
	preMax := math.MinInt
	ans := math.MinInt
	for i := range n - 1 {
		preMax = max(preMax, prefix[i])
		ans = max(ans, preMax+suffix[i+1])
	}
	return ans
}
func main() {
	nums := randomArray(10, 100)
	fmt.Print(nums, maxSumReverse1(nums), maxSumReverse2(nums))
}
