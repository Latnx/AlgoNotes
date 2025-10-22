package main

// 相似字符串组
// 测试链接 : https://leetcode.cn/problems/similar-string-groups/
var arr []int
var sets int

func find(i int) int {
	if arr[i] != i {
		arr[i] = find(arr[i])
	}
	return arr[i]
}
func union(x, y int) {
	findX, findY := find(x), find(y)
	if findX != findY {
		sets--
	}
	arr[findX] = findY
}
func isSimilar(str1, str2 string) bool {
	diff := 0
	for i := range str1 {
		if str1[i] != str2[i] {
			diff++
		}
	}
	return diff == 0 || diff == 2
}
func numSimilarGroups(strs []string) int {
	arr = make([]int, len(strs))
	for i := range arr {
		arr[i] = i
	}
	sets = len(strs)

	for i := 0; i < len(strs)-1; i++ {
		for j := i + 1; j < len(strs); j++ {
			if isSimilar(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}
	return sets
}
