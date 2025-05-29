package main

// https://leetcode.cn/problems/minimum-operations-to-halve-array-sum
// 将数组和减半的最少操作次数

import "container/heap"

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { (*h) = append(*h, x.(int)) }
func (h *hp) Pop() any          { x := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return x }

func halveArray(nums []int) int {
	sum := 0
	n := 0
	var h hp = make([]int, len(nums))
	for i, num := range nums {
		// 要点，需要左移
		h[i] = num << 20
		sum += h[i]
	}
	heap.Init(&h)
	subtotal := 0
	for ; sum/2 > subtotal; n++ {
		h[0] /= 2
		subtotal += h[0]
		heap.Fix(&h, 0)
	}
	return n
}

func main() {
	halveArray([]int{})
}
