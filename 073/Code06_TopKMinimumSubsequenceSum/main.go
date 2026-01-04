package main

// 非负数组前k个最小的子序列累加和
// 给定一个数组nums，含有n个数字，都是非负数
// 给定一个正数k，返回所有子序列中累加和最小的前k个累加和
// 子序列是包含空集的
// 1 <= n <= 10^5
// 1 <= nums[i] <= 10^6
// 1 <= k <= 10^5
// 注意这个数据量，用01背包的解法是不行的，时间复杂度太高了
// 对数器验证
import "sort"

func topKSum1(nums []int, k int) []int {
	allSubsequences := make([]int, 0)
	f1(nums, 0, 0, &allSubsequences)
	sort.Ints(allSubsequences)
	return allSubsequences
}
func f1(nums []int, index int, sum int, res *[]int) {
	if index == len(nums) {
		*res = append(*res, sum)
		return
	}
	f1(nums, index+1, sum, res)
	f1(nums, index+1, sum+nums[index], res)
}

func main() {

}
