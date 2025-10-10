package main

// 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水
// 测试链接 : https://leetcode.cn/problems/trapping-rain-water/

// 转换为每一列能接多少水，再相加
// 转换为每一列，max（min(max(左)，max(右))-当前格子高度， 0）
// 因此只需要一个结构，求出左侧最大值和右侧最大值
// 只需要两个最大前缀数组即可

// 动态规划
func trap(height []int) int {
	length := len(height)
	left, right := make([]int, length), make([]int, length)
	left[0] = height[0]
	right[length-1] = height[length-1]
	for i := 1; i < length; i++ {
		left[i] = max(left[i-1], height[i])
	}
	for i := length - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	sum := 0
	for i := 1; i < length-1; i++ {
		sum += max(min(left[i-1], right[i+1])-height[i], 0)
	}
	return sum
}

// 双指针
func trap2(height []int) int {
	length := len(height)
	lMax, rMax := height[0], height[length-1]
	l, r := 1, length-2
	sum := 0
	for l <= r {
		// 右侧是真实的，因此可以求出右指针的水量
		if lMax > rMax {
			sum += max(rMax-height[r], 0)
			rMax = max(height[r], rMax)
			r--
		} else {
			sum += max(lMax-height[l], 0)
			lMax = max(height[l], lMax)
			l++
		}
	}
	return sum
}

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
