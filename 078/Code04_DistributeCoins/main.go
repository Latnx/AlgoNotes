package main

// 在二叉树中分配硬币
// 给你一个有 n 个结点的二叉树的根结点 root
// 其中树中每个结点 node 都对应有 node.val 枚硬币
// 整棵树上一共有 n 枚硬币
// 在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点
// 移动可以是从父结点到子结点，或者从子结点移动到父结点
// 返回使每个结点上 只有 一枚硬币所需的 最少 移动次数
// 测试链接 : https://leetcode.cn/problems/distribute-coins-in-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	_, _, res := dfs(root)
	return res
}

// 节点数，硬币数量和移动次数
func dfs(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}
	lNode, lCoin, lStep := dfs(root.Left)
	rNode, rCoin, rStep := dfs(root.Right)

	return lNode + rNode + 1, lCoin + rCoin + root.Val, lStep + rStep + abs(lCoin-lNode) + abs(rCoin-rNode)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
