package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// 涂色 & 奇怪打印机
// 假设你有一条长度为5的木板，初始时没有涂过任何颜色
// 你希望把它的5个单位长度分别涂上红、绿、蓝、绿、红
// 用一个长度为5的字符串表示这个目标：RGBGR
// 每次你可以把一段连续的木板涂成一个给定的颜色，后涂的颜色覆盖先涂的颜色
// 例如第一次把木板涂成RRRRR
// 第二次涂成RGGGR
// 第三次涂成RGBGR，达到目标
// 返回尽量少的涂色次数
// 测试链接 : https://www.luogu.com.cn/problem/P4170
// 测试链接 : https://leetcode.cn/problems/strange-printer/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	str := scanner.Text()
	dp = make([][]int, len(str))
	for i := range dp {
		dp[i] = make([]int, len(str))
	}
	writer.WriteString(strconv.Itoa(f1([]byte(str), 0, len(str)-1)))
}

var dp [][]int

// [l...r]需要变成指定颜色
func f1(str []byte, l, r int) int {
	if dp[l][r] != 0 {
		return dp[l][r]
	}
	// 退出条件
	if l == r {
		return 1
	}
	if l+1 == r {
		if str[l] == str[r] {
			return 1
		} else {
			return 2
		}
	}
	// 划分
	res := math.MaxInt
	if str[l] == str[r] {
		res = min(res, f1(str, l, r-1))
	}
	for k := l; k < r; k++ {
		res = min(res, f1(str, l, k)+f1(str, k+1, r))
	}
	dp[l][r] = res
	return res
}
