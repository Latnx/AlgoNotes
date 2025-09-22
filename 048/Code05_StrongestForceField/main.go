package main

import (
	"slices"
)

// 最强祝福力场
// 小扣在探索丛林的过程中，无意间发现了传说中"落寞的黄金之都"
// 而在这片建筑废墟的地带中，小扣使用探测仪监测到了存在某种带有「祝福」效果的力场
// 经过不断的勘测记录，小扣将所有力场的分布都记录了下来
// forceField[i] = [x,y,side]
// 表示第 i 片力场将覆盖以坐标 (x,y) 为中心，边长为 side 的正方形区域。
// 若任意一点的 力场强度 等于覆盖该点的力场数量
// 请求出在这片地带中 力场强度 最强处的 力场强度
// 注意：力场范围的边缘同样被力场覆盖。
// 测试链接 : https://leetcode.cn/problems/xepqZ5/
var diff [][]int

func fieldOfGreatestBlessing(forceField [][]int) int {
	xs, ys := make([]int, 2*len(forceField)), make([]int, 2*len(forceField))
	for i, field := range forceField {
		x := field[0]
		y := field[1]
		r := field[2]
		xs[2*i] = 2*x - r
		xs[2*i+1] = 2*x + r
		ys[2*i] = 2*y - r
		ys[2*i+1] = 2*y + r
	}
	// 下标更改
	xm := sort(xs)
	ym := sort(ys)
	diff = make([][]int, len(xs)+2)
	for i := range diff {
		diff[i] = make([]int, len(ys)+2)
	}

	for _, field := range forceField {
		x := field[0]
		y := field[1]
		r := field[2]
		set(xm[2*x-r]+1, ym[2*y-r]+1, xm[2*x+r]+1, ym[2*y+r]+1)
	}
	build()
	res := 0
	for i := range len(xs) {
		for j := range len(ys) {
			res = max(res, diff[i+1][j+1])
		}
	}
	return res
}

func sort(s []int) map[int]int {
	slices.Sort(s)
	size := 0
	m := make(map[int]int)
	for _, val := range s {
		if s[size] != val {
			size++
			s[size] = val
		}
		m[val] = size
	}
	return m
}

func set(row1, col1, row2, col2 int) {
	diff[row1][col1] += 1
	diff[row1][col2+1] -= 1
	diff[row2+1][col1] -= 1
	diff[row2+1][col2+1] += 1
}

func build() {
	for i := 1; i < len(diff); i++ {
		for j := 1; j < len(diff[0]); j++ {
			diff[i][j] += diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1]
		}
	}
}
