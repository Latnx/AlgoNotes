package main

import "fmt"

// 布尔运算
// 给定一个布尔表达式和一个期望的布尔结果 result
// 布尔表达式由 0 (false)、1 (true)、& (AND)、 | (OR) 和 ^ (XOR) 符号组成
// 布尔表达式一定是正确的，不需要检查有效性
// 但是其中没有任何括号来表示优先级
// 你可以随意添加括号来改变逻辑优先级
// 目的是让表达式能够最终得出result的结果
// 返回最终得出result有多少种不同的逻辑计算顺序
// 测试链接 : https://leetcode.cn/problems/boolean-evaluation-lcci/
func countEval(s string, result int) int {

	dp = make([][][2]int, len(s))
	for i := range dp {
		dp[i] = make([][2]int, len(s))
	}
	return dfs([]byte(s), 0, len(s)-1)[result]
}

var dp [][][2]int

func dfs(s []byte, l, r int) [2]int {
	if dp[l][r][0] != 0 || dp[l][r][1] != 0 {
		return dp[l][r]
	}
	if l == r {
		if s[l] == '0' {
			return [2]int{1, 0}
		} else {
			return [2]int{0, 1}
		}
	}
	res := [2]int{}
	for i := l + 1; i < r; i += 2 {
		pre := dfs(s, l, i-1)
		end := dfs(s, i+1, r)
		switch s[i] {
		case '&':
			res[1] += pre[1] * end[1]
			res[0] += pre[0]*end[1] + pre[1]*end[0] + pre[0]*end[0]
		case '|':
			res[1] += pre[1]*end[1] + pre[0]*end[1] + pre[1]*end[0]
			res[0] += pre[0] * end[0]
		case '^':
			res[1] += pre[0]*end[1] + pre[1]*end[0]
			res[0] += pre[1]*end[1] + pre[0]*end[0]
		}
	}
	dp[l][r] = res
	return res
}
func main() {
	fmt.Print(countEval("1^0|0|1", 0))
}
