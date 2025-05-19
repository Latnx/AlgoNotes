package main

import "math/rand"

// 随即选择算法
// 数组中的第K个最大元素,O(n)
// https://leetcode.cn/problems/kth-largest-element-in-an-array

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1, k)
	return nums[len(nums)-k]
}

func quickSort(nums []int, l, r, k int) {
	if l >= r {
		return
	}
	a, b := partition(nums, l, r)
	pos := len(nums) - k
	if a <= pos && pos <= b {
		return
	} else if pos > b {
		quickSort(nums, b+1, r, k)
	} else if pos < a {
		quickSort(nums, l, a-1, k)
	}

}
func partition(nums []int, l, r int) (int, int) {
	pivot := nums[l+rand.Intn(r-l)]
	var a, i = l, l
	var b = r
	for i <= b {
		if nums[i] < pivot {
			nums[i], nums[a] = nums[a], nums[i]
			a++
			i++
		} else if nums[i] > pivot {
			nums[i], nums[b] = nums[b], nums[i]
			b--
		} else {
			i++
		}
	}
	return a, b
}

func main() {
	findKthLargest([]int{}, 0)
}
