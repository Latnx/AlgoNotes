package main

import (
	"fmt"
	"math"
)

// 有装下8个苹果的袋子、装下6个苹果的袋子，一定要保证买苹果时所有使用的袋子都装满
// 对于无法装满所有袋子的方案不予考虑，给定n个苹果，返回至少要多少个袋子
// 如果不存在每个袋子都装满的方案返回-1

// 打表发现规律？？
func f(rest int) int {
	if rest < 0 {
		return math.MaxInt32
	}
	if rest == 0 {
		return 0
	}

	return min(f(rest-6), f(rest-8)) + 1
}
func bags1(apple int) int {
	bag := f(apple)
	if bag >= math.MaxInt32 {
		return -1
	}
	return bag
}
func main() {
	for i := range 100 {
		fmt.Println(i, bags1(i))
	}
}
