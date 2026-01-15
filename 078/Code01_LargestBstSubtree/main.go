package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// 最大BST子树
// 给定一个二叉树，找到其中最大的二叉搜索树（BST）子树，并返回该子树的大小
// 其中，最大指的是子树节点数最多的
// 二叉搜索树（BST）中的所有节点都具备以下属性：
// 左子树的值小于其父（根）节点的值
// 右子树的值大于其父（根）节点的值
// 注意：子树必须包含其所有后代
// 测试链接 : https://leetcode.cn/problems/largest-bst-subtree/
// https://www.nowcoder.com/practice/380d49d7f99242709ab4b91c36bf2acc
var left, right, parent []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	root, _ := strconv.Atoi(scanner.Text())

	left = make([]int, n+1)
	right = make([]int, n+1)
	parent = make([]int, n+1)

	for i := 0; i < n; i++ {
		//  fa，lch，rch，表示 fa 的左儿子为 lch，右儿子为 rch
		scanner.Scan()
		fa, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		lch, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		rch, _ := strconv.Atoi(scanner.Text())
		parent[fa] = fa
		left[fa] = lch
		right[fa] = rch
	}
	writer.WriteString(strconv.Itoa(Bst(root)[0]))
}

// []bst节点数，最小值，最大值, 0/1bst
func Bst(root int) [4]int {
	if root == 0 {
		return [4]int{0, math.MaxInt32, math.MinInt32, 1}
	}
	bstL := Bst(left[root])
	bstR := Bst(right[root])
	// 1. 不包含x
	// 2. 包含x
	if bstL[3] == 1 && bstR[3] == 1 && bstL[2] < parent[root] && bstR[1] > parent[root] {
		return [4]int{bstL[0] + bstR[0] + 1, min(bstL[1], root), max(bstR[2], root), 1}
	} else {
		return [4]int{max(bstL[0], bstR[0]), math.MaxInt32, math.MinInt32, 0}
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
