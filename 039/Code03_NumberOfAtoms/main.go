package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 含有嵌套的分子式求原子数量
// 测试链接 : https://leetcode.cn/problems/number-of-atoms/
/*
思路：
	有五类字符
	数字跟元素或者元素的表
套模板
 1. 先设计不存在()的解法
    a.遇到大写，清算时刻！！
	b.遇到小写，手机名字
	c.遇到数字，cur = cur*10 + int(val) - '0'


 2. 使用递归解决括号问题
    a.遇到（，清算时刻！！
    b.遇到）就返回表
    -.可不在break后再对i++,因为递归结束后会执行i++
*/

var index int

func parseName(formula string) string {
	curEle := []byte{formula[index]}
	index++
	for index < len(formula) && formula[index] >= 'a' && formula[index] <= 'z' {
		curEle = append(curEle, formula[index])
		index++
	}
	return string(curEle)
}

func parseNum(formula string) int {
	n := len(formula)
	if index >= n || formula[index] < '0' || formula[index] > '9' {
		return 1
	}
	num := 0
	for index < n && formula[index] >= '0' && formula[index] <= '9' {
		num = num*10 + int(formula[index]-'0')
		index++
	}
	return num
}

func atoms(formula string) map[string]int {
	// 总表
	m := map[string]int{}
	for index < len(formula) {
		val := formula[index]
		if val >= 'A' && val <= 'Z' {
			m[parseName(formula)] += parseNum(formula)
		} else if val == '(' {
			index++
			pre := atoms(formula)
			num := parseNum(formula)
			for k, v := range pre {
				m[k] += v * num
			}
		} else if val == ')' {
			index++
			break
		}
	}
	return m
}

func countOfAtoms(formula string) string {
	index = 0
	m := atoms(formula)
	res := []byte{}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		res = append(res, key...)
		if m[key] != 1 {
			res = append(res, strconv.Itoa(m[key])...)
		}
	}
	return string(res)
}

func main() {
	fmt.Print(countOfAtoms("Mg2(O2(OH)3H)2"))
}
