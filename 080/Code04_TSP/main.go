package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// 售货员的难题 - TSP问题
// 某乡有n个村庄(1<=n<=20)，有一个售货员，他要到各个村庄去售货
// 各村庄之间的路程s(1<=s<=1000)是已知的
// 且A村到B村的路程，与B到A的路大多不同(有向带权图)
// 为了提高效率，他从商店出发到每个村庄一次，然后返回商店所在的村，
// 假设商店所在的村庄为1
// 他不知道选择什么样的路线才能使所走的路程最短
// 请你帮他选择一条最短的路
// 测试链接 : https://www.luogu.com.cn/problem/P1171
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for i := range n {
		for j := range n {
			scanner.Scan()
			graph[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}
	for i := range 1 << (n - 1) {
		for j := range n {
			dp[i][j] = -1
		}
	}
	writer.WriteString(strconv.Itoa(TSP(0, 0)))
}

var n int
var graph [20][20]int
var dp [1 << 19][20]int

func TSP(status, index int) int {
	if status == (1<<(n-1))-1 {
		return graph[index][0]
	}
	if dp[status][index] != -1 {
		return dp[status][index]
	}
	res := math.MaxInt
	for i := 1; i < n; i++ {
		if status&(1<<(i-1)) == 0 {
			res = min(res, TSP(status|(1<<(i-1)), i)+graph[index][i])
		}
	}
	dp[status][index] = res
	return res
}
