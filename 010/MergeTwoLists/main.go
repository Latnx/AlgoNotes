package main

// https://leetcode.cn/problems/merge-two-sorted-lists/
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var pre *ListNode = nil
	var cur1, cur2 = list1, list2

	if cur1.Val < cur2.Val {
		pre = cur1
		cur1 = cur1.Next
	} else {
		pre = cur2
		cur2 = cur2.Next
	}
	head := pre
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}
	if cur1 == nil {
		pre.Next = cur2
	}
	if cur2 == nil {
		pre.Next = cur1
	}
	return head
}
