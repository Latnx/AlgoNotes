package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for p := head; p != nil; p = p.Next.Next {
		p.Next = &Node{
			Val:  p.Val,
			Next: p.Next,
		}
	}
	for p := head; p != nil; p = p.Next.Next {
		cur := p.Next
		if p.Random == nil {
			cur.Random = nil
		} else {
			cur.Random = p.Random.Next
		}
	}
	next := head
	res := head.Next
	for p := head; p != nil; {
		next = p.Next.Next
		cur := p.Next
		p.Next = next
		p = p.Next
		if p == nil {
			cur.Next = nil
		} else {
			cur.Next = p.Next
		}
	}
	return res
}
