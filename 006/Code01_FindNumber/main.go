package main

import "fmt"

// 二分法
// O(logn)
func exist(arr []int, num int) bool {
	if len(arr) == 0 {
		return false
	}
	var l, r int = 0, len(arr) - 1
	for l <= r {
		m := (l + r) / 2
		if arr[m] == num {
			return true
		} else if arr[m] > num {
			r = m - 1
		} else if arr[m] < num {
			l = m + 1
		}
	}

	return false
}
func main() {
	fmt.Println(exist([]int{1, 2, 3, 4, 5}, 3)) // true
}
