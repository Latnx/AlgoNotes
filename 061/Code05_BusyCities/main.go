package main

// 繁忙的都市
// 测试链接 : https://www.luogu.com.cn/problem/P2330

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

var arr []int

const MAXM = 200001

var edges [][3]int = make([][3]int, MAXM)

func find(i int) int {
	if arr[i] != i {
		arr[i] = find(arr[i])
	}
	return arr[i]
}

func union(x, y int) {
	arr[find(x)] = arr[find(y)]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	for i := range m {
		scanner.Scan()
		edges[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		edges[i][1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		edges[i][2], _ = strconv.Atoi(scanner.Text())
	}
	slices.SortFunc(edges, func(a, b [3]int) int { return a[2] - b[2] })
	arr = make([]int, n+1)
	for i := range arr {
		arr[i] = i
	}

	ans := 0
	edgeCnt := 0
	for _, edge := range edges {
		if find(edge[0]) != find(edge[1]) {
			union(edge[0], edge[1])
			edgeCnt++
			ans = max(ans, edge[2])
		}
	}
	writer.WriteString(strconv.Itoa(edgeCnt) + " " + strconv.Itoa(ans))
}
