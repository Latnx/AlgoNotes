package main

// 移除最多的同行或同列石头
// n 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头
// 如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头
// 给你一个长度为 n 的数组 stones ，其中 stones[i] = [xi, yi] 表示第 i 块石头的位置
// 返回 可以移除的石子 的最大数量。
// 测试链接 : https://leetcode.cn/problems/most-stones-removed-with-same-row-or-column/

// stones 是稀疏矩阵
var arr []int
var sets int

func find(i int) int {
	if arr[i] != i {
		arr[i] = find(arr[i])
	}
	return arr[i]
}
func inSameSet(x, y int) bool {
	return find(x) == find(y)
}
func union(x, y int) {
	findX, findY := find(x), find(y)
	arr[findX] = find(findY)
	if findX != findY {
		sets--
	}
}
func removeStones(stones [][]int) int {
	arr = make([]int, len(stones))
	for i := range arr {
		arr[i] = i
	}
	sets = len(stones)
	// 做行列映射
	mapCol := make(map[int]int)
	mapRow := make(map[int]int)
	for i, val := range stones {
		x, y := val[0], val[1]
		if index, ok := mapRow[x]; ok {
			union(i, index)
		} else {
			mapRow[x] = i
		}
		if index, ok := mapCol[y]; ok {
			union(i, index)
		} else {
			mapCol[y] = i
		}
	}
	return len(stones) - sets
}
func main() {
	removeStones([][]int{{0, 1}, {1, 0}})
}
