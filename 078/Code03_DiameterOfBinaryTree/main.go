package main

// 二叉树的直径
// 给你一棵二叉树的根节点，返回该树的直径
// 二叉树的 直径 是指树中任意两个节点之间最长路径的长度
// 这条路径可能经过也可能不经过根节点 root
// 两节点之间路径的 长度 由它们之间边数表示
// 测试链接 : https://leetcode.cn/problems/diameter-of-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最大值，单边最大值
func diameterOfBinaryTree(root *TreeNode) int {
	res, _ := f1(root)
	return res
}
func f1(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lMax, lMaxS := f1(root.Left)
	rMax, rMaxS := f1(root.Right)
	return max(lMax, rMax, lMaxS+rMaxS), max(lMaxS, rMaxS) + 1
}
