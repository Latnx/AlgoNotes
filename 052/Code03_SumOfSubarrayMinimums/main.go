package main

// 子数组的最小值之和
// 给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
// 由于答案可能很大，答案对 1000000007 取模
// 测试链接 : https://leetcode.cn/problems/sum-of-subarray-minimums/

// (cur - left) * (right - cur) * val
// 包含cur并不小于的开头*结尾
// 其中在右侧相等的时候, 不更新, 否则会计算重复
func sumSubarrayMins(arr []int) int {
	stack := make([]int, len(arr))
	lMin := make([]int, len(arr))
	rMin := make([]int, len(arr))
	r := 0
	for i, val := range arr {
		for r > 0 && val <= arr[stack[r-1]] {
			r--
			if r == 0 {
				lMin[stack[r]] = -1
			} else {
				lMin[stack[r]] = stack[r-1]
			}
			rMin[stack[r]] = i
		}
		stack[r] = i
		r++
	}
	for r > 0 {
		r--
		if r == 0 {
			lMin[stack[r]] = -1
		} else {
			lMin[stack[r]] = stack[r-1]
		}
		rMin[stack[r]] = len(arr)
	}
	// 不需要更新右侧相等
	res := 0
	for i := range len(arr) {
		res += (i - lMin[i]) * (rMin[i] - i) * arr[i]
		res %= 1000000007
	}
	return res
}
