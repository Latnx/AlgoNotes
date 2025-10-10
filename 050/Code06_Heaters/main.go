package main

import (
	"math"
	"sort"
)

// 供暖器
// 冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
// 在加热器的加热半径范围内的每个房屋都可以获得供暖。
// 现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置
// 请你找出并返回可以覆盖所有房屋的最小加热半径。
// 说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
// 测试链接 : https://leetcode.cn/problems/heaters/
func abs(num int) int {
	return int(math.Abs(float64(num)))
}
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	l, r := 0, 0
	res := 0
	for l < len(houses) {
		if r < len(heaters)-1 && abs(heaters[r]-houses[l]) >= abs(heaters[r+1]-houses[l]) {
			r++
		} else {
			res = max(res, abs(heaters[r]-houses[l]))
			l++
		}
	}
	return res
}
