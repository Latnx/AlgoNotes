package main

import "maps"

// K个不同整数的子数组
// 给定一个正整数数组 nums和一个整数 k，返回 nums 中 「好子数组」 的数目。
// 如果 nums 的某个子数组中不同整数的个数恰好为 k
// 则称 nums 的这个连续、不一定不同的子数组为 「好子数组 」。
// 例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。
// 子数组 是数组的 连续 部分。
// 测试链接 : https://leetcode.cn/problems/subarrays-with-k-different-integers/
func subarraysSmallKDistinct(nums []int, k int) int {
	res := 0
	set := make(map[int]int)
	for l, r := 0, 0; r < len(nums); r++ {
		set[nums[r]]++
		for ; len(set) > k && l <= r; l++ {
			set[nums[l]]--
			if i := set[nums[l]]; i == 0 {
				delete(set, nums[l])
			}
		}
		res += r - l + 1
	}
	return res
}
func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysSmallKDistinct(nums, k) - subarraysSmallKDistinct(nums, k-1)
}

// 超时
func subarraysWithKDistinct2(nums []int, k int) int {
	res := 0
	set := make(map[int]int)
	for l, r := 0, 0; r < len(nums); r++ {
		set[nums[r]]++
		for l = r; len(set) > k && l <= r; l++ {
			set[nums[l]]--
			if i := set[nums[l]]; i == 0 {
				delete(set, nums[l])
			}
		}
		// 假收缩
		set_c := maps.Clone(set)
		for j := l; len(set_c) == k && j <= r; j++ {
			res++
			set_c[nums[j]]--
			if i := set_c[nums[j]]; i == 0 {
				delete(set_c, nums[l])
			}
		}
	}
	return res
}
func main() {
	subarraysSmallKDistinct([]int{1, 2, 1, 2, 3}, 2)
}
