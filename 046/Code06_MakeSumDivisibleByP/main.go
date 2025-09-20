package main

// 使数组和能被P整除
// 给你一个正整数数组 nums，请你移除 最短 子数组（可以为 空）
// 使得剩余元素的 和 能被 p 整除。 不允许 将整个数组都移除。
// 请你返回你需要移除的最短子数组的长度，如果无法满足题目要求，返回 -1 。
// 子数组 定义为原数组中连续的一组元素。
// 测试链接 : https://leetcode.cn/problems/make-sum-divisible-by-p/

// 哈希表记录余数, 根据同余原理
//  (a - b) % c == (a % c - b % c +c) % c
// 想让 (a - b) % c = 0
// 只需让a % c = (b % c +c) % c

// 哈希表记录 前缀和的余数
// 如果当前前缀余数大于整体余数(6, 4), 则找到6-4=2的余数最晚出现的位置
// 如果当前前缀余数大于整体余数(4, 6), 则找到4-6+p=-2+p的余数最晚出现的位置
func minSubarray(nums []int, p int) int {
	preSumMap := make(map[int]int)
	preSumMap[0] = -1
	sum := 0
	total := 0
	for _, val := range nums {
		total += val
	}
	totalYu := total % p
	if totalYu == 0 {
		return 0
	}
	res := len(nums)
	for i, val := range nums {
		sum += val
		if index, ok := preSumMap[(sum-totalYu+p)%p]; ok {
			res = min(res, i-index)
		}
		preSumMap[sum%p] = i
	}
	if res == len(nums) {
		return -1
	}
	return res
}

func main() {
	minSubarray([]int{1, 2, 3}, 7)
}
