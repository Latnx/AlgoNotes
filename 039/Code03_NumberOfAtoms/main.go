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

// 对之前收集到的名字或者表进行结算
func clear(curEle *[]byte, cur *int, m map[string]int, preMap *map[string]int) {
	if *cur == 0 {
		*cur = 1
	}
	// 来自哪个
	if len(*curEle) > 0 {
		eleName := string(*curEle)
		m[eleName] += *cur
	} else if len(*preMap) > 0 {
		for k, v := range *preMap {
			m[k] += v * (*cur)
		}
	}
	*cur = 0
	*preMap = make(map[string]int)
	*curEle = []byte{}
}

func atoms(formula string) map[string]int {
	// 总表
	m := map[string]int{}
	cur := 0
	// 之前收集到的名字，历史的一部分
	curEle := []byte{}
	// 之前收集到的表，历史的一部分
	preMap := map[string]int{}
	for ; index < len(formula); index++ {
		val := formula[index]
		if val >= 'A' && val <= 'Z' {
			clear(&curEle, &cur, m, &preMap)
			curEle = []byte{val}
		} else if val >= 'a' && val <= 'z' {
			curEle = append(curEle, val)
		} else if val >= '0' && val <= '9' {
			cur = cur*10 + int(val) - '0'
		} else if val == '(' {
			clear(&curEle, &cur, m, &preMap)
			index++
			preMap = atoms(formula)
		} else if val == ')' {
			break
		}
	}
	clear(&curEle, &cur, m, &preMap)
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
	fmt.Print(countOfAtoms("Mg(OH)2"))
}
