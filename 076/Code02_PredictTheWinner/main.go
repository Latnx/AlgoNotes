package main

// 预测赢家
// 给你一个整数数组 nums 。玩家 1 和玩家 2 基于这个数组设计了一个游戏
// 玩家 1 和玩家 2 轮流进行自己的回合，玩家 1 先手
// 开始时，两个玩家的初始分值都是 0
// 每一回合，玩家从数组的任意一端取一个数字
// 取到的数字将会从数组中移除，数组长度减1
// 玩家选中的数字将会加到他的得分上
// 当数组中没有剩余数字可取时游戏结束
// 如果玩家 1 能成为赢家，返回 true
// 如果两个玩家得分相等，同样认为玩家 1 是游戏的赢家，也返回 true
// 你可以假设每个玩家的玩法都会使他的分数最大化
// 测试链接 : https://leetcode.cn/problems/predict-the-winner/
func predictTheWinner(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	player1 := f2(nums)
	return sum-player1 <= player1
}
func f1(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	if l+1 == r {
		return max(nums[l], nums[r])
	}
	p1 := nums[l] + min(f1(nums, l+2, r), f1(nums, l+1, r-1))
	p2 := nums[r] + min(f1(nums, l+1, r-1), f1(nums, l, r-2))
	return max(p1, p2)
}
func f2(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for l := range n {
		dp[l][l] = nums[l]
	}
	for l := 0; l+1 < n; l++ {
		dp[l][l+1] = max(nums[l], nums[l+1])
	}
	for l := n - 3; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			p1 := nums[l] + min(dp[l+2][r], dp[l+1][r-1])
			p2 := nums[r] + min(dp[l+1][r-1], dp[l][r-2])
			dp[l][r] = max(p1, p2)
		}
	}
	return dp[0][n-1]
}
