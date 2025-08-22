package main

import "container/heap"

// 295. 数据流的中位数
// https://leetcode.cn/problems/find-median-from-data-stream/

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

type MedianFinder struct {
	minH minHeap
	maxH maxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minH: make(minHeap, 0),
		maxH: make(maxHeap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxH.Len() == 0 {
		heap.Push(&this.maxH, num)
		return
	}
	if num < this.maxH[0] {
		heap.Push(&this.maxH, num)
	} else {
		heap.Push(&this.minH, num)
	}
	if this.minH.Len() > this.maxH.Len() {
		x := heap.Pop(&this.minH).(int)
		heap.Push(&this.maxH, x)
	} else if this.minH.Len()+1 < this.maxH.Len() {
		x := heap.Pop(&this.maxH).(int)
		heap.Push(&this.minH, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// fmt.Println(this.minH, this.maxH)
	if this.maxH.Len() == this.minH.Len() {
		return (float64(this.maxH[0]) + float64(this.minH[0])) / 2
	} else {
		return float64(this.maxH[0])
	}
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.FindMedian()
}
