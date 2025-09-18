package main

import "strconv"

// 牛牛和他的朋友们约定了一套接头密匙系统，用于确认彼此身份
// 密匙由一组数字序列表示，两个密匙被认为是一致的，如果满足以下条件：
// 密匙 b 的长度不超过密匙 a 的长度。
// 对于任意 0 <= i < length(b)，有b[i+1] - b[i] == a[i+1] - a[i]
// 现在给定了m个密匙 b 的数组，以及n个密匙 a 的数组
// 请你返回一个长度为 m 的结果数组 ans，表示每个密匙b都有多少一致的密匙
// 数组 a 和数组 b 中的元素个数均不超过 10^5
// 1 <= m, n <= 1000
// 测试链接 : https://www.nowcoder.com/practice/c552d3b4dfda49ccb883a6371d9a6932
const MAXN = 10005

var tree [][11]int
var pass []int
var cnt int

func countConsistentKeys(b [][]int, a [][]int) []int {
	cnt = 1
	pass = make([]int, MAXN)
	tree = make([][11]int, MAXN)
	res := make([]int, 0)
	for _, val := range a {
		pre := make([]byte, 0)
		for j := 1; j < len(val); j++ {
			pre = append(pre, strconv.Itoa(val[j]-val[j-1])...)
		}
		Insert(string(pre))
	}
	for _, val := range b {
		pre := make([]byte, 0)
		for j := 1; j < len(val); j++ {
			pre = append(pre, strconv.Itoa(val[j]-val[j-1])...)
		}
		res = append(res, StartsWith(string(pre)))
	}
	return res
}

func Insert(word string) {
	var next = 1
	pass[next]++
	for _, val := range word {
		path := val - '0'
		if val == '-' {
			path = 10
		}
		if tree[next][path] == 0 {
			cnt++
			tree[next][path] = cnt
		}
		next = tree[next][path]
		pass[next]++
	}
}

func StartsWith(pre string) int {
	var next = 1
	for _, val := range pre {
		path := val - '0'
		if val == '-' {
			path = 10
		}
		if tree[next][path] == 0 {
			return 0
		}
		next = tree[next][path]
	}
	return pass[next]
}
