package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const OFFSET = 30001
const MAXN = 1000001

var arr []int = make([]int, OFFSET+MAXN+OFFSET)
var n, m int

func set(l, r, s, e, d int) {
	arr[l] += s
	arr[l+1] += -s + d
	arr[r+1] += -d - e
	arr[r+2] += e
}

func build() {
	for i := 1; i <= OFFSET+m; i++ {
		arr[i] += arr[i-1]
	}
	for i := 1; i <= OFFSET+m; i++ {
		arr[i] += arr[i-1]
	}
}
func setP(v, x int) {
	// 统计四段等比
	set(x-3*v+1, x-2*v, 1, v, 1)
	set(x-2*v+1, x, v-1, -v, -1)
	set(x+1, x+2*v, -v+1, v, 1)
	set(x+2*v+1, x+3*v-1, v-1, 1, -1)
}

// 一群人落水后求每个位置的水位高度
// 问题描述比较复杂，见测试链接
// 测试链接 : https://www.luogu.com.cn/problem/P5026
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	// 数目与湖泊的宽度
	fmt.Scan(&n, &m)
	for range n {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		setP(v, x+OFFSET)
	}
	build()
	for i := 1; i <= m; i++ {
		fmt.Print(arr[i+OFFSET], " ")
	}
}
