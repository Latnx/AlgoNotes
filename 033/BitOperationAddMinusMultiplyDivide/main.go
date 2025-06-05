package main

import (
	"fmt"
)

// https://leetcode.cn/problems/divide-two-integers/

func add(a, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}

func sub(a, b int) int {
	return add(a, add(^b, 1))
}

func multi(a, b int) int {
	res := 0
	for b != 0 {
		if (b & 1) == 1 {
			res += a
		}
		a <<= 1
		b >>= 1
	}
	return res
}

func neg(a int) int {
	return ^a + 1
}

func divide(dividend int, divisor int) int {
	res := 0
	sign := dividend ^ divisor
	if dividend < 0 {
		dividend = neg(dividend)
	}
	if divisor < 0 {
		divisor = neg(divisor)
	}
	for i := 62; i >= 0; i-- {
		if dividend>>i >= divisor {
			dividend -= divisor << i
			res += 1 << i
		}
	}
	if sign < 0 {
		res = neg(res)
	}
	return res
}

func main() {

	fmt.Print(add(15, -27), sub(15, 27), multi(9, 8), divide(47328, -1))
}
