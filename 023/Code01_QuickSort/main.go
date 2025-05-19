package main

import "math/rand"

func sort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := partition(nums, l, r)
	quickSort(nums, l, mid-1)
	quickSort(nums, mid+1, r)
}

// 随即一个位置，返回划分点，这个划分点位置正确
func partition(nums []int, l, r int) int {
	pivot := nums[l+rand.Intn(r-l)]
	// a 是 小于等于 pivot 区域的越界位置
	i, a := l, l
	mid := 0
	for i <= r {
		if nums[i] <= pivot {
			nums[i], nums[a] = nums[a], nums[i]
			if nums[a] == pivot {
				mid = a
			}
			a++
		}
		i++
	}
	nums[a-1], nums[mid] = nums[mid], nums[a-1]
	return a - 1
}

// 波兰国旗改进快排,两个区域，小于区和大于区
// < 中轴，放在左边，a++, i++
// = 中轴，i++
// > 中轴，b--

func partition_flag(nums []int, l, r int) (int, int) {
	pivot := nums[l+rand.Intn(r-l)]
	i, a, b := l, l, r
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
	return a - 1, b + 1
}
func main() {
	sort([]int{})
}
