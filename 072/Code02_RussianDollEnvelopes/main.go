package main

import (
	"sort"
)

// 俄罗斯套娃信封问题
// 给你一个二维整数数组envelopes ，其中envelopes[i]=[wi, hi]
// 表示第 i 个信封的宽度和高度
// 当另一个信封的宽度和高度都比这个信封大的时候
// 这个信封就可以放进另一个信封里，如同俄罗斯套娃一样
// 请计算 最多能有多少个信封能组成一组“俄罗斯套娃”信封
// 即可以把一个信封放到另一个信封里面，注意不允许旋转信封
// 测试链接 : https://leetcode.cn/problems/russian-doll-envelopes/
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})
	ends := make([]int, len(envelopes))
	ends[0] = envelopes[0][1]
	r := 1
	for i := range envelopes {
		L, R := 0, r
		for L < R {
			mid := (L + R) / 2
			if ends[mid] >= envelopes[i][1] {
				R = mid
			} else {
				L = mid + 1
			}
		}
		ends[L] = envelopes[i][1]
		if L == r {
			ends[r] = envelopes[i][1]
			r++
		}
	}
	return r
}
