package main

import "fmt"

// 计数排序
// 基数排序 非比较的排序
// 技巧一：前缀数量分区（将所有小于等于0的位置找到，从后向前遍历，可知位置）将一个一个桶，压成一个
// 技巧二：提取数字
func sort(nums []int) []int {
	return radixSort(nums, 10)
}
func radixSort(nums []int, bits int) []int {
	help := make([]int, len(nums))
	for offset := 1; bits > 0; bits-- {
		// 计数
		cnt := [10]int{}
		for _, v := range nums {
			cnt[(v/offset)%10]++
		}
		// 分区
		for i := 1; i < len(cnt); i++ {
			cnt[i] += cnt[i-1]
		}
		// 前缀统计
		for i := len(nums) - 1; i >= 0; i-- {
			cnt[(nums[i]/offset)%10]--
			help[cnt[(nums[i]/offset)%10]] = nums[i]
		}
		copy(nums, help)
		offset *= 10
	}
	return nums
}
func main() {

	fmt.Print(sort([]int{1, 312, 3123, 41, 4, 45, 6347, 645443, 243, 565437, 545, 6654, 643, 6}))
}
