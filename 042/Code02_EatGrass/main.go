package main

import "fmt"

// 草一共有n的重量，两只牛轮流吃草，A牛先吃，B牛后吃
// 每只牛在自己的回合，吃草的重量必须是4的幂，1、4、16、64....
// 谁在自己的回合正好把草吃完谁赢，根据输入的n，返回谁赢

func f(rest int, cur string) string {
	var enemy string
	if cur == "B" {
		enemy = "A"
	} else {
		enemy = "B"
	}
	if rest == 0 || rest == 2 {
		return enemy
	} else if rest < 5 {
		return cur
	}
	// enemy 遍历所有方法都不能使enemy自己赢，则cur自己赢
	for i := 1; i < rest; i *= 4 {
		win := f(rest-i, enemy)
		if win == cur {
			return cur
		}
	}
	return enemy
}

func f2(rest int, cur string) string {
	var enemy string
	if cur == "B" {
		enemy = "A"
	} else {
		enemy = "B"
	}
	if rest%5 == 0 || rest%5 == 2 {
		return cur
	} else {
		return enemy
	}
}

func main() {
	for i := range 50 {
		fmt.Println(i, f(i, "A"))
	}
}
