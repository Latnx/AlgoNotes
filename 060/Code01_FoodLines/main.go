package main

import (
	"bufio"
	"os"
	"strconv"
)

// 最大食物链计数
// a -> b，代表a在食物链中被b捕食
// 给定一个有向无环图，返回
// 这个图中从最初级动物到最顶级捕食者的食物链有几条
// 测试链接 : https://www.luogu.com.cn/problem/P4017
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	indegree := make([]int, n+1)
	for range m {
		scanner.Scan()
		A, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		B, _ := strconv.Atoi(scanner.Text())
		adj[A] = append(adj[A], B)
		indegree[B]++
	}
	resInt := make([]int, n+1)
	queue := make([]int, n+1)
	l, r := 0, 0
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			resInt[i] = 1
			queue[r] = i
			r++
		}
	}
	for l < r {
		for j := range adj[queue[l]] {
			resInt[adj[queue[l]][j]] = (resInt[adj[queue[l]][j]] + resInt[queue[l]]) % 80112002
			indegree[adj[queue[l]][j]]--
			if indegree[adj[queue[l]][j]] == 0 {
				queue[r] = adj[queue[l]][j]
				r++
			}
		}
		l++
	}
	res := 0
	for i := 1; i <= n; i++ {
		if len(adj[i]) == 0 { // 出度为 0
			res = (res + resInt[i]) % 80112002
		}
	}

	writer.WriteString(strconv.Itoa(res))
}
