package main

// https://leetcode.cn/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	return n > 0 && n^(n&(-n)) == 0
}
