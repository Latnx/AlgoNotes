package main

// 求二叉树的最大、最小深度
// 测试链接 : https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 测试链接 : https://leetcode.cn/problems/minimum-depth-of-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else if root.Left == nil {
		return minDepth(root.Right) + 1
	} else {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}
