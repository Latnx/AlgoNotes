package main

// 求最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
