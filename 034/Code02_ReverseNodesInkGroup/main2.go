package main

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var res *ListNode = &ListNode{Next: head}
	var pre, cur, next *ListNode = res, head, nil
	for cur != nil {
		tmpHead := pre

		i := 0
		for t := cur; i < k && t != nil; i++ {
			t = t.Next
		}
		if i != k {
			break
		}
		i = 0
		for ; i < k && cur != nil; i++ {
			next = cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		tmpHead.Next.Next = cur
		tmpHead.Next, pre = pre, tmpHead.Next
	}
	return res.Next
}

func createList() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	return head
}

func main() {
	reverseKGroup(createList(), 2)
}
