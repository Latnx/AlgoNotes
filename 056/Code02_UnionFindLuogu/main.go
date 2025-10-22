package main

// 并查集模版(洛谷)
// 本实现用递归函数实现路径压缩，而且省掉了小挂大的优化，一般情况下可以省略
// 测试链接 : https://www.luogu.com.cn/problem/P3367

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAXN = 1000001

var arr [MAXN]int

func build() {
	for i := 0; i < MAXN; i++ {
		arr[i] = i
	}
}

func find(i int) int {
	if arr[i] != i {
		arr[i] = find(arr[i])
	}
	return arr[i]
}

func isSameSet(a, b int) bool {
	return find(a) == find(b)
}
func union(a, b int) {
	arr[find(a)] = find(b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	N, M := 0, 0
	build()
	fmt.Scan(&N, &M)
	for i := 0; i < M; i++ {
		var opt, x, y int
		scanner.Scan()
		opt, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		x, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ = strconv.Atoi(scanner.Text())
		switch opt {
		case 2:
			if isSameSet(x, y) {
				writer.WriteString("Y\n")
			} else {
				writer.WriteString("N\n")
			}
		case 1:
			union(x, y)
		}
	}
}
