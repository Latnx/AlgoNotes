package main

func existRight(arr []int, num int) int {
	res := -1
	if len(arr) == 0 {
		return res
	}

	left, right := 0, len(arr)
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] <= num {
			left = mid + 1
			res = mid
		} else if arr[mid] > num {
			right = mid - 1
		}
	}
	return res
}
