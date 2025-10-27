package main

import "sort"

// 找出知晓秘密的所有专家
// 链接测试 : https://leetcode.cn/problems/find-all-people-with-secret/
type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][2] < a[j][2] }

var arr []int

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
	arr[find(x)] = find(y)
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	arr = make([]int, n)
	for i := range n {
		arr[i] = i
	}
	sort.Sort(SortBy(meetings))
	union(0, firstPerson)
	time := 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i][2] != time {
			for j := i - 1; j >= 0 && meetings[j][2] == time; j-- {
				if !inSameSet(meetings[j][0], firstPerson) {
					arr[meetings[j][0]] = meetings[j][0]
				}
				if !inSameSet(meetings[j][1], firstPerson) {
					arr[meetings[j][1]] = meetings[j][1]
				}
			}
		}
		time = meetings[i][2]
		union(meetings[i][1], meetings[i][0])
	}
	res := make([]int, 0)
	for i := range arr {
		if inSameSet(i, firstPerson) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	findAllPeople(5, [][]int{{1, 4, 3}, {0, 4, 3}}, 3)
}
