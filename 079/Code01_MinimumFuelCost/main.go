package main

// 到达首都的最少油耗
// 给你一棵 n 个节点的树（一个无向、连通、无环图）
// 每个节点表示一个城市，编号从 0 到 n - 1 ，且恰好有 n - 1 条路
// 0 是首都。给你一个二维整数数组 roads
// 其中 roads[i] = [ai, bi] ，表示城市 ai 和 bi 之间有一条 双向路
// 每个城市里有一个代表，他们都要去首都参加一个会议
// 每座城市里有一辆车。给你一个整数 seats 表示每辆车里面座位的数目
// 城市里的代表可以选择乘坐所在城市的车，或者乘坐其他城市的车
// 相邻城市之间一辆车的油耗是一升汽油
// 请你返回到达首都最少需要多少升汽油
// 测试链接 : https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital/
func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], road[1])
		graph[road[1]] = append(graph[road[1]], road[0])
	}
	_, cost := dfs(graph, seats, 0, -1)
	return int64(cost)
}

// 用图建树，从父到子遍历
func dfs(graph [][]int, seats, now, parent int) (int, int) {
	sizeN, costN := 1, 0
	for _, v := range graph[now] {
		if v != parent {
			size, cost := dfs(graph, seats, v, now)
			sizeN += size
			costN += cost
			costN += (size + seats - 1) / seats
		}
	}
	return sizeN, costN
}
