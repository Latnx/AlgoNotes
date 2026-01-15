package main

// 路径总和 III
// 给定一个二叉树的根节点 root ，和一个整数 targetSum
// 求该二叉树里节点值之和等于 targetSum 的 路径 的数目
// 路径 不需要从根节点开始，也不需要在叶子节点结束
// 但是路径方向必须是向下的（只能从父节点到子节点）
// 测试链接 : https://leetcode.cn/problems/path-sum-iii/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	preSum = make(map[int]int)
	preSum[0] = 1
	res = 0
	dfs(root, targetSum, 0)
	return res
}

var preSum map[int]int
var res int

func dfs(root *TreeNode, targetSum, pre int) {
	if root == nil {
		return
	}
	sum := pre + root.Val
	res += preSum[sum-targetSum]
	preSum[sum]++
	dfs(root.Left, targetSum, sum)
	dfs(root.Right, targetSum, sum)
	preSum[sum]--
}
