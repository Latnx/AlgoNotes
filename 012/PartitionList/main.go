package main

// https://leetcode.cn/problems/partition-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	p1, p2 := &ListNode{}, &ListNode{}
	var cur1, cur2 *ListNode = p1, p2

	var cur *ListNode = head
	for cur != nil {
		if cur.Val < x {
			cur1.Next = cur
			cur1 = cur1.Next
		} else if cur.Val >= x {
			cur2.Next = cur
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	cur1.Next = p2.Next
	cur2.Next = nil
	return p1
}
