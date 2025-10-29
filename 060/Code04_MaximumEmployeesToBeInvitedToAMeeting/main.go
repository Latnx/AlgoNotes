package main

// 参加会议的最多员工数
// 测试链接 : https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/
// 基环图
// 1. 小环：所有小环的两个顶点的最长链
// 2. 若安排大环，只能有一个大环（最大）
// 流程：
// 1. 拓扑排序判环
// 计算环个数中点数若小环则+= 两顶点最大，若大环则计算大环节点数
func maximumInvitations(favorite []int) int {
	indegree := make([]int, len(favorite))
	preLenth := make([]int, len(favorite)) // 判断小环的长链数
	queue := make([]int, len(favorite))
	l, r := 0, 0
	for i := range favorite {
		indegree[favorite[i]]++
		preLenth[i] = 1
	}
	for i := range indegree {
		if indegree[i] == 0 {
			queue[r] = i
			r++
		}
	}
	for l < r {
		preLenth[favorite[queue[l]]] = max(preLenth[favorite[queue[l]]], preLenth[queue[l]]+1)
		indegree[favorite[queue[l]]]--
		if indegree[favorite[queue[l]]] == 0 {
			queue[r] = favorite[queue[l]]
			r++
		}
		l++
	}
	// 现在所有未成环的都已遍历
	// 统计小环和大环
	resMinC := 0
	resMaxC := 0
	for i := range indegree {
		// 小环
		if indegree[i] != 0 && favorite[favorite[i]] == i {
			indegree[i]--
			indegree[favorite[i]]--
			resMinC += preLenth[i] + preLenth[favorite[i]]
		} else if indegree[i] != 0 && favorite[favorite[i]] != i {
			resMaxCTemp := 0
			for ; indegree[i] != 0; i = favorite[i] {
				indegree[i]--
				resMaxCTemp++
			}
			resMaxC = max(resMaxCTemp, resMaxC)
		}
	}
	return max(resMinC, resMaxC)
}
