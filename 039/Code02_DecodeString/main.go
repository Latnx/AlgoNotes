package main

// 含有嵌套的字符串解码
// 测试链接 : https://leetcode.cn/problems/decode-string/
/*
1. 先设计不存在括号的解法
	a.遍历每一个字符
    c.遇到字母就添加
    d.遇到数字就计算，cur = cur*10 + int(s[index]) - int('0')

 2. 使用递归解决括号问题
    a.遇到[就将下一坐标递归
    b.遇到]就返回当前字符串
    -.可不在break后再对i++,因为递归结束后会执行i++
*/
var index int

func decode(s string) string {
	res := []byte{}
	cur := 0
	for ; index < len(s); index++ {
		if s[index] >= 'a' && s[index] <= 'z' {
			res = append(res, s[index])
		} else if s[index] >= '0' && s[index] <= '9' {
			cur = cur*10 + int(s[index]) - int('0')
		} else if s[index] == '[' {
			index++
			r := decode(s)
			for range cur {
				res = append(res, r...)
			}
			cur = 0
		} else if s[index] == ']' {
			break
		}
	}
	return string(res)
}

func decodeString(s string) string {
	index = 0
	return decode(s)
}
