package main

// 求完全二叉树的节点个数, 时间复杂度小于 O(N)
// 测试链接 : https://leetcode.cn/problems/count-complete-tree-nodes/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func highLeft(root *TreeNode) int {
	h := 0
	for r := root; r != nil; r = r.Left {
		h++
	}
	return h
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 分别查看左右两棵树哪个是满二叉树
	// 		左子树的高度和右子树的最左是否相等
	// 			如果相等，左子树是满的
	// 			如果不相等，右子树是满的
	hL := highLeft(root.Left)
	hR := highLeft(root.Right)
	if hL == hR {
		return 1<<hL - 1 + countNodes(root.Left) + 1
	} else {
		return 1<<hR - 1 + countNodes(root.Right) + 1
	}
}
