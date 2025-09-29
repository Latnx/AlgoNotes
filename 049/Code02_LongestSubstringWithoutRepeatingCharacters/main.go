package main

// 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// 测试链接 : https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	res := 0
	set := [256]bool{}
	for l, r := 0, 0; r < len(s); r++ {
		for set[s[r]] {
			set[s[l]] = false
			l++
		}
		set[s[r]] = true
		res = max(res, r-l+1)
	}
	return res
}
