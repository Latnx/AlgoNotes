package main

// 没有重复项数字的全排列
// 测试链接 : https://leetcode.cn/problems/permutations/

func dfs(nums []int, index int, res *[][]int) {
	if index == len(nums) {
		*res = append(*res, make([]int, len(nums)))
		copy((*res)[len(*res)-1], nums)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1, res)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

// 同样抽象成一个多叉树
func permute(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, 0, &res)
	return res
}
