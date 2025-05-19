package main

import "fmt"

// 归并分治：
// 1. 一个问题能不能看成左+右+跨
// 2. 若计算跨时，若左右有序，能不能增加便利

// 小和问题
//  https://www.nowcoder.com/practice/edfe05a1d45c4ea89101d936cac32469
// 1 3 5 2 4 6， 左右产生的小和+跨左右的小和，例如3产生的小和一定来自于0-2和左跨右
var Nums []int
var mergeNums []int

func calArray(nums []int) int64 {
	Nums = nums
	mergeNums = make([]int, len(nums))
	return int64(mergeSort(0, len(nums)-1))
}

func mergeSort(l, r int) int {
	if l == r {
		return 0
	}

	m := (l + r) / 2
	lmin := mergeSort(l, m)
	rmin := mergeSort(m+1, r)
	kuamin := merge(l, m, r)
	return lmin + rmin + kuamin
}

func merge(l, m, r int) int {
	var i, j = l, m + 1
	var index = i
	var sum = 0
	var forSum = 0
	for i <= m && j <= r {
		if Nums[i] <= Nums[j] {
			mergeNums[index] = Nums[i]
			forSum += Nums[i]
			i++
		} else if Nums[i] > Nums[j] {
			mergeNums[index] = Nums[j]
			sum += forSum
			j++
		}
		index++
	}
	for i <= m {
		mergeNums[index] = Nums[i]
		i++
		index++
	}
	for j <= r {
		mergeNums[index] = Nums[j]
		j++
		index++
		sum += forSum
	}
	for ; l <= r; l++ {
		Nums[l] = mergeNums[l]
	}
	return sum
}

func main() {
	fmt.Print(calArray(Nums))
}
