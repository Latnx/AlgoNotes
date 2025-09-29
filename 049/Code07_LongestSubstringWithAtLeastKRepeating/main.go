package main

// 至少有K个重复字符的最长子串
// 给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串
// 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度
// 如果不存在这样的子字符串，则返回 0。
// 测试链接 : https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/

// 双指针问题分解，至少有 K 个重复字符的最长子串，这个子串有1种字符，2种字符
// 数量不够就扩窗口，种类超过就缩窗口
func longestSubstring(s string, k int) int {
	res := 0
	for require := 1; require <= 26; require++ {
		cnts := make([]int, 256)
		category := 0
		satisfy := 0
		for l, r := 0, 0; r < len(s); r++ {
			if cnts[s[r]] == 0 {
				category++
			}
			cnts[s[r]]++
			if cnts[s[r]] == k {
				satisfy++
			}
			for ; l <= r && category > require; l++ {
				if cnts[s[l]] == k {
					satisfy--
				}
				cnts[s[l]]--
				if cnts[s[l]] == 0 {
					category--
				}
			}
			if category == satisfy {
				res = max(res, r-l+1)
			}
		}
	}
	return res
}
