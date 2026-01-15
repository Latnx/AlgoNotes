package main

// 监控二叉树
// 给定一个二叉树，我们在树的节点上安装摄像头
// 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象
// 计算监控树的所有节点所需的最小摄像头数量
// 测试链接 : https://leetcode.cn/problems/binary-tree-cameras/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	t, res := dfs(root)
	if t == 0 {
		return res + 1
	}
	return res
}

// 0,1,2 三状态表明，未覆盖，已覆盖无摄像头，已覆盖有摄像头
// 摄像头数
func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 1, 0
	}
	lType, lNum := dfs(root.Left)
	rType, rNum := dfs(root.Right)
	if lType == 1 && rType == 1 {
		return 0, lNum + rNum
	}
	if lType == 0 || rType == 0 {
		return 2, lNum + rNum + 1
	}
	return 1, lNum + rNum
}
