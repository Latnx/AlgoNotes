package main

// 二叉树打家劫舍问题
// 测试链接 : https://leetcode.cn/problems/house-robber-iii/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(calRob(root))
}

func calRob(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	calL, nCalL := calRob(root.Left)
	calR, nCalR := calRob(root.Right)
	return root.Val + nCalL + nCalR, max(calL, nCalL) + max(calR, nCalR)
}
