package main

import "math"

// 二叉搜索子树的最大键值和
// 给你一棵以 root 为根的二叉树
// 请你返回 任意 二叉搜索子树的最大键值和
// 二叉搜索树的定义如下：
// 任意节点的左子树中的键值都 小于 此节点的键值
// 任意节点的右子树中的键值都 大于 此节点的键值
// 任意节点的左子树和右子树都是二叉搜索树
// 测试链接 : https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Info struct {
	sum       int
	minVal    int
	maxVal    int
	isBst     bool
	maxBstSum int
}

func maxSumBST(root *TreeNode) int {
	return dfs(root).maxBstSum
}
func dfs(root *TreeNode) Info {
	if root == nil {
		return Info{maxVal: math.MinInt, minVal: math.MaxInt, isBst: true}
	}
	lTree := dfs(root.Left)
	rTree := dfs(root.Right)
	if lTree.isBst && rTree.isBst && lTree.maxVal < root.Val && rTree.minVal > root.Val {
		return Info{
			sum:       lTree.sum + rTree.sum + root.Val,
			maxVal:    max(rTree.maxVal, root.Val),
			minVal:    min(lTree.minVal, root.Val),
			isBst:     true,
			maxBstSum: max(lTree.sum+rTree.sum+root.Val, lTree.maxBstSum, rTree.maxBstSum),
		}
	} else {
		return Info{
			// sum:       max(lTree.sum, rTree.sum),
			// maxVal:    max(rTree.maxVal, root.Val),
			// minVal:    min(lTree.minVal, root.Val),
			isBst:     false,
			maxBstSum: max(lTree.maxBstSum, rTree.maxBstSum),
		}
	}
}
