package main

import "math"

// 替换子串得到平衡字符串
// 有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
// 假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
// 给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
// 你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
// 请返回待替换子串的最小可能长度。
// 如果原字符串自身就是一个平衡字符串，则返回 0。
// 测试链接 : https://leetcode.cn/problems/replace-the-substring-for-balanced-string/

// 其实是找出多的字母，比如Q多两个，E多一个，那么只需要找出，2Q和1E的最短序列，就是替换子串的最小可能长度
// 因此将问题转换
func balancedString(s string) int {
	minLenth := math.MaxInt
	cnt := 0
	set := make([]int, 256)
	// 记录字符分布
	for i := range s {
		set[s[i]]++
	}
	// 计算多出的字符数量
	for i := range set {
		if set[i] <= len(s)/4 {
			set[i] = 0
		} else {
			set[i] -= len(s) / 4
		}
		cnt += set[i]
		set[i] = -set[i]
	}

	// 套覆盖模板
	for l, r := 0, 0; r < len(s); r++ {
		if set[s[r]] < 0 {
			cnt--
		}
		set[s[r]]++
		for ; l <= r && set[s[l]] > 0; l++ {
			set[s[l]]--
		}
		if cnt == 0 && minLenth > r-l+1 {
			minLenth = r - l + 1
		}
	}
	if minLenth == math.MaxInt {
		return 0
	}
	return minLenth
}
func main() {
	balancedString("QQQW")
}
