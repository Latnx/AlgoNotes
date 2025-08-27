package main

import (
	"strconv"
	"strings"
)

// 二叉树先序序列化和反序列化
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
	this.f(root)
	return string(this.code)
}

func (this *Codec) f(root *TreeNode) {
	if root == nil {
		this.code = append(this.code, '#', ',')
		return
	}
	this.code = append(this.code, strconv.Itoa(root.Val)...)
	this.code = append(this.code, ',')
	this.f(root.Left)
	this.f(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.strArr = strings.Split(string(data), ",")
	return this.d()
}
func (this *Codec) d() *TreeNode {
	if this.strArr[0] == "#" {
		this.strArr = this.strArr[1:]
		return nil
	}
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(this.strArr[0])
	this.strArr = this.strArr[1:]
	root.Left = this.d()
	root.Right = this.d()
	return root
}
