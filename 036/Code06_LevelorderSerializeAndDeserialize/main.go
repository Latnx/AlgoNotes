package main

import (
	"strconv"
	"strings"
)

// 二叉树按层序列化和反序列化
// 测试链接 : https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
	code   []byte
	strArr []string
}

func Constructor() Codec {
	return Codec{
		code: make([]byte, 0),
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	this.code = append(this.code, strconv.Itoa(root.Val)...)
	this.code = append(this.code, ',')
	for len(queue) != 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
			this.code = append(this.code, strconv.Itoa(queue[0].Left.Val)...)
			this.code = append(this.code, ',')
		} else {
			this.code = append(this.code, '#', ',')
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
			this.code = append(this.code, strconv.Itoa(queue[0].Right.Val)...)
			this.code = append(this.code, ',')

		} else {
			this.code = append(this.code, '#', ',')
		}
		queue = queue[1:]
	}
	return string(this.code)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.strArr = strings.Split(string(data), ",")
	queue := []*TreeNode{}
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(this.strArr[0])
	queue = append(queue, root)
	for i := 1; i < len(this.strArr) && len(queue) != 0; i += 2 {
		root := queue[0]
		if this.strArr[i] != "#" {
			left := &TreeNode{}
			left.Val, _ = strconv.Atoi(this.strArr[i])
			root.Left = left
			queue = append(queue, left)
		}
		if this.strArr[i+1] != "#" {
			right := &TreeNode{}
			right.Val, _ = strconv.Atoi(this.strArr[i+1])
			root.Right = right
			queue = append(queue, right)
		}
		queue = queue[1:]
	}
	return root
}
