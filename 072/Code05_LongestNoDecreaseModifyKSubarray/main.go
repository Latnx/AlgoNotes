package main

import (
	"bufio"
	"os"
	"strconv"
)

// 有一次修改机会的最长不下降子序列
// 给定一个长度为n的数组arr，和一个整数k
// 只有一次机会可以将其中连续的k个数全修改成任意一个值
// 这次机会你可以用也可以不用，请返回最长不下降子序列长度
// 1 <= k, n <= 10^5
// 1 <= arr[i] <= 10^6
// 测试链接 : https://www.luogu.com.cn/problem/P8776

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	K, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, N)
	for i := range N {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	writer.WriteString(strconv.Itoa(solute(nums, N, K)))
}
func solute(nums []int, n, k int) int {
	ends := make([]int, n)
	r := 0
	ans := 0
	right := reverseRight(nums)
	for i := k; i < n; i++ {
		ans = max(ans, binFind(ends, nums[i], 0, r)+k+right[i])
		// 释放一个数
		pos := binFind(ends, nums[i-k], 0, r)
		if pos == r {
			ends[pos] = nums[i-k]
			r++
		} else {
			ends[pos] = nums[i-k]
		}
	}
	ans = max(ans, r+k)
	return ans
}

func reverseRight(nums []int) []int {
	n := len(nums)
	right := make([]int, n)
	ends := make([]int, n)
	lenEnds := 0 // ends 中有效元素的个数

	// 从右往左遍历，计算最长不上升子序列
	// ends 数组是降序的
	for i := n - 1; i >= 0; i-- {
		// 在降序数组 ends[0:lenEnds] 中找到第一个 < nums[i] 的位置
		pos := binFindLess(ends, nums[i], lenEnds)

		if pos == -1 {
			// 没找到，nums[i] 是当前最小的，追加到末尾
			ends[lenEnds] = nums[i]
			lenEnds++
			right[i] = lenEnds
		} else {
			// 找到了，更新该位置
			ends[pos] = nums[i]
			right[i] = pos + 1
		}
	}

	return right
}
func binFindLess(arr []int, target int, len int) int {
	l, r := 0, len-1
	ans := -1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < target {
			ans = mid
			r = mid - 1 // 继续向左找
		} else {
			l = mid + 1
		}
	}

	return ans
}

// 返回首个大于target的下标
func binFind(arr []int, target int, l, r int) int {
	for l < r {
		mid := (l + r) / 2
		if arr[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
