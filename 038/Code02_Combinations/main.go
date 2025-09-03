package main

import (
	"sort"
)

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的组合
// 比如输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
// 测试链接 : https://leetcode.cn/problems/subsets-ii/

// 使用如上题方法，出现一些问题
// 即不容易对子集序列进行去重，可能需排序后转字符串去重，复杂度较高

// 本题先对目标进行排序，后选择从不同的元素中抽取0-x个，作为结果
// 按上题经验，先找到多叉树，然后对树进行遍历，因此感觉使用集合效果更佳

var keyArr []int

func dfs(nums map[int]int, index int, result *[][]int, path []int) {
	if len(nums) == index {
		var temp []int = make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := range nums[keyArr[index]] + 1 {
		var temp []int = make([]int, len(path))
		copy(temp, path)
		for range i {
			temp = append(temp, keyArr[index])
		}
		dfs(nums, index+1, result, temp)
	}
}
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	keyArr = nil
	numsSet := map[int]int{}
	for _, value := range nums {
		numsSet[value]++
	}
	for key := range numsSet {
		keyArr = append(keyArr, key)
	}
	dfs(numsSet, 0, &res, []int{})
	return res
}

// 原方法

var pathes [][]int
var Nums []int

func subsetsWithDup2(nums []int) [][]int {
	sort.Ints(nums)
	pathes = make([][]int, 0)
	Nums = nums
	f(0, make([]int, 0))
	return pathes
}

func f(i int, path []int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	if i >= len(Nums) {
		pathes = append(pathes, make([]int, len(path)))
		copy(pathes[len(pathes)-1], path)
		return
	}
	k := i
	// 记录下一个数的第一个位置
	for ; k < len(Nums) && Nums[i] == Nums[k]; k++ {
	}
	// 这组要零个
	f(k, tmp)
	for ; i < k; i++ {
		tmp = append(tmp, Nums[i])
		f(k, tmp)
	}
}
