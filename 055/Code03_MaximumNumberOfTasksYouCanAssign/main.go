package main

import (
	"sort"
)

// 测试链接 : https://leetcode.cn/problems/maximum-number-of-tasks-you-can-assign/

// 二分答案法
// cal 函数可以用 单调队列实现
// 没任务可做就吃药
// 没吃药做最小任务, 吃了药做最大任务(能力范围进队列)
// 1. 没吃药进队列
// 2. 不空头部拿
// 3. 为空 吃药, 没有药返回false
// 4. 尾部拿, 拿不到返回false
func cal(task, workers []int, pills, strength int) bool {
	queue := make([]int, len(task))
	l, r := 0, 0
	taskIndex := 0
	add := func(val int) {
		for taskIndex < len(task) && task[taskIndex] <= val {
			queue[r] = taskIndex
			taskIndex++
			r++
		}
	}
	for _, val := range workers {
		add(val)
		if l < r && task[queue[l]] <= val {
			l++
		} else {
			if pills <= 0 {
				return false
			}
			add(val + strength)
			if l < r {
				pills--
				r--
			} else {
				return false
			}
		}
	}
	return true
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	minTask, maxTask := 0, min(len(tasks), len(workers))
	res := 0
	for minTask <= maxTask {
		mid := (minTask + maxTask) / 2
		if cal(tasks[:mid], workers[len(workers)-mid:], pills, strength) { // >= mid
			minTask = mid + 1
			res = mid
		} else {
			maxTask = mid - 1
		}
	}
	return res
}
