package main

// 最长递增子序列和最长不下降子序列
// 给定一个整数数组nums
// 找到其中最长严格递增子序列长度、最长不下降子序列长度
// 测试链接 : https://leetcode.cn/problems/longest-increasing-subsequence/

// O(N2)
func lengthOfLIS_N2(nums []int) int {
	dp := make([]int, len(nums))
	ans := 0
	for i := range nums {
		dp[i] = 1
		for j := range i {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(dp[i], ans)
	}
	return ans
}

// O(nlogn)
func lengthOfLIS_ENDs(nums []int) int {
	// dp[i]:以i结尾的最长递增子序列长度
	dp := make([]int, len(nums))
	// ends[i]:长度为i+1的递增子序列的最小结尾, 单调增，可二分
	ends := make([]int, len(nums))
	ends[0] = nums[0]
	dp[0] = 1
	r := 1
	ans := 0
	for i := range nums {
		for j := range r {
			if nums[i] <= ends[j] {
				ends[j] = nums[i]
				dp[i] = dp[j]
				break
			}
			if j == r-1 {
				ends[r] = nums[i]
				dp[i] = r + 1
				r++
			}
		}
		ans = max(dp[i], ans)
	}
	return ans
}
func lengthOfLIS(nums []int) int {
	// dp[i]:以i结尾的最长递增子序列长度
	dp := make([]int, len(nums))
	// ends[i]:长度为i+1的递增子序列的最小结尾, 单调增，可二分
	ends := make([]int, len(nums))
	ends[0] = nums[0]
	dp[0] = 1
	r := 1
	ans := 0
	for i := range nums {
		L, R := 0, r
		for L < R {
			mid := L + (R-L)/2
			if nums[i] <= ends[mid] { // mid 可行
				R = mid // 保留 mid
			} else {
				L = mid + 1
			}
		}
		ends[L] = nums[i]
		if L == r {
			dp[i] = r + 1
			r++
		}
		ans = max(dp[i], ans)
	}
	return ans
}
func lengthOfLIS_new(nums []int) int {
	// ends[i]:长度为i+1的递增子序列的最小结尾, 单调增，可二分
	ends := make([]int, len(nums))
	r := 0
	for i := range nums {
		L, R := 0, r
		for L < R {
			mid := L + (R-L)/2
			if ends[mid] >= nums[i] { // mid 可行
				R = mid // 保留 mid
			} else {
				L = mid + 1
			}
		}
		ends[L] = nums[i]
		if L == r {
			r++
		}
	}
	return r
}
func main() {
	lengthOfLIS([]int{7, 7, 7})
}
