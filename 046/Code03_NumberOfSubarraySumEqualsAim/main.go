package main

// 返回无序数组中累加和为给定值的子数组个数
// 测试链接 : https://leetcode.cn/problems/subarray-sum-equals-k/

// 统计以当前下标为末尾的所有可能（已经覆盖所有情况），因此向前查找
// 对前缀和进行记录，为方便查找使用哈希表存储
// 添加0下标,表示当前前缀和就可以表示,不需要相减[:i]
//
func subarraySum(nums []int, k int) int {
	res := 0
	preMap := make(map[int]int)
	preMap[0] = 1
	sum := 0
	for _, val := range nums {
		sum += val
		res += preMap[sum-k]
		preMap[sum]++
	}
	return res
}
func main() {
	subarraySum([]int{1, 1, 1}, 2)
}
