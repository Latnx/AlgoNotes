package main

import (
	"math"
	"slices"
)

// 最接近目标值的子序列和
// 测试链接 : https://leetcode.cn/problems/closest-subsequence-sum/
func dfs(index, end int, nums []int, sumArr *[]int, sum int) {
	if index >= end {
		*sumArr = append(*sumArr, sum)
		return
	}
	dfs(index+1, end, nums, sumArr, sum)
	dfs(index+1, end, nums, sumArr, sum+nums[index])

}
func minAbsDifference(nums []int, goal int) int {
	var lsum []int = make([]int, 0)
	var rsum []int = make([]int, 0)
	dfs(0, len(nums)/2, nums, &lsum, 0)
	dfs(len(nums)/2, len(nums), nums, &rsum, 0)
	slices.Sort(lsum)
	slices.Sort(rsum)
	ans := math.MaxInt
	// 找到前一半i=0的最小值
	// 为什么j不回退, i++后lsum[i](相等或变大), 若j取得比现在大(则rsum[j]相等或变大, )
	for i, j := 0, len(rsum)-1; i < len(lsum); i++ {
		for j > 0 && abs(goal-lsum[i]-rsum[j-1]) <= abs(goal-lsum[i]-rsum[j]) {
			j--
		}
		ans = min(ans, abs(goal-lsum[i]-rsum[j]))
	}
	return ans
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
