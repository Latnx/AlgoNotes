package main

import "fmt"

// 一个正整数如果能被 a 或 b 整除，那么它是神奇的。
// 给定三个整数 n , a , b ，返回第 n 个神奇的数字。
// 因为答案可能很大，所以返回答案 对 1000000007 取模
// 测试链接 : https://leetcode.cn/problems/nth-magical-number/

// 二分答案法
// 1. 先定范围: (1, n*min(a,b))
// 2.f(min(a, b))
// 容斥原理
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// 在[1， x],范围中有多少数字能够被a或b整除
func f(x, a, b int) int {
	return x/a + x/b - x/lcm(a, b)
}

func nthMagicalNumber(n int, a int, b int) int {
	res := 0
	for l, r, mid := 0, n*min(a, b), 0; l <= r; {
		mid = (r-l)/2 + l
		cnt := f(mid, a, b)
		if cnt >= n {
			r = mid - 1
			res = mid
		} else if cnt < n {
			l = mid + 1
		}
	}
	return res % 1000000007
}
func main() {
	fmt.Print(nthMagicalNumber(4, 2, 3))
}
