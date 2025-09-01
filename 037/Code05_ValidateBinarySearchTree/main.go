package main

import (
	"math"
)

// 验证搜索二叉树
// 测试链接 : https://leetcode.cn/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var isValid bool

func isValidBST(root *TreeNode) bool {
	isValid = true
	mmNum(root)
	return isValid
}

func mmNum(root *TreeNode) (int, int) {
	if root == nil {
		return math.MaxInt64, math.MinInt64
	}
	lMin, lMax := mmNum(root.Left)
	rMin, rMax := mmNum(root.Right)
	if lMax >= root.Val || rMin <= root.Val {
		isValid = false
	}
	return min(lMin, root.Val), max(rMax, root.Val)
}

// 中序遍历递增
var preNodeVal []int

func isValidBST2(root *TreeNode) bool {
	preNodeVal = nil
	for i := 0; i < len(preNodeVal)-1; i++ {
		if preNodeVal[i] >= preNodeVal[i+1] {
			return false
		}
	}
	return true
}
func midOrder(root *TreeNode) {
	if root == nil {
		return
	}
	midOrder(root.Left)
	preNodeVal = append(preNodeVal, root.Val)
	midOrder(root.Right)
}
