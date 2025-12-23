package main

// 使数组K递增的最少操作次数
// 给你一个下标从0开始包含n个正整数的数组arr，和一个正整数k
// 如果对于每个满足 k <= i <= n-1 的下标 i
// 都有 arr[i-k] <= arr[i] ，那么称 arr 是K递增的
// 每一次操作中，你可以选择一个下标i并将arr[i]改成任意正整数
// 请你返回对于给定的 k ，使数组变成K递增的最少操作次数
// 测试链接 : https://leetcode.cn/problems/minimum-operations-to-make-the-array-k-increasing/

func lengthOfLIS(nums []int) int {
	ends := make([]int, len(nums)+1)
	r := 0
	for i := range nums {
		L, R := 0, r
		for L < R {
			mid := (L + R) / 2
			if ends[mid] > nums[i] {
				R = mid
			} else {
				L = mid + 1
			}
		}
		ends[L] = nums[i]
		if L == r {
			r++
		}
	}
	return r
}

func kIncreasing(arr []int, k int) int {
	n := len(arr)
	//分成K组
	res := 0
	for i := 0; i < k; i++ {
		nums := make([]int, 0, n/k)
		for j := i; j < n; j += k {
			nums = append(nums, arr[j])
		}
		res += len(nums) - lengthOfLIS(nums)
	}
	return res
}
