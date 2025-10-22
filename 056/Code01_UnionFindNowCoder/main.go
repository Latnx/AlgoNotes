package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 并查集模版(牛客)
// 路径压缩 + 小挂大
// 测试链接 : https://www.nowcoder.com/practice/e7ed657974934a30b2010046536a5372

// 问题场景:
// 给定一个没有重复值的整形数组arr，初始时认为arr中每一个数各自都是一个单独的集合。请设计一种叫UnionFind的结构，并提供以下两个操作。
// boolean isSameSet(int a, int b): 查询a和b这两个数是否属于一个集合
// void union(int a, int b): 把a所在的集合与b所在的集合合并在一起，原本两个集合各自的元素以后都算作同一个集合
// O(1)
// 小挂大
// 扁平化
const MAXN = 1000001

var arr [MAXN]int
var size [MAXN]int

func build() {
	for i := 0; i < MAXN; i++ {
		arr[i] = i
		size[i] = 1
	}
}

func find(i int) int {
	temp := i
	for arr[i] != i {
		i = arr[i]
	}
	for arr[temp] != temp {
		temp = arr[temp]
		arr[temp] = i
	}
	return i
}

func isSameSet(a, b int) bool {
	return find(a) == find(b)
}
func union(a, b int) {
	findA, findB := find(a), find(b)
	if size[findA] > size[findB] {
		arr[findB] = findA
		size[findB] += size[findA]
	} else {
		arr[findA] = findB
		size[findA] += size[findB]
	}
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
		case 1:
			if isSameSet(x, y) {
				writer.WriteString("Yes\n")
			} else {
				writer.WriteString("No\n")
			}
		case 2:
			union(x, y)
		}
	}
}
