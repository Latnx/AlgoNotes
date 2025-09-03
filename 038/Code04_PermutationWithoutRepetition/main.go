package main

// 有重复项数组的去重全排列
// 测试链接 : https://leetcode.cn/problems/permutations-ii/
func dfs(nums []int, index int, res *[][]int) {
	if index == len(nums) {
		*res = append(*res, make([]int, len(nums)))
		copy((*res)[len(*res)-1], nums)
		return
	}
	set := map[int]struct{}{}
	for i := index; i < len(nums); i++ {
		// 保证每个来到index的数字都是不一样的
		_, ok := set[nums[i]]
		if ok {
			continue
		}
		set[nums[i]] = struct{}{}
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1, res)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

// 同样抽象成一个多叉树
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, 0, &res)
	return res
}
