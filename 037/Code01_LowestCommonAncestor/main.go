package main

// 普通二叉树上寻找两个节点的最近公共祖先
// 测试链接 : https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	a := lowestCommonAncestor(root.Left, p, q)
	b := lowestCommonAncestor(root.Right, p, q)
	if a == nil && b == nil {
		return nil
	} else if a == nil && b != nil {
		return b
	} else if a != nil && b == nil {
		return a
	} else {
		return root
	}
}
