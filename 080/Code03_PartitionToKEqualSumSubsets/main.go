package main

// 划分为k个相等的子集
// 给定一个整数数组  nums 和一个正整数 k，
// 找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
// 测试链接 : https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	dp = make([]int, 1<<len(nums))
	return f(nums, sum/k, 1<<len(nums)-1, 0, k)
}

var dp []int

func f(nums []int, limit int, status, cur, rest int) bool {
	if rest == 0 {
		return status == 0
	}
	if dp[status] != 0 {
		return dp[status] == 1
	}
	ans := false
	for i := range nums {
		if status&(1<<i) != 0 && cur+nums[i] <= limit {
			if cur+nums[i] == limit {
				ans = f(nums, limit, status^(1<<i), 0, rest-1)
			} else {
				ans = f(nums, limit, status^(1<<i), cur+nums[i], rest)
			}
			if ans {
				dp[status] = 1
				return true
			}
		}
	}
	dp[status] = 2
	return ans
}
