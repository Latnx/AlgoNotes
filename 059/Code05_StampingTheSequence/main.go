package main

import (
	"fmt"
	"slices"
)

// 戳印序列
// 测试链接 : https://leetcode.cn/problems/stamping-the-sequence/
func movesToStamp(stamp string, target string) []int {
	// 遍历每个位置，记录错误点，所以可以在每个位置统计错误点，盖上后一次，可以取消前一次的错误点
	// 1. 建图
	// 节点 = 印章起点
	// 边   = “错误位置 → 印章起点”
	// 入度 = “该印章还有多少字符不匹配”
	n, m := len(target), len(stamp)
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	indegree := make([]int, n-m+1)
	for i := range indegree {
		indegree[i] = m
	}
	// 存放度为零的初始位置
	queue := make([]int, n-m+1)
	l, r := 0, 0
	// 2. 错误位置 → 印章起点
	for i := range n - m + 1 {
		for j := range stamp {
			if target[i+j] == stamp[j] {
				indegree[i]--
				if indegree[i] == 0 {
					queue[r] = i
					r++
				}
			} else {
				adj[i+j] = append(adj[i+j], i)
			}
		}
	}
	// 3. 拓扑排序, 逐个解放初始位置所在的每一个位置
	visited := make([]bool, n)
	res := make([]int, 0)
	for l < r {
		res = append(res, queue[l])
		for j := range stamp {
			for i := range adj[queue[l]+j] {
				if !visited[queue[l]+j] {
					indegree[adj[queue[l]+j][i]]--
					if indegree[adj[queue[l]+j][i]] == 0 {
						queue[r] = adj[queue[l]+j][i]
						r++
					}
				}
			}
			visited[queue[l]+j] = true
		}
		l++
	}
	if len(res) != m-n+1 {
		return []int{}
	}
	slices.Reverse(res)
	return res
}
func main() {
	fmt.Println(movesToStamp("abc", "abcbc"))    // 输出 [2 0]
	fmt.Println(movesToStamp("abca", "aabcaca")) // 输出 [3 0 1]
}
