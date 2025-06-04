package main

//https://leetcode.cn/problems/single-number-ii/

func singleNumber(nums []int) int {
	bits := [32]int{}
	for _, num := range nums {
		for j := range 32 {
			bits[j] += (num) >> j & 1
		}
	}
	res := 0
	for i, bit := range bits {
		b := 0
		if bit%3 != 0 {
			b = 1
		}
		res |= b << i
	}
	return int(int32(res))
}
