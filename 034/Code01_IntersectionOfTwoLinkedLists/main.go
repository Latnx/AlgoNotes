package main

// Code01_IntersectionOfTwoLinkedLists
// https://leetcode.cn/problems/intersection-of-two-linked-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0
	for p := headA; p != nil; p = p.Next {
		lenA++
	}
	for p := headB; p != nil; p = p.Next {
		lenB++
	}
	p1, p2 := headA, headB
	if lenA < lenB {
		for range lenB - lenA {
			p2 = p2.Next
		}
	} else {
		for range lenA - lenB {
			p1 = p1.Next
		}
	}
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}
