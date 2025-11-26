package main

// 环绕字符串中唯一的子字符串
// 测试链接 : https://leetcode.cn/problems/unique-substrings-in-wraparound-string/
func findSubstringInWraproundString(s string) int {
	dic := make([]int, 26)
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i]-1 || s[i-1] == s[i]+25 {
			dp[i] = dp[i-1] + 1
		}
		dic[s[i]-'a'] = max(dp[i], dic[s[i]-'a'])
	}
	sum := 0
	for _, val := range dic {
		sum += val
	}
	return sum
}
