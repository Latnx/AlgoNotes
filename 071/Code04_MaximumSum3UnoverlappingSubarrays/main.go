package main

// 三个无重叠子数组的最大和
// 给你一个整数数组 nums 和一个整数 k
// 找出三个长度为 k 、互不重叠、且全部数字和（3 * k 项）最大的子数组
// 并返回这三个子数组
// 以下标的数组形式返回结果，数组中的每一项分别指示每个子数组的起始位置
// 如果有多个结果，返回字典序最小的一个
// 测试链接 : https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays/
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	prefix, suffix := make([]int, n), make([]int, n)
	sum := make([]int, n)
	num := 0
	for i := range k {
		num += nums[i]
	}
	sum[0] = num
	for i := 0; i+k < n; i++ {
		num = num - nums[i] + nums[i+k]
		sum[i+1] = num
	}
	preMaxIndex := 0
	for i := range sum {
		if sum[i] > sum[preMaxIndex] {
			preMaxIndex = i
		}
		prefix[i] = preMaxIndex
	}
	sufMaxIndex := n - 1
	for i := n - 1; i >= 0; i-- {
		if sum[i] >= sum[sufMaxIndex] {
			sufMaxIndex = i
		}
		suffix[i] = sufMaxIndex
	}
	res := make([]int, 3)
	total := 0
	for i := 0; i+2*k < n; i++ {
		l := prefix[i]
		mid := i + k
		r := suffix[i+2*k]
		if total < sum[l]+sum[mid]+sum[r] {
			total = sum[l] + sum[mid] + sum[r]
			res[0], res[1], res[2] = l, mid, r
		}
	}
	return res
}

func main() {
	maxSumOfThreeSubarrays([]int{9, 8, 7, 6, 2, 2, 2, 2}, 2)
}
