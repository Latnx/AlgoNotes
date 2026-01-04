package main

// 最后一块石头的重量 II
// 有一堆石头，用整数数组 stones 表示
// 其中 stones[i] 表示第 i 块石头的重量。
// 每一回合，从中选出任意两块石头，然后将它们一起粉碎
// 假设石头的重量分别为 x 和 y，且 x <= y
// 那么粉碎的可能结果如下：
// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x
// 最后，最多只会剩下一块 石头，返回此石头 最小的可能重量
// 如果没有石头剩下，就返回 0
// 测试链接 : https://leetcode.cn/problems/last-stone-weight-ii/
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, val := range stones {
		sum += val
	}
	// 找到小于等于目标的最大值
	A := find(stones, sum/2)
	return sum - 2*A
}
func find(stones []int, target int) int {
	dp := make([]int, target+1)
	for i := range stones {
		for j := target; j >= 1; j-- {
			if j-stones[i] >= 0 {
				dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
			}
		}
	}
	return dp[target]
}
