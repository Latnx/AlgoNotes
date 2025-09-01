package main

// 验证平衡二叉树
// 测试链接 : https://leetcode.cn/problems/balanced-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var isBalance bool = true

func isBalanced(root *TreeNode) bool {
	isBalance = true
	high(root)
	return isBalance
}

func high(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHigh := high(root.Left)
	rHigh := high(root.Right)
	if lHigh > rHigh+1 || rHigh > lHigh+1 {
		isBalance = false
	}
	return max(lHigh, rHigh) + 1
}
