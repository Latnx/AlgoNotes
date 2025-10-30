package main

import (
	"slices"
)

// 检查边长度限制的路径是否存在
// 测试链接 : https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/
var father []int

func build(n int) {
	father = make([]int, n)
	for i := range n {
		father[i] = i
	}
}
func find(i int) int {
	if father[i] != i {
		father[i] = find(father[i])
	}
	return father[i]
}
func inSameSet(x, y int) bool {
	return find(x) == find(y)
}
func union(x, y int) {
	father[find(x)] = find(y)
}

// queries[j] = [pj, qj, limitj]
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	build(n)
	// 给每个quries加序号并排序
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	slices.SortFunc(queries, func(a, b []int) int { return a[2] - b[2] })
	slices.SortFunc(edgeList, func(a, b []int) int { return a[2] - b[2] })
	res := make([]bool, len(queries))
	j := 0
	for i := range queries {
		for ; j < len(edgeList) && edgeList[j][2] < queries[i][2]; j++ {
			if !inSameSet(edgeList[j][0], edgeList[j][1]) {
				union(edgeList[j][0], edgeList[j][1])
			}
		}
		if inSameSet(queries[i][0], queries[i][1]) {
			res[queries[i][3]] = true
		}
	}
	return res
}
