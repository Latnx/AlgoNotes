package main

// 二叉树的最大特殊宽度
// 测试链接 : https://leetcode.cn/problems/maximum-width-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type node struct {
	TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*node
	queue = append(queue, &node{*root, 1})
	res := 1
	for len(queue) != 0 {
		l := queue
		queue = nil
		for _, curNode := range l {
			if curNode.Left != nil {
				queue = append(queue, &node{*curNode.Left, curNode.index * 2})
			}
			if curNode.Right != nil {
				queue = append(queue, &node{*curNode.Right, curNode.index*2 + 1})
			}
		}
		if len(queue) > 0 {
			res = max(res, queue[len(queue)-1].index-queue[0].index+1)

		}
	}
	return res
}
