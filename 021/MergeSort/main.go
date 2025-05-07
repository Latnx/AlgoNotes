package main

func sortArray(nums []int) []int {
	mergeNums := make([]int, len(nums))
	mergeSort(nums, mergeNums, 0, len(mergeNums)-1)
	return nums
}

func mergeSort(nums []int, mergeNums []int, l int, r int) {
	if l == r {
		return
	}
	middle := (l + r) / 2
	mergeSort(nums, mergeNums, l, middle)
	mergeSort(nums, mergeNums, middle+1, r)
	merge(nums, mergeNums, l, middle, r)
}

func merge(nums []int, mergeNums []int, l int, middle int, r int) {
	i, j := l, middle+1
	index := l
	for i <= middle && j <= r {
		if nums[i] < nums[j] {
			mergeNums[index] = nums[i]
			i++
		} else if nums[i] >= nums[j] {
			mergeNums[index] = nums[j]
			j++
		}
		index++
	}
	for i <= middle {
		mergeNums[index] = nums[i]
		i++
		index++
	}
	for j <= r {
		mergeNums[index] = nums[j]
		j++
		index++
	}
	for ; l <= r; l++ {
		nums[l] = mergeNums[l]
	}
}
