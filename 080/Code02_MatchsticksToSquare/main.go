package main

// 火柴拼正方形
// 你将得到一个整数数组 matchsticks
// 其中 matchsticks[i] 是第 i 个火柴棒的长度
// 你要用 所有的火柴棍 拼成一个正方形
// 你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次
// 如果你能拼出正方形，则返回 true ，否则返回 false
// 测试链接 : https://leetcode.cn/problems/matchsticks-to-square/
func makesquare(matchsticks []int) bool {
	sum := 0
	for _, stick := range matchsticks {
		sum += stick
	}
	if sum%4 != 0 {
		return false
	}
	dp = make([]int, 1<<len(matchsticks))
	return f(matchsticks, sum/4, (1<<len(matchsticks))-1, 0, 4)
}

var dp []int

func f(mathsticks []int, bian int, status int, cur, rest int) bool {
	if rest == 0 {
		return status == 0
	}
	if dp[status] != 0 {
		return dp[status] == 1
	}
	ans := false
	for i := range mathsticks {
		if status&(1<<i) != 0 && mathsticks[i]+cur <= bian {
			if mathsticks[i]+cur == bian {
				ans = f(mathsticks, bian, status^(1<<i), 0, rest-1)
			} else {
				ans = f(mathsticks, bian, status^(1<<i), cur+mathsticks[i], rest)
			}
			if ans {
				dp[status] = 1
				return ans
			}
		}
	}
	dp[status] = 2
	return ans
}
