package main

// https://leetcode.cn/problems/single-number-iii/
// 找到唯二出现的奇数次的数
// BK算法：提取一个数 二进制状态 最右侧的1 (aeorb & -aeorb)

func singleNumber(nums []int) []int {
	aeorb := 0
	for _, num := range nums {
		aeorb ^= num
	}
	last1 := (aeorb & -aeorb)
	a, b := 0, 0
	for _, num := range nums {
		if num&last1 == last1 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
