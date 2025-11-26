package main

// 解码方法
// 测试链接 : https://leetcode.cn/problems/decode-ways/

// 基础dfs
func f1(s []byte, index int) int {
	if index >= len(s) {
		return 1
	}
	ans := 0
	if s[index] == '0' {
		return 0
	}
	if index+1 < len(s) && (s[index] == '1' || (s[index] == '2' && s[index+1] >= '0' && s[index+1] <= '6')) {
		ans += f1(s, index+2)
	}
	ans += f1(s, index+1)
	return ans
}

// 增加缓存
var arr = make([]int, 100)

func f2(s []byte, index int) int {
	if index >= len(s) {
		return 1
	}
	if arr[index] != -1 {
		return arr[index]
	}
	ans := 0
	if s[index] == '0' {
		return 0
	}
	if index+1 < len(s) && (s[index] == '1' || (s[index] == '2' && s[index+1] >= '0' && s[index+1] <= '6')) {
		ans += f2(s, index+2)
	}
	ans += f2(s, index+1)
	arr[index] = ans
	return ans
}

func numDecodings(s string) int {
	for i := range arr {
		arr[i] = -1
	}
	return f1([]byte(s), 0)
}
