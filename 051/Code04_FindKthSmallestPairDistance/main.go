package main

import (
	"sort"
)

// 找出第K小的数对距离
// 数对 (a,b) 由整数 a 和 b 组成，其数对距离定义为 a 和 b 的绝对差值。
// 给你一个整数数组 nums 和一个整数 k
// 数对由 nums[i] 和 nums[j] 组成且满足 0 <= i < j < nums.length
// 返回 所有数对距离中 第 k 小的数对距离。
// 测试链接 : https://leetcode.cn/problems/find-k-th-smallest-pair-distance/

// 返回arr中任意两数的差值 <= limit，这样的数字配对，有几对
func cal(nums []int, limit int) (k int) {
	for i, j := 0, 1; i < len(nums); i++ {
		for j < len(nums) && nums[j]-nums[i] <= limit {
			j++
		}
		k += j - 1 + i
	}
	return
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, nums[len(nums)-1]-nums[0]
	res := 0
	for l <= r {
		mid := (l + r) / 2
		if cal(nums, mid) >= k {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func main() {
	smallestDistancePair([]int{1, 3, 1}, 1)
}
