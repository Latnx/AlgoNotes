package main

// 测试链接 : https://leetcode.cn/problems/maximum-running-time-of-n-computers/
// 碎片电池总电量 >= 电脑数量*分钟数 一定可以 反之一定不行（碎片电池的电量小于分钟数）
// 若有大于的电量，就可以同时减去，和电脑的数量

// 最多满足的电脑数量
func cal(batteries []int, minute int) int {
	computers := 0
	sum := 0
	for _, val := range batteries {
		if val >= minute {
			computers++
		} else {
			sum += val
		}
	}
	computers += sum / minute
	return computers
}

func maxRunTime(n int, batteries []int) int64 {
	sum := 0
	for _, val := range batteries {
		sum += val
	}
	l, r := 1, sum/n
	res := 0
	for l <= r {
		mid := (r-l)>>2 + l
		if cal(batteries, mid) < n {
			r = mid - 1
		} else {
			res = mid
			l = mid + 1
		}
	}
	return int64(res)
}
