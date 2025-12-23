package main

import (
	"sort"
)

// 最长数对链
// 给你一个由n个数对组成的数对数组pairs
// 其中 pairs[i] = [lefti, righti] 且 lefti < righti
// 现在，我们定义一种 跟随 关系，当且仅当 b < c 时
// 数对 p2 = [c, d] 才可以跟在 p1 = [a, b] 后面
// 我们用这种形式来构造 数对链
// 找出并返回能够形成的最长数对链的长度
// 测试链接 : https://leetcode.cn/problems/maximum-length-of-pair-chain/
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	ends := make([]int, len(pairs))
	r := 0
	for _, pair := range pairs {
		first, second := pair[0], pair[1]
		L, R := 0, r
		for L < R {
			mid := (L + R) / 2
			if ends[mid] >= first {
				R = mid
			} else {
				L = mid + 1
			}
		}
		if L == r {
			ends[L] = second
			r++
		} else {
			ends[L] = min(ends[L], second)
		}
	}
	return r
}
func main() {
	findLongestChain([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	})
}
