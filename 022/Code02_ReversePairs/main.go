package main

// 翻转对
// https://leetcode.cn/problems/reverse-pairs/

var Nums []int
var sortNums []int

func reversePairs(nums []int) int {
	Nums = nums
	sortNums = make([]int, len(nums))
	return mergeSort(0, len(nums)-1)
}

func mergeSort(l, r int) int {
	if l == r {
		return 0
	}

	m := (l + r) / 2
	b := mergeSort(l, m)
	a := mergeSort(m+1, r)
	c := calPairs(l, m, r)
	merge(l, m, r)
	return a + b + c
}

func calPairs(l, m, r int) int {
	i, j := l, m+1
	sum := 0
	forSum := 0
	for i <= m && j <= r {
		if Nums[i] > 2*Nums[j] {
			forSum++
			j++
		} else {
			sum += forSum
			i++
		}
	}
	for i <= m {
		sum += forSum
		i++
	}
	return sum
}

func merge(l, m, r int) {
	i, j := l, m+1
	index := i
	for ; i <= m && j <= r; index++ {
		if Nums[i] <= Nums[j] {
			sortNums[index] = Nums[i]
			i++
		} else {
			sortNums[index] = Nums[j]
			j++
		}
	}
	for i <= m {
		sortNums[index] = Nums[i]
		i++
		index++
	}
	for j <= r {
		sortNums[index] = Nums[j]
		j++
		index++
	}
	for ; l <= r; l++ {
		Nums[l] = sortNums[l]
	}
}

func main() {
	reversePairs([]int{})
}
