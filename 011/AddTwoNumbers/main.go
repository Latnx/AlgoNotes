package main

// https://leetcode.cn/problems/add-two-numbers/
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	sum := 0
	head := ListNode{}
	pre := &head
	for l1 != nil && l2 != nil {
		sum = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		l1 = l1.Next
		l2 = l2.Next
		pre.Next = &ListNode{}
		pre = pre.Next
		pre.Val = sum
	}

	for l1 != nil {
		sum = (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		l1 = l1.Next
		pre.Next = &ListNode{}
		pre = pre.Next
		pre.Val = sum
	}
	for l2 != nil {
		sum = (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		l2 = l2.Next
		pre.Next = &ListNode{}
		pre = pre.Next
		pre.Val = sum
	}
	if carry != 0 {
		pre.Next = &ListNode{}
		pre = pre.Next
		pre.Val = carry
	}
	return head.Next
}
