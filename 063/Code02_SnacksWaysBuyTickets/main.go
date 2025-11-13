package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

// 牛牛的背包问题 & 世界冰球锦标赛
// 测试链接 : https://www.luogu.com.cn/problem/P4799
// 测试链接 : https://www.nowcoder.com/practice/bf877f837467488692be703735db84e6
// 单价和总金额太大，背包不能解决。特征 1≤N≤40
const MAXN = 1 << 20

var lsum []int = make([]int, MAXN)
var rsum []int = make([]int, MAXN)
var cost [40]int
var lIndex int
var rIndex int
var n, m int

func build() {
	for i := 0; i < MAXN; i++ {
		lsum[i] = 0
		rsum[i] = 0
	}
	for i := range cost {
		cost[i] = 0
	}
	lIndex = 0
	rIndex = 0
}

func dfs(index int, end int, sum int, sumArr []int, sumIndex *int) {
	if index >= end {
		if sum <= m {
			sumArr[*sumIndex] = sum
			*sumIndex++
		}
		return
	}
	dfs(index+1, end, sum, sumArr, sumIndex)
	dfs(index+1, end, sum+cost[index], sumArr, sumIndex)
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
	for i := 0; i < n; i++ {
		scanner.Scan()
		cost[i], _ = strconv.Atoi(scanner.Text())
	}
	dfs(0, n/2, 0, lsum, &lIndex)
	dfs(n/2, n, 0, rsum, &rIndex)
	lpart := lsum[:lIndex]
	rpart := rsum[:rIndex]
	sort.Ints(lpart)
	sort.Ints(rpart)
	res := 0
	for i, j := lIndex-1, 0; i >= 0; i-- {
		for j < rIndex && lpart[i]+rpart[j] <= m {
			j++
		}
		res += j
	}
	writer.WriteString(strconv.Itoa(res))
}
