package main

// https://leetcode.cn/problems/reverse-bits
const (
	m1 = 0x55555555
	m2 = 0x33333333
	m3 = 0x0f0f0f0f
	m4 = 0x00ff00ff
)

func reverseBits(num uint32) uint32 {
	num = num>>1&m1 | num&m1<<1
	num = num>>2&m2 | num&m2<<2
	num = num>>4&m3 | num&m3<<4
	num = num>>8&m4 | num&m4<<8
	return num>>16 | num<<16
}
