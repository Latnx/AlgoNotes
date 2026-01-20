package main

import "math"

// 从树中删除边的最小分数
// 存在一棵无向连通树，树中有编号从0到n-1的n个节点，以及n-1条边
// 给你一个下标从0开始的整数数组nums长度为n，其中nums[i]表示第i个节点的值
// 另给你一个二维整数数组edges长度为n-1
// 其中 edges[i] = [ai, bi] 表示树中存在一条位于节点 ai 和 bi 之间的边
// 删除树中两条不同的边以形成三个连通组件，对于一种删除边方案，定义如下步骤以计算其分数：
// 分别获取三个组件每个组件中所有节点值的异或值
// 最大 异或值和 最小 异或值的 差值 就是这种删除边方案的分数
// 返回可能的最小分数
// 测试链接 : https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/

// 可以看作，将一个节点视作根。可求得整棵树的异或和
// 再选择两个节点与整棵树断开。分别求两个节点子树的异或和即可
// 因为是异或，可以先记录再进行计算来进行化简
// 计算每个节点的子树异或和（dfs）
// 分别尝试断开不同节点
// 1. 若两个节点不包含(使用dfn序列，判断)，则删掉的异或和=A^b^c
// 2. 若包含则
func minimumScore(nums []int, edges [][]int) int {
	clear(dfn)
	clear(size)
	time = 0
	n := len(nums)
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	dfs(graph, nums, 0, -1)
	ans := math.MaxInt
	for i := 0; i < len(edges)-1; i++ {
		for j := i + 1; j < len(edges); j++ {
			u1, v1, u2, v2 := edges[i][0], edges[i][1], edges[j][0], edges[j][1]
			edge1 := max(dfn[u1], dfn[v1])
			edge2 := max(dfn[u2], dfn[v2])
			pre := min(edge1, edge2)
			pos := max(edge1, edge2)
			// 如果包含
			var a, b, c int
			total := xor[dfn[0]]
			if pos < pre+size[pre] {
				c = xor[pos]
				b = xor[pre] ^ c
				a = total ^ xor[pre]
			} else {
				c = xor[pos]
				b = xor[pre]
				a = total ^ b ^ c
			}
			ans = min(ans, max(a, b, c)-min(a, b, c))
		}
	}
	return ans
}

var time = 0
var dfn = make([]int, 1000)
var size = make([]int, 1000)
var xor = make([]int, 1000)

func dfs(graph [][]int, nums []int, now, p int) (int, int) {
	cnt := time
	time++
	dfn[now] = cnt
	s, or := 1, nums[now]
	for i := range graph[now] {
		if graph[now][i] == p {
			continue
		}
		a, b := dfs(graph, nums, graph[now][i], now)
		s += a
		or ^= b
	}
	size[cnt], xor[cnt] = s, or
	return s, or
}
