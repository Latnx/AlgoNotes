package main

// 每日温度
// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer
// 其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 测试链接 : https://leetcode.cn/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, len(temperatures))
	r := 0
	answerIndex := make([]int, len(temperatures))
	for i, val := range temperatures {
		for r > 0 && val >= temperatures[stack[r-1]] {
			r--
			answerIndex[stack[r]] = i
		}
		stack[r] = i
		r++
	}
	for r > 0 {
		r--
		answerIndex[stack[r]] = -1
	}
	for i := len(temperatures) - 1; i >= 0; i-- {
		if answerIndex[i] != -1 && temperatures[i] == temperatures[answerIndex[i]] {
			answerIndex[i] = answerIndex[answerIndex[i]]
		}
	}
	answer := make([]int, len(answerIndex))
	for i, val := range answerIndex {
		if val != -1 {
			answer[i] = val - i
		}
	}
	return answerIndex
}
