package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// 完成配对需要的最少字符数量
// 给定一个由'['、']'、'('，')'组成的字符串
// 请问最少插入多少个括号就能使这个字符串的所有括号正确配对
// 例如当前串是 "([[])"，那么插入一个']'即可满足
// 输出最少需要插入多少个字符
// 测试链接 : https://www.nowcoder.com/practice/e391767d80d942d29e6095a935a5b96b
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	str := scanner.Text()
	dp = make([][]int, len(str))
	for i := range dp {
		dp[i] = make([]int, len(str))
	}
	writer.WriteString(strconv.Itoa(f1([]byte(str), 0, len(str)-1)))

}

var dp [][]int

func f1(str []byte, l, r int) int {
	if dp[l][r] != 0 {
		return dp[l][r]
	}
	// 退出条件
	if l == r {
		return 1
	}
	if l+1 == r {
		if (str[l] == '[' && str[r] == ']') || (str[l] == '(' && str[r] == ')') {
			return 0
		} else {
			return 2
		}
	}
	// 划分
	res := math.MaxInt
	if (str[l] == '[' && str[r] == ']') || (str[l] == '(' && str[r] == ')') {
		res = f1(str, l+1, r-1)
	}
	for k := l; k < r; k++ {
		res = min(res, f1(str, l, k)+f1(str, k+1, r))
	}
	if res == math.MaxInt {
		res = 0
	}
	dp[l][r] = res
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
