package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

// 暴力解
func base(arr []int) []int {
	return []int{}
}

// 最优解
func quickFunc(arr []int) []int {
	return []int{}
}

func generateRandomArray() []int {
	arr := make([]int, rand.Intn(100))
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func main() {
	for range 1000 {
		arr := generateRandomArray()
		ans1 := base(arr)
		ans2 := quickFunc(arr)
		if !reflect.DeepEqual(ans1, ans2) {
			fmt.Println("arr:", arr, "ans1:", ans1, "ans2:", ans2)
			return
		}
	}
}
