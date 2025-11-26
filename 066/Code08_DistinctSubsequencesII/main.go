package main

// 不同的子序列 II
// 测试链接 : https://leetcode.cn/problems/distinct-subsequences-ii/

// 若字符不同，算空集，非空子序列数量2^n个
// 若字符相同，
// 思路：all = 1;dp[i]有多少个子序列以i为结尾
// 1. 纯新增：all- 当前字符上次记录
// 2. 当前字符记录+=纯新增
// 3. all+=纯新增
const mod = 1000000007

func distinctSubseqII(s string) int {
	all := 1
	dp := make([]int, 26)
	for i := range s {
		c := s[i] - 'a'
		new := all - dp[c]%mod
		dp[c] += new % mod
		all += new % mod
	}
	return (all - 1) % mod
}
