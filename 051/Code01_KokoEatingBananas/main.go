package main

// 爱吃香蕉的珂珂
// 珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉
// 警卫已经离开了，将在 h 小时后回来。
// 珂珂可以决定她吃香蕉的速度 k （单位：根/小时)
// 每个小时，她将会选择一堆香蕉，从中吃掉 k 根
// 如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉
// 珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
// 返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）
// 测试链接 : https://leetcode.cn/problems/koko-eating-bananas/

// 二分答案法，找到答案的最大最小值，定出范围
func cal(piles []int, h int) int {
	res := 0
	for _, val := range piles {
		res += val / h
		if val%h != 0 {
			res++
		}
	}
	return res
}

func minEatingSpeed(piles []int, h int) int {
	maxSpeed := 0
	for _, val := range piles {
		maxSpeed = max(maxSpeed, val)
	}
	minSpeed := 1
	res := 0
	for minSpeed <= maxSpeed {
		mid := (minSpeed + maxSpeed) / 2
		if cal(piles, mid) > h {
			minSpeed = mid + 1
		} else {
			res = mid
			maxSpeed = mid - 1
		}
	}
	return res
}
