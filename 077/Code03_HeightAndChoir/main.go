package main

import (
	"bufio"
	"os"
	"strconv"
)

const MOD int = 19650827

// 合唱队
// 具体描述情打开链接查看
// 测试链接 : https://www.luogu.com.cn/problem/P3205
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, n)
	for i := range n {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}
	dp = make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, n)
	}
	// a, b := f1(arr, 0, n-1)
	writer.WriteString(strconv.Itoa(f2(arr)))
}

// 左端点是最后进来的，右端点是最后进来的
var dp [][][2]int

// [l...r]有多少种方式
func f1(arr []int, l, r int) (int, int) {
	if l == r {
		return 1, 1
	}
	if l+1 == r {
		if arr[l] < arr[r] {
			return 1, 1
		}
		return 0, 0
	}
	//左端点是最后进来的
	a, b := f1(arr, l+1, r)
	c, d := f1(arr, l, r-1)
	resLeft, resRight := 0, 0
	if arr[l] < arr[l+1] {
		resLeft += a
	}
	if arr[l] < arr[r] {
		resLeft += b
	}
	if arr[r] > arr[l] {
		resRight += c
	}
	if arr[r] > arr[r-1] {
		resRight += d
	}
	return resLeft, resRight
}

func f2(arr []int) int {
	n := len(arr)
	for i := range arr {
		dp[i][i] = [2]int{1, 1}
	}
	for i := 0; i+1 < n; i++ {
		if arr[i] < arr[i+1] {
			dp[i][i+1] = [2]int{1, 1}
		} else {
			dp[i][i+1] = [2]int{0, 0}
		}
	}
	for l := n - 3; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			a, b := dp[l+1][r][0], dp[l+1][r][1]
			c, d := dp[l][r-1][0], dp[l][r-1][1]
			resLeft, resRight := 0, 0
			if arr[l] < arr[l+1] {
				resLeft += a
			}
			if arr[l] < arr[r] {
				resLeft += b
			}
			if arr[r] > arr[l] {
				resRight += c
			}
			if arr[r] > arr[r-1] {
				resRight += d
			}
			dp[l][r] = [2]int{resLeft % MOD, resRight % MOD}
		}
	}
	return (dp[0][n-1][0] + dp[0][n-1][1]) % MOD
}
