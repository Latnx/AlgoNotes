package main

// https://leetcode.cn/problems/palindrome-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var pre, cur, next *ListNode = nil, slow, nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	p := head
	for pre != nil && p != nil {
		if pre.Val != p.Val {
			return false
		}
		pre = pre.Next
		p = p.Next
	}
	return true
}
