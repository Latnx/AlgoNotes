package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(head *TreeNode) {
	if head == nil {
		return
	}
	print(head.Val, " ")
	preOrder(head.Left)
	preOrder(head.Right)
}

// 中序打印所有节点，递归版
func inOrder(head *TreeNode) {
	if head == nil {
		return
	}
	inOrder(head.Left)
	print(head.Val, " ")
	inOrder(head.Right)
}

// 后序打印所有节点，递归版
func posOrder(head *TreeNode) {
	if head == nil {
		return
	}
	posOrder(head.Left)
	posOrder(head.Right)
	print(head.Val, " ")
}
