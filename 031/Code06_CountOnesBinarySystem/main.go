package main

// https://leetcode.cn/problems/hamming-distance/

const (
	m1 = 0x55555555
	m2 = 0x33333333
	m3 = 0x0f0f0f0f
	m4 = 0x00ff00ff
	m5 = 0x0000ffff
)

func hammingDistance(x int, y int) int {
	xeory := x ^ y
	xeory = xeory>>1&m1 + xeory&m1
	xeory = xeory>>2&m2 + xeory&m2
	xeory = xeory>>4&m3 + xeory&m3
	xeory = xeory>>8&m4 + xeory&m4
	xeory = xeory>>16&m5 + xeory&m5
	return xeory
}
