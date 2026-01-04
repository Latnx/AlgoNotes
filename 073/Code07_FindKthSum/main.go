package main

import (
	"container/heap"
	"slices"
)

// 同学找到的在线测试，思路来自课上的题目6
// 找出数组的第K大和
// 给定一个数组nums和正数k
// 可以选择数组的任何子序列并对其元素求和
// 希望找到第k大的子序列和，子序列和允许出现重复
// 空子序列的和视作0，数组中的值可正、可负、可0
// 测试链接 : https://leetcode.cn/problems/find-the-k-sum-of-an-array/description/
type ele struct {
	idx int
	sum int
}
type hp []ele

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { (*h) = append((*h), x.(ele)) }
func (h *hp) Pop() any          { x := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return x }

// 堆为{下标序列， sum}
// 展开策略
// 1. 删除最后下标，将下标的下一个加入堆
// 2. 不删除下标，将下标的下一个加入堆
func kSum(nums []int, k int) int64 {
	sum := 0
	for i, val := range nums {
		if val > 0 {
			sum += val
		} else {
			nums[i] = -nums[i]
		}
	}
	slices.Sort(nums)
	h := make(hp, 0)
	heap.Init(&h)
	heap.Push(&h, ele{-1, 0})
	for range k - 1 {
		e := heap.Pop(&h).(ele)
		lastIndex := e.idx
		if lastIndex+1 < len(nums) {
			heap.Push(&h, ele{e.idx + 1, e.sum + nums[lastIndex+1]})
			if lastIndex >= 0 {
				heap.Push(&h, ele{e.idx + 1, e.sum - nums[lastIndex] + nums[lastIndex+1]})
			}
		}
	}
	return int64(sum - heap.Pop(&h).(ele).sum)
}
func main() {
	kSum([]int{2, 4, -2}, 5)
}
