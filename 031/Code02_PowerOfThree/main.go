package main

// https://leetcode.cn/problems/power-of-three/

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
