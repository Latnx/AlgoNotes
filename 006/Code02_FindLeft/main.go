package main

import "fmt"

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

func nextGreatestLetter(letters []byte, target byte) byte {
	ans := letters[0]
	var l, r = 0, len(letters) - 1
	for l <= r {
		middle := (l + r) / 2
		if letters[middle] >= target {
			r = middle - 1
			ans = letters[middle]
		} else if letters[middle] < target {
			l = middle + 1
		}
	}
	return ans
}
func main() {
	fmt.Print(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
}
