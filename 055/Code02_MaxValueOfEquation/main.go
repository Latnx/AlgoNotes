package main

import "math"

// 满足不等式的最大值
// 给你一个数组 points 和一个整数 k
// 数组中每个元素都表示二维平面上的点的坐标，并按照横坐标 x 的值从小到大排序
// 也就是说 points[i] = [xi, yi]
// 并且在 1 <= i < j <= points.length 的前提下，xi < xj 总成立
// 请你找出 yi + yj + |xi - xj| 的 最大值，
// 其中 |xi - xj| <= k 且 1 <= i < j <= points.length
// 题目测试数据保证至少存在一对能够满足 |xi - xj| <= k 的点。
// 测试链接 : https://leetcode.cn/problems/max-value-of-equation/

// 单调队列:
// 1. j > i
// 2. xj - xi <= k
// 3. 求max(yi + yj + |xi - xj|) 也就是求在j固定且 > i 情况下, yj+xj+max(yi -xi)
// 查找del和add条件
// del(xj - xi > k)// 超出范围就弹出
// add(yi + yj) // 单调增, max(yi -xi)

// 其实就是维护了一个k大小(xj-xi<= k)范围内的极大值

func findMaxValueOfEquation(points [][]int, k int) int {
	// 因为需要淘汰, 所以没办法只维持一个最值,需要维持一个k的窗口,存放i
	queue := make([]int, len(points))
	l, r := 0, 0
	res := math.MinInt
	for i := range points {
		// 弹出不符合规则(xj - xi >k)的队头
		for l < r && points[i][0]-points[queue[l]][0] > k {
			l++
		}
		if l < r {
			res = max(res, points[queue[l]][1]-points[queue[l]][0]+points[i][1]+points[i][0])
		}
		// 对内维护值为yi -xi 递减情况, 因为要求max(yi -xi),
		for l < r && points[queue[r-1]][1]-points[queue[r-1]][0] <= points[i][1]-points[i][0] {
			r--
		}
		queue[r] = i
		r++
	}
	return res
}

func main() {
	points := [][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}
	k := 1
	findMaxValueOfEquation(points, k)
}
