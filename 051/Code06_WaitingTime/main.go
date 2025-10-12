package main

// 给你一个数组 time ，其中 time[i] 表示第 i 辆公交车完成 一趟旅途 所需要花费的时间。
// 每辆公交车可以 连续 完成多趟旅途，也就是说，一辆公交车当前旅途完成后，可以 立马开始 下一趟旅途。每辆公交车 独立 运行，也就是说可以同时有多辆公交车在运行且互不影响。
// 给你一个整数 totalTrips ，表示所有公交车 总共 需要完成的旅途数目。请你返回完成 至少 totalTrips 趟旅途需要花费的 最少 时间。
// https://leetcode.cn/problems/minimum-time-to-complete-trips

// 计算给点时间下，能够完成多少趟旅行
func cal(time []int, totalTime int) (totalTrips int) {
	for _, val := range time {
		if val > totalTime {
			continue
		}
		totalTrips += totalTime / val
	}
	return
}

func minimumTime(time []int, totalTrips int) int64 {
	maxTime := 0
	for _, val := range time {
		maxTime = max(maxTime, val)
	}
	l, r := 0, maxTime*totalTrips
	res := 0
	for l <= r {
		mid := (r-l)>>2 + l
		if cal(time, mid) >= totalTrips {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return int64(res)
}
