package main

import (
	"bufio"
	"os"
	"strconv"
)

// 夏季特惠
// 某公司游戏平台的夏季特惠开始了，你决定入手一些游戏
// 现在你一共有X元的预算，平台上所有的 n 个游戏均有折扣
// 标号为 i 的游戏的原价a_i元，现价只要b_i元
// 也就是说该游戏可以优惠 a_i - b_i，并且你购买该游戏能获得快乐值为w_i
// 由于优惠的存在，你可能做出一些冲动消费导致最终买游戏的总费用超过预算
// 只要满足 : 获得的总优惠金额不低于超过预算的总金额
// 那在心理上就不会觉得吃亏。
// 现在你希望在心理上不觉得吃亏的前提下，获得尽可能多的快乐值。
// 测试链接 : https://leetcode.cn/problems/tJau2o/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	w := make([]int, 1)
	v := make([]int, 1)

	ans := 0
	for range n {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		if 2*b-a <= 0 {
			ans += val
			x -= 2*b - a
			continue
		}
		w = append(w, 2*b-a)
		v = append(v, val)
	}
	M := len(w) - 1
	dp := make([][]int, M+1)
	for i := range dp {
		dp[i] = make([]int, x+1)
	}
	for i := 1; i <= M; i++ {
		for j := 1; j <= x; j++ {
			if j >= w[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	writer.WriteString(strconv.Itoa(ans + dp[M][x]))
}
