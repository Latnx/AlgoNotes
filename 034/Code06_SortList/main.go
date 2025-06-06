package main

// https://leetcode.cn/problems/sort-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head, nil)
}

func mergeSort(l, r *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	if l.Next == r {
		l.Next = nil
		return l
	}
	fast, slow := l, l
	for fast != r && fast.Next != r {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return merge(mergeSort(l, slow), mergeSort(slow, r))
}

func merge(l, m *ListNode) *ListNode {
	head := &ListNode{}
	pre := head
	for l != nil && m != nil {
		if l.Val <= m.Val {
			pre.Next = l
			l = l.Next
		} else {
			pre.Next = m
			m = m.Next
		}
		pre = pre.Next
	}
	if l != nil {
		pre.Next = l
	} else {
		pre.Next = m
	}
	return head.Next
}

func main() {
	sortList(&ListNode{})
}
