package main

// 扰乱字符串
// 使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
// 步骤1 : 如果字符串的长度为 1 ，算法停止
// 步骤2 : 如果字符串的长度 > 1 ，执行下述步骤：
//
//	在一个随机下标处将字符串分割成两个非空的子字符串
//	已知字符串s，则可以将其分成两个子字符串x和y且满足s=x+y
//	可以决定是要 交换两个子字符串 还是要 保持这两个子字符串的顺序不变
//	即s可能是 s = x + y 或者 s = y + x
//	在x和y这两个子字符串上继续从步骤1开始递归执行此算法
//
// 给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串
// 如果是，返回true ；否则，返回false
// 测试链接 : https://leetcode.cn/problems/scramble-string/

var set map[[2]string]bool

func isScramble(s1 string, s2 string) bool {
	set = make(map[[2]string]bool)
	return dfs(s1, s2)
}
func dfs(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) == 0 {
		return false
	}
	if res, ok := set[[2]string{s1, s2}]; ok {
		return res
	}
	if !sameFreq(s1, s2) {
		set[[2]string{s1, s2}] = false
		return false
	}
	for i := 1; i < len(s1); i++ {
		if dfs(s1[:i], s2[:i]) && dfs(s1[i:], s2[i:]) {
			return true
		}
	}
	for i := 1; i < len(s1); i++ {
		if dfs(s1[i:], s2[:len(s1)-i]) && dfs(s1[:i], s2[len(s1)-i:]) {
			return true
		}
	}
	set[[2]string{s1, s2}] = false
	return false
}
func sameFreq(a, b string) bool {
	cnt := [26]int{}
	for i := range a {
		cnt[a[i]-'a']++
		cnt[b[i]-'a']--
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
