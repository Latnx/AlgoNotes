package main

import "fmt"

func hanoi(n int) {
	if n > 0 {
		f(n, "左", "右", "中")
	}
}

func f(i int, from, to, other string) {
	if i == 1 {
		fmt.Println("移动圆盘 1 从", from, "到", to)
	} else {
		f(i-1, from, other, to)
		fmt.Println("移动圆盘", i, "从", from, "到", to)
		f(i-1, other, to, from)
	}
}

func main() {
	hanoi(3)
}
