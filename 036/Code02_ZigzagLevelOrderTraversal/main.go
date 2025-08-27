package main

// 二叉树的锯齿形层序遍历
// 测试链接 : https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	*TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var res [][]int
	for level := 0; len(queue) > 0; level++ {
		l := queue
		queue = nil
		r := []int{}
		for _, curNode := range l {
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			r = append(r, curNode.Val)
		}
		if level%2 == 1 {
			for i := 0; i < len(r); i++ {
				r[i], r[len(r)-i] = r[len(r)-i], r[i]
			}
		}
		res = append(res, r)
	}
	return res
}
