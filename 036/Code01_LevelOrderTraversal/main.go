package main

// 二叉树的层序遍历
// 测试链接 : https://leetcode.cn/problems/binary-tree-level-order-traversal/
// 不是普通的层序遍历，需要按层分组
// 需要记录层信息(a, 0), 使用哈希表
// 在队列中记录层信息是否可行？
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
	var l []*node
	nowLevel := -1
	res := [][]int{}
	l = append(l, &node{root, 0})
	for len(l) != 0 {
		curNode := l[0]
		if curNode.level != nowLevel {
			res = append(res, []int{})
			nowLevel = curNode.level
		}
		res[len(res)-1] = append(res[len(res)-1], curNode.Val)

		if curNode.Left != nil {
			l = append(l, &node{curNode.Left, curNode.level + 1})
		}
		if curNode.Right != nil {
			l = append(l, &node{curNode.Right, curNode.level + 1})
		}
		l = l[1:]
	}
	return res
}
