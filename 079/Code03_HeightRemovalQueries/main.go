package main

// 移除子树后的二叉树高度
// 给你一棵 二叉树 的根节点 root ，树中有 n 个节点
// 每个节点都可以被分配一个从 1 到 n 且互不相同的值
// 另给你一个长度为 m 的数组 queries
// 你必须在树上执行 m 个 独立 的查询，其中第 i 个查询你需要执行以下操作：
// 从树中 移除 以 queries[i] 的值作为根节点的子树
// 题目所用测试用例保证 queries[i] 不等于根节点的值
// 返回一个长度为 m 的数组 answer
// 其中 answer[i] 是执行第 i 个查询后树的高度
// 注意：
// 查询之间是独立的，所以在每个查询执行后，树会回到其初始状态
// 树的高度是从根到树中某个节点的 最长简单路径中的边数
// 测试链接 : https://leetcode.cn/problems/height-of-binary-tree-after-subtree-removal-queries/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var n = 100000
var dfn = make([]int, n)
var size = make([]int, n)
var deep = make([]int, n)
var maxL = make([]int, n)
var maxR = make([]int, n)

func treeQueries(root *TreeNode, queries []int) []int {
	// 1. val->dfn
	// 2. dfn->子树大小
	// 3. dfn->深度
	clear(dfn)
	clear(size)
	clear(deep)
	time = 0
	dfs(root, 0)
	// 记录左右最大值，方便查找
	maxL[0] = deep[0]
	for i := 1; i < time; i++ {
		maxL[i] = max(maxL[i-1], deep[i])
	}
	maxR[time] = deep[time]
	for i := time - 1; i >= 0; i-- {
		maxR[i] = max(maxR[i+1], deep[i])
	}

	res := make([]int, len(queries))
	for i := range queries {
		index := dfn[queries[i]]
		res[i] = max(maxL[index-1], maxR[index+size[index]])
	}
	return res
}

var time = 0

func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return 0
	}
	cur := time
	dfn[root.Val] = cur
	deep[cur] = depth
	time++

	leftSize := dfs(root.Left, depth+1)
	rightSize := dfs(root.Right, depth+1)

	size[cur] = leftSize + rightSize + 1
	return size[cur]
}
