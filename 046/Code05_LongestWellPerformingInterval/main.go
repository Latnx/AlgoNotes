package main

// 表现良好的最长时间段
// 给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数
// 我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是 劳累的一天
// 所谓 表现良好的时间段 ，意味在这段时间内，「劳累的天数」是严格 大于 不劳累的天数
// 请你返回 表现良好时间段 的最大长度
// 测试链接 : https://leetcode.cn/problems/longest-well-performing-interval/

// 劳累天数是-1 不劳累是1, 计算区间和小于零的最长段
// 这样的数组,变化是连续的,因为都是一点一点计算出来的,所以 若存在-4 ,-4之前肯定存在-3,因此只需要找到最接近零的边界即可

func longestWPI(hours []int) int {
	maxLength := 0
	preSum := make([]int, len(hours)+1)
	for k := range hours {
		v := 0
		if hours[k] > 8 {
			v = -1
		} else {
			v = 1
		}
		preSum[k+1] = preSum[k] + v
	}
	preSumMap := make(map[int]int)
	// 哈希表记录前缀
	for i := len(preSum) - 1; i >= 0; i-- {
		preSumMap[preSum[i]] = i
	}
	for i := range preSum {
		if preSum[i] < 0 {
			maxLength = max(maxLength, i)
		} else {
			index, ok := preSumMap[preSum[i]+1]
			if ok {
				maxLength = max(maxLength, i-index)
			}
		}
	}
	return maxLength
}
