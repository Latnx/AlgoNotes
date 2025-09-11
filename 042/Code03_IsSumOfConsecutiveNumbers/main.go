package main

import "fmt"

// 判断一个数字是否是若干数量(数量>1)的连续正整数的和
// 狠狠地暴力，数学规律？不关心
func is1(num int) bool {

	for begin := 1; begin < num; begin++ {
		sum := 0
		for i := begin; sum < num; i++ {
			sum += i
			if sum == num {
				return true
			}
		}
	}
	return false
}

// 判断是2的某次方
func is2(num int) bool {
	return (num & (num - 1)) != 0
}
func main() {
	for i := range 1000 {
		if !is1(i) {
			fmt.Println(i)
		}
	}
}
