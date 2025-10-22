package main

import "math"

// 和至少为K的最短子数组
// 给定一个数组arr，其中的值有可能正、负、0
// 给定一个正数k
// 返回累加和>=k的所有子数组中，最短的子数组长度
// 测试链接 : https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/

func shortestSubarrayBase(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	for index, val := range nums {
		sum[index+1] = sum[index] + val
	}
	res := math.MaxInt
	for index := range sum {
		for i := 0; index-i >= 0; i++ {
			if sum[index]-sum[index-i] >= k {
				res = min(res, i)
				break
			}
		}
	}
	return res
}

// 基础做法超时，
// 需要维持一个单调队列：队列中维持了从小到大排序的前缀和（下标）
// 分别从对头，和队尾淘汰对象，
// 头部为满足的答案, 队尾为不再可能成为答案的对象..
// 先分析题目,随着前缀和index的靠后, 可以根据单调性淘汰掉
func shortestSubarray(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	for index, val := range nums {
		sum[index+1] = sum[index] + val
	}
	queue := make([]int, len(nums)+1)
	l, r := 0, 0
	res := math.MaxInt
	for index, val := range sum {
		// 队列不空，且头达标 del
		for l < r && val-sum[queue[l]] >= k {
			res = min(res, index-queue[l])
			l++
		}
		// 弹出大于等于当前值的值，因为val达标 队尾肯定达标，而又需要更短，因此保留后面的
		for l < r && sum[queue[r-1]] >= val {
			r--
		}
		queue[r] = index
		r++
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
func main() {
	shortestSubarray([]int{56, -21, 56, 35, -9}, 61)
}
