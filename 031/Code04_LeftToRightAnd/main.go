package main

// https://leetcode.cn/problems/bitwise-and-of-numbers-range/

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right ^= (right & -right)
	}
	return right
}
