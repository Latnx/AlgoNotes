package main

// 堆排序 https://leetcode.cn/problems/sort-an-array/
/*
位置： 父：(i-1)/2 子：i*2+1, i*2+2
两种调整：上浮和下沉O(logn)
两种操作：Pop需要下沉，Push需要上浮(logn)
初始化需要从一半开始不断下沉O(nlogn)
*/

import (
	"fmt"
)

type heap []int

// i位置数变小了 向下调整
func (h heap) heapify(i int) {
	size := len(h)
	l := 2*i + 1
	for l < size {
		best := l
		if l+1 < size && h[l+1] > h[l] {
			best = l + 1
		}
		if h[best] <= h[i] {
			break
		}
		h[best], h[i] = h[i], h[best]
		i = best
		l = 2*i + 1
	}
}

// i数变大 向上调整
func (h heap) heapInsert(i int) {
	for ; h[i] > h[(i-1)/2]; i = (i - 1) / 2 {
		h[i], h[(i-1)/2] = h[(i-1)/2], h[i]
	}
}

func (h *heap) Push(sum int) {
	*h = append(*h, sum)
	h.heapInsert(len(*h) - 1)
}

func (h *heap) Pop() int {
	x := (*h)[0]
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	(*h) = (*h)[:len(*h)-1]
	h.heapify(0)
	return x
}

// 从顶到底建堆，使用Push逐个建堆，O(N*logN)
// 从底到顶建堆，找有孩子的节点开始向下调整，逐渐包含整个堆 O(N)
func (h heap) Init() {
	for i := (len(h) - 1) / 2; i >= 0; i-- {
		h.heapify(i)
	}
}

func main() {
	var h heap = heap{3, 2, 1, 4, 5, 6, 32, 1, 2, 4, 6, 3, 12}
	h.Init()
	for range h {
		fmt.Print(h.Pop(), " ")
	}
}
