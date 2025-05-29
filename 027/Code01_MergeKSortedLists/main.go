package main

import "container/heap"

// https://leetcode.cn/problems/merge-k-sorted-lists
// 合并 K 个升序链表
type ListNode struct {
	Val  int
	Next *ListNode
}

type hp [](*ListNode)

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { (*h) = append(*h, x.(*ListNode)) }
func (h *hp) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

func mergeKLists(lists []*ListNode) *ListNode {
	var h hp
	for index, list := range lists {
		if list != nil {
			h.Push(list)
			lists[index] = list.Next
		}
	}
	heap.Init(&h)
	head := &ListNode{}
	p := head
	for len(h) > 0 {
		p.Next = heap.Pop(&h).(*ListNode)
		p = p.Next
		if p.Next != nil {
			heap.Push(&h, p.Next)
		}
	}
	return head.Next
}

func main() {
	mergeKLists([]*ListNode{})
}
