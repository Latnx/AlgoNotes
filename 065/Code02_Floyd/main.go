package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// Floyd算法模版（洛谷）
// 测试链接 : https://www.luogu.com.cn/problem/P2910

const MAXN = 101
const MAXM = 10001

var path = [MAXM]int{}
var distance = [MAXN][MAXN]int{}
var n, m, ans int

func build() {
	for i := range MAXN {
		for j := range MAXN {
			distance[i][j] = math.MaxInt
		}
	}
}

func main() {
	build()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())
	for i := range m {
		scanner.Scan()
		path[i], _ = strconv.Atoi(scanner.Text())
		path[i]--
	}
	for i := range n {
		for j := range n {
			scanner.Scan()
			distance[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}
	floyd()
	for i := 1; i < m; i++ {
		ans += distance[path[i-1]][path[i]]
	}
	writer.WriteString(strconv.Itoa(ans))
}

func floyd() {
	for k := range n {
		for i := range n {
			for j := range n {
				if distance[i][k] != math.MaxInt && distance[k][j] != math.MaxInt {
					distance[i][j] = min(distance[i][j], distance[i][k]+distance[k][j])
				}
			}
		}
	}
}
