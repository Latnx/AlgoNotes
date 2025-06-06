package main

// https://leetcode.cn/problems/linked-list-cycle-ii
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	fast, slow := head.Next.Next, head.Next
	for fast != slow && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != slow && fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
