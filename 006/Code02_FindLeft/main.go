package main

// 大于等于num的最左位置
func existLeft(arr []int, num int) int {
	ans := -1
	if len(arr) == 0 {
		return ans
	}

	var l, r int = 0, len(arr) - 1
	for l <= r {
		m := (l + r) / 2
		if arr[m] >= num {
			r = m - 1
			ans = m
		} else if arr[m] < num {
			l = m + 1
		}
	}

	return ans
}
