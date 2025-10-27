package main

import (
	"sort"
)

// 好路径的数目
// 长度为 n 下标从 0 开始的整数数组 vals ，分别表示每个节点的值。
// 二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。

// 开始节点和结束节点的值 相同 。
// 开始节点和结束节点中间的所有节点值都 小于等于 开始节点的值（也就是说开始节点的值应该是路径上所有节点的最大值）。
// 测试链接 : https://leetcode.cn/problems/number-of-good-paths/

// 每个点初始是一条“好路径”；
//
// 用 DSU（并查集）逐步合并；
// 当合并的两个集合中最大值相同，就能产生新的好路径数量 = 两集合内该值的节点数量乘积；

// 排序路径
type SortBy [][]int

var val []int

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	return max(val[a[i][0]], val[a[i][1]]) < max(val[a[j][0]], val[a[j][1]])
}

// 集合中最大值的数量
var setNum []int

// 集合中最大值
var setMax []int

// 并查集
var arr []int

func find(i int) int {
	if arr[i] != i {
		arr[i] = find(arr[i])
	}
	return arr[i]
}
func union(x, y int) {
	findX, findY := find(x), find(y)
	if setMax[findX] > setMax[findY] {
		arr[findY] = findX
	} else if setMax[findX] < setMax[findY] {
		arr[findX] = findY
	} else if setMax[findX] == setMax[findY] {
		arr[findY] = findX
		setNum[findX] += setNum[findY]
	}
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	// 先处理两个端点最大值小的边,按照边两端最大节点值升序排序；
	val = vals
	sort.Sort(SortBy(edges))

	arr = make([]int, len(vals))
	setNum = make([]int, len(vals))
	setMax = make([]int, len(vals))
	for i := range arr {
		arr[i] = i
		setNum[i] = 1
		setMax[i] = vals[i]
	}
	res := len(vals)

	for i := range edges {
		if setMax[find(edges[i][0])] == setMax[find(edges[i][1])] {
			res += setNum[find(edges[i][0])] * setNum[find(edges[i][1])]
		}
		union(edges[i][0], edges[i][1])
	}
	return res
}
