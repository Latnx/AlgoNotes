package main

import (
	"fmt"
	"math"
)

// 返回无序数组中正数和负数个数相等的最长子数组长度
// 给定一个无序数组arr，其中元素可正、可负、可0
// 求arr所有子数组中正数与负数个数相等的最长子数组的长度
// 测试链接 : https://www.nowcoder.com/practice/545544c060804eceaed0bb84fcd992fb
func main() {
	N := 0
	fmt.Scan(&N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
		if arr[i] < 0 {
			arr[i] = -1
		} else if arr[i] > 0 {
			arr[i] = 1
		}
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
		j, ok := preSumMap[val-0]
		if ok {
			maxLength = int(math.Max(float64(maxLength), float64(i-j)))
		}
	}
	fmt.Print(maxLength)
}
