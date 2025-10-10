package main

// 寻找重复数
// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n）
// 可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
// 测试链接 : https://leetcode.cn/problems/find-the-duplicate-number/

// 其实是快慢指针判断链表成环
// 静态链表
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// 仿 Code07
func findDuplicate2(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == l+1 {
			l++
		} else if nums[nums[l]-1] == nums[l] {
			return nums[l]
		} else {
			nums[l], nums[nums[l]-1] = nums[nums[l]-1], nums[l]
		}
	}
	return 0
}
func main() {
	findDuplicate([]int{1, 3, 4, 2, 2})
}
