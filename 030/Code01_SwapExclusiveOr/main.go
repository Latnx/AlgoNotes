package main

func swap(a, b int) {
	a ^= b
	b ^= a
	a ^= b
}
