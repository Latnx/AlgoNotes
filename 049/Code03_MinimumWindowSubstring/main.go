package main

import (
	"fmt"
	"math"
)

// 最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
// 测试链接 : https://leetcode.cn/problems/minimum-window-substring/

func minWindow1(s string, t string) string {
	resL, resR := 0, 0
	mixLength := math.MaxInt
	target := [256]int{}
	now := [256]int{}
	cnt := 0

	for index := range t {
		target[t[index]]++
	}

	// 双指针模板
	for l, r := 0, 0; r < len(s); r++ {
		if target[s[r]] > 0 && now[s[r]] < target[s[r]] {
			cnt++
		}
		now[s[r]]++
		for ; l <= r && now[s[l]] > target[s[l]]; l++ {
			now[s[l]]--
		}
		if cnt == len(t) && mixLength > r-l+1 {
			resL, resR = l, r+1
			mixLength = r - l + 1
		}
	}

	return s[resL:resR]
}

// 负债模型
func minWindow2(s string, t string) string {
	resL, resR := 0, 0
	mixLength := math.MaxInt
	target := [256]int{}
	cnt := 0

	for index := range t {
		target[t[index]]--
	}

	// 双指针模板
	for l, r := 0, 0; r < len(s); r++ {
		if target[s[r]] < 0 {
			cnt++
		}
		target[s[r]]++
		for ; l <= r && target[s[l]] > 0; l++ {
			target[s[l]]--
		}
		if cnt == len(t) && mixLength > r-l+1 {
			resL, resR = l, r+1
			mixLength = r - l + 1
		}
	}
	return s[resL:resR]
}

func main() {
	fmt.Println(minWindow1("ADOBECODEBANC", "ABC")) // "BANC"
}
