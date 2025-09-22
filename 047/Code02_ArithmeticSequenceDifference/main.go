package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 一开始1~n范围上的数字都是0，一共有m个操作，每次操作为(l,r,s,e,d)
// 表示在l~r范围上依次加上首项为s、末项为e、公差为d的数列
// m个操作做完之后，统计1~n范围上所有数字的最大值和异或和
// 测试链接 : https://www.luogu.com.cn/problem/P4231
// 第一行 2 个整数 N,M，用空格隔开，下同。
// 接下来 M 行，每行4个整数 l,r,s,e，含义见题目描述。
var arr []int
var N, M int

func set(l, r, s, e, d int) {
	arr[l] += s
	arr[l+1] += d - s
	arr[r+1] -= d + e
	arr[r+2] += e
}
func build() {
	for i := 1; i <= N; i++ {
		arr[i] += arr[i-1]
	}
	for i := 1; i <= N; i++ {
		arr[i] += arr[i-1]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fmt.Scan(&N, &M)
	arr = make([]int, N+3)
	for range M {
		scanner.Scan()
		l, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		r, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		s, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		e, _ := strconv.Atoi(scanner.Text())
		set(l, r, s, e, (e-s)/(r-l))
	}

	build()
	ma := 0
	Xor := 0
	for i := 1; i <= N; i++ {
		ma = max(ma, arr[i])
		Xor ^= arr[i]
	}
	fmt.Println(Xor, ma)
}
