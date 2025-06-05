package main

// 返回大于等于 n 最小的 2 的幂
func near2Power(n int) int {
	if n <= 0 {
		return 0
	}
	// 防止 n自己是2的幂
	n -= 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n + 1
}
