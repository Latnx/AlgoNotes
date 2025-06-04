package main

// https://leetcode.cn/problems/missing-number/
func missingNumber(nums []int) int {
	eorAll, eorHas := 0, 0
	for i, num := range nums {
		eorAll ^= i
		eorHas ^= num
	}
	eorAll ^= len(nums)
	return eorAll ^ eorHas
}
