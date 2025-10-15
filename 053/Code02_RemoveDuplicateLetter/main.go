package main

// 去除重复字母保证剩余字符串的字典序最小
// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次
// 需保证 返回结果的字典序最小
// 要求不能打乱其他字符的相对位置
// 测试链接 : https://leetcode.cn/problems/remove-duplicate-letters/

// 栈中是保证词频的最小字典序，栈中大压小(尽量), (满足条件为出栈的字符后面仍然存在)
func removeDuplicateLetters(s string) string {
	m := make(map[byte]int)
	stack := make([]int, len(s))
	exict := make(map[byte]int)
	r := 0
	for _, val := range s {
		m[byte(val)]++
	}
	for i, val := range s {
		m[s[i]]--
		for r > 0 && s[stack[r-1]] >= byte(val) && m[s[stack[r-1]]] > 0 && exict[s[i]] == 0 {
			r--
			exict[s[stack[r]]]--
		}
		if exict[s[i]] == 0 {
			exict[s[i]]++
			stack[r] = i
			r++
		}
	}
	var bytes []byte
	for i := 0; i < len(m); i++ {
		bytes = append(bytes, s[stack[i]])
	}
	res := string(bytes)
	return res
}
func main() {
	removeDuplicateLetters("cdadabcc")
}
