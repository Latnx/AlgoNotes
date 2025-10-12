package main

// 分割数组的最大值(画匠问题)
// 给定一个非负整数数组 nums 和一个整数 m
// 你需要将这个数组分成 m 个非空的连续子数组。
// 设计一个算法使得这 m 个子数组各自和的最大值最小。
// 测试链接 : https://leetcode.cn/problems/split-array-largest-sum/

// 都是先假定一个结果，算给定的另外一个参数，然后二分
// 需要单调性
// 给定最大值，求需要分成几组
func cal(nums []int, res int) int {
	k := 0
	sum := 0
	for _, val := range nums {
		if val > res {
			return len(nums) + 1
		}
		if sum+val > res {
			sum = 0
			sum += val
			k++
		} else if sum+val == res {
			sum = 0
			k++
		} else if sum+val < res {
			sum += val
		}
	}
	if sum > 0 {
		k++
	}
	return k
}
func splitArray(nums []int, k int) int {
	l, r := 0, 0
	for _, val := range nums {
		r += val
	}
	res := 0
	for l <= r {
		mid := (l + r) / 2
		if cal(nums, mid) <= k {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func main() {
	splitArray([]int{1, 4, 4}, 3)
}
