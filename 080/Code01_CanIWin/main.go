package main

// 我能赢吗
// 给定两个整数n和m
// 两个玩家可以轮流从公共整数池中抽取从1到n的整数（不放回）
// 抽取的整数会累加起来（两个玩家都算）
// 谁在自己的回合让累加和 >= m，谁获胜
// 若先出手的玩家能稳赢则返回true，否则返回false
// 假设两位玩家游戏时都绝顶聪明，可以全盘为自己打算
// 测试链接 : https://leetcode.cn/problems/can-i-win/
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal == 0 {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	dp := make([]int, 1<<(maxChoosableInteger+1))
	return f(maxChoosableInteger, 0, desiredTotal, dp)
}
func f(n, status, rest int, dp []int) bool {
	if rest <= 0 {
		return false
	}
	if dp[status] != 0 {
		return dp[status] == 1
	}
	for i := 1; i <= n; i++ {
		// 没被使用过
		if status>>i&1 == 0 {
			if !f(n, status|(1<<i), rest-i, dp) {
				dp[status] = 1
				return true
			}
		}
	}
	dp[status] = 2
	return false
}
