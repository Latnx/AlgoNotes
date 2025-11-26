package main

// 最长有效括号
// 测试链接 : https://leetcode.cn/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	maxLength := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			p := i - 1 - dp[i-1]
			if p >= 0 {
				if s[p] == ')' {
					dp[i] = 0
				} else if p-1 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[p-1]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxLength = max(maxLength, dp[i])
	}
	return maxLength
}
func main() {
	longestValidParentheses(")()())")
}
