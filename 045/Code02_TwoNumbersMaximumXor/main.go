package main

// 数组中两个数的最大异或值
// 给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0<=i<=j<=n
// 1 <= nums.length <= 2 * 10^5
// 0 <= nums[i] <= 2^31 - 1
// 测试链接 : https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/

const MAXN = 10000

var tree [][2]int = make([][2]int, MAXN)
var cnt = 1

func findMaximumXOR(nums []int) int {
	clear()
	cnt = 1
	for _, val := range nums {
		Insert(val)
	}
	res := 0
	for _, val := range nums {
		res = max(findMax(val), res)
	}
	return res
}

func clear() {
	for i := range tree {
		for j := range tree[i] {
			tree[i][j] = 0
		}
	}
}

func Insert(num int) {
	next := 1
	for i := 31; i >= 0; i-- {
		path := (num >> i) & 1
		if tree[next][path] == 0 {
			cnt++
			tree[next][path] = cnt
		}
		next = tree[next][path]
	}
}

func findMax(num int) int {
	var res = 0
	next := 1
	for i := 31; i >= 0; i-- {
		path := (num >> i) & 1
		if path == 0 && tree[next][1] != 0 {
			next = tree[next][1]
			res = res<<1 | 1
		} else if path == 1 && tree[next][0] != 0 {
			next = tree[next][0]
			res = res<<1 | 1
		} else {
			next = tree[next][path]
			res = res<<1 | 0
		}
	}
	return res
}
