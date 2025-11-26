package main

// 斐波那契数
// 测试链接 : https://leetcode.cn/problems/fibonacci-number/
// 1.递归O(2^n)
// 2.+缓存表，从顶到底记录，记忆化搜索O(N)
// 3.从简单到复杂，从底到顶记录,O(N)
// 4.优化空间
var arr []int

func f1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	if arr[n] != -1 {
		return arr[n]
	}
	arr[n] = f1(n-1) + f1(n-2)
	return arr[n]
}
func fib(n int) int {
	arr = make([]int, n+1)
	for i := range arr {
		arr[i] = -1
	}
	return f1(n)
}

func main() {
	fib(30)
}
