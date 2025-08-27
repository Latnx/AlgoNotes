package main

// 验证完全二叉树
// 测试链接 : https://leetcode.cn/problems/check-completeness-of-a-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	flag := true
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		curNode := queue[0]
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
			if !flag {
				return flag
			}
		} else {
			flag = false
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
			if !flag {
				return flag
			}
		} else {
			flag = false
		}
		queue = queue[1:]
	}
	return true
}
