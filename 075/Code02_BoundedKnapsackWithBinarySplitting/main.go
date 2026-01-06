package main

import (
	"bufio"
	"os"
	"strconv"
)

// 多重背包通过二进制分组转化成01背包(模版)
// 宝物筛选
// 一共有n种货物, 背包容量为t
// 每种货物的价值(v[i])、重量(w[i])、数量(c[i])都给出
// 请返回选择货物不超过背包容量的情况下，能得到的最大的价值
// 测试链接 : https://www.luogu.com.cn/problem/P1776
var v, w = [10005]int{}, [40005]int{}
var dp = [40005]int{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	// 虚拟出的总商品数量
	m := 0
	for i := 1; i <= n; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		weight, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		cnt, _ := strconv.Atoi(scanner.Text())

		// 二进制分组
		for k := 1; k <= cnt; k <<= 1 {
			m++
			v[m] = k * val
			w[m] = k * weight
			cnt -= k
		}
		if cnt > 0 {
			m++
			v[m] = cnt * val
			w[m] = cnt * weight
		}
	}
	writer.WriteString(strconv.Itoa(compute(m, t)))
}

// 压缩的01背包
func compute(n, t int) int {
	for i := 1; i <= n; i++ {
		for j := t; j >= w[i]; j-- {
			dp[j] = max(dp[j], dp[j-w[i]]+v[i])
		}
	}
	return dp[t]
}
