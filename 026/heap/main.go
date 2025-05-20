package main

import (
	"container/heap"
	"fmt"
)

type pair struct{ x, y int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].x < h[j].x }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *hp) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

func main() {
	var h hp
	h.Push(pair{3, 4})
	h.Push(pair{2, 1})
	h.Push(pair{1, 3})
	h.Push(pair{10, 3})
	h.Push(pair{11, 3})
	h.Push(pair{9, 3})
	h.Push(pair{6, 3})
	h.Push(pair{3, 3})
	heap.Init(&h)
	for h.Len() > 0 {
		fmt.Print(heap.Pop(&h))
	}
}
