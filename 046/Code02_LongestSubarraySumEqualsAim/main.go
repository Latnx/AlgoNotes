package main

import (
	"fmt"
	"math"
)

// 返回无序数组中累加和为给定值的最长子数组长度
// 给定一个无序数组arr, 其中元素可正、可负、可0
// 给定一个整数aim
// 求arr所有子数组中累加和为aim的最长子数组长度

// 测试链接 : https://www.nowcoder.com/practice/36fb0fd3c656480c92b569258a1223d5
func main() {
	// 第一行两个整数N, k。N表示数组长度，k的定义已在题目描述中给出
	// 第二行N个整数表示数组内的数
	N, k := 0, 0
	fmt.Scan(&N, &k)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}
	maxLength := 0
	preSum := make([]int, N+1)
	preSumMap := make(map[int]int)
	for k, v := range arr {
		preSum[k+1] = preSum[k] + v
	}
	for i := len(preSum) - 1; i >= 0; i-- {
		preSumMap[preSum[i]] = i
	}
	for i, val := range preSum {
		j, ok := preSumMap[val-k]
		if ok {
			maxLength = int(math.Max(float64(maxLength), float64(i-j)))
		}
	}
	fmt.Print(maxLength)
}
