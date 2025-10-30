package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

// Prim算法模版（洛谷）
// 动态空间实现
// 测试链接 : https://www.luogu.com.cn/problem/P3366
const MAXM = 200001

type hp [][2]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(i any)        { *h = append(*h, i.([2]int)) }
func (h *hp) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	adj := make([][][2]int, n+1)
	for i := range adj {
		adj[i] = make([][2]int, 0)
	}
	h := make(hp, 0)
	for range m {
		scanner.Scan()
		from, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		to, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		adj[from] = append(adj[from], [2]int{to, val})
		adj[to] = append(adj[to], [2]int{from, val})

	}
	for j := range adj[1] {
		heap.Push(&h, adj[1][j])
	}
	ans := 0
	edgeCnt := 0
	set := make([]bool, n+1)
	set[1] = true
	for len(h) != 0 {
		edge := heap.Pop(&h).([2]int)
		if set[edge[0]] {
			continue
		}
		set[edge[0]] = true
		ans += edge[1]
		edgeCnt++
		for j := range adj[edge[0]] {
			heap.Push(&h, adj[edge[0]][j])
		}
	}
	if edgeCnt == n-1 {
		writer.WriteString(strconv.Itoa(ans))
	} else {
		writer.WriteString("orz")
	}
}
