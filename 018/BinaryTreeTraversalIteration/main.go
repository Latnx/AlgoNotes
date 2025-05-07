package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 一个栈，先弹出再打印，先压右再压左
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	sta := [](*TreeNode){}
	sta = append(sta, root)

	for len(sta) != 0 && sta[len(sta)-1] != nil {
		p := sta[len(sta)-1]
		sta = sta[0 : len(sta)-1]
		res = append(res, p.Val)
		if p.Right != nil {
			sta = append(sta, p.Right)
		}
		if p.Left != nil {
			sta = append(sta, p.Left)
		}
	}
	return res
}

// 一个栈，整条左边界进栈，弹出节点打印，右树重复
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	sta := [](*TreeNode){}
	head := root
	for len(sta) != 0 || head != nil {
		if head == nil {
			head = sta[len(sta)-1]
			sta = sta[:len(sta)-1]
			res = append(res, head.Val)
			head = head.Right
		} else {
			sta = append(sta, head)
			head = head.Left
		}
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	sta := []*TreeNode{}
	sta = append(sta, root)
	sta2 := []*TreeNode{}
	for len(sta) != 0 {
		cur := sta[len(sta)-1]
		sta2 = append(sta2, cur)
		sta = sta[:len(sta)-1]
		if cur.Left != nil {
			sta = append(sta, cur.Left)
		}
		if cur.Right != nil {
			sta = append(sta, cur.Right)
		}
	}
	for len(sta2) != 0 {
		res = append(res, sta2[len(sta2)-1].Val)
		sta2 = sta2[:len(sta2)-1]
	}
	return res
}
