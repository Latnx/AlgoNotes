package main

import "fmt"

/*
含有嵌套的表达式求值:请写一个整数计算器，支持加减乘三种运算和括号。
测试链接 : https://leetcode.cn/problems/basic-calculator-iii/
测试链接 : https://www.nowcoder.com/practice/c215ba61c8b1443b996351df929dc4d4

 1. 先设计不存在括号的解法
    a.准备两个栈，数字栈和符号栈
    b.遍历每一个字符
    c.遇到数字就记录, cur = cur*10 + int(val) - int('0')
    d.遇到符号就压栈, cur = 0
    c.若压栈时，栈顶为*或/，使用栈顶符号和数字计算后再压栈
    e.读取最后一个数字后直接压入num,+

 2. 使用递归解决括号问题
    a.遇到（就将下一坐标递归
    b.遇到）就返回当前计算值
    -.可不在break后再对i++,因为递归结束后会执行i++
*/
func calStack(number *[]int, ops *[]byte) int {
	// 最后再计算栈的时候，符号栈内只有+-
	res := (*number)[0]
	for index := 1; index < len(*number); index++ {
		switch (*ops)[index-1] {
		case '+':
			res += (*number)[index]
		case '-':
			res -= (*number)[index]
		}
	}
	return res
}

func push(numbers *[]int, ops *[]byte, cur int, op byte) {
	if len(*ops) == 0 {
		*numbers = append(*numbers, cur)
		*ops = append(*ops, op)
		return
	}
	switch (*ops)[len(*ops)-1] {
	case '+', '-':
		*numbers = append(*numbers, cur)
		*ops = append(*ops, op)
	case '*':
		(*numbers)[len(*numbers)-1] *= cur
		(*ops)[len(*ops)-1] = op
	case '/':
		(*numbers)[len(*numbers)-1] /= cur
		(*ops)[len(*ops)-1] = op
	}
}

var i = 0

func calculator(s string) int {
	numbers := []int{}
	ops := []byte{}
	cur := 0 //记录一个数字
	for val := byte(0); i < len(s); i++ {
		val = s[i]
		if val >= '0' && val <= '9' {
			// 遇到数字就统计
			cur = cur*10 + int(val) - int('0')
		} else if val == '+' || val == '-' || val == '*' || val == '/' {
			// 遇到符号就入栈
			push(&numbers, &ops, cur, val)
			cur = 0
		} else if val == '(' {
			// 左括号就入栈
			i++
			cur = calculator(s)
		} else if val == ')' {
			break
		}
	}
	push(&numbers, &ops, cur, '+')
	return calStack(&numbers, &ops)
}

func solve(s string) int {
	i = 0
	return calculator(s)
}

func main() {
	fmt.Print(solve("(3+4)*(5+(2-3))"))
}
