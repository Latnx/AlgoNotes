package main

// 移除盒子
// 给出一些不同颜色的盒子boxes，盒子的颜色由不同的正数表示
// 你将经过若干轮操作去去掉盒子，直到所有的盒子都去掉为止
// 每一轮你可以移除具有相同颜色的连续 k 个盒子（k >= 1）
// 这样一轮之后你将得到 k * k 个积分
// 返回你能获得的最大积分总和
// 测试链接 : https://leetcode.cn/problems/remove-boxes/
func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp = make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	return f1(boxes, 0, n-1, 0)
}

var dp [][][]int

// [l...r]范围最大积分，k为前面跟了k个与boxes[l]相同颜色的盒子
func f1(boxes []int, l, r int, k int) int {
	// 终止条件
	if l > r {
		return 0
	}
	if dp[l][r][k] > 0 {
		return dp[l][r][k]
	}
	// 递归
	s := l
	// 找出连续的颜色
	for ; s+1 <= r && boxes[s+1] == boxes[l]; s++ {
	}
	cnt := k + s - l + 1
	// 递归，1.只算这段2.和下一段相同颜色一块移除
	res := f1(boxes, s+1, r, 0) + cnt*cnt
	// 找出下一个相同的颜色
	for next := s + 2; next <= r; next++ {
		if boxes[next] == boxes[l] && boxes[next-1] != boxes[next] {
			res = max(res, f1(boxes, s+1, next-1, 0)+f1(boxes, next, r, cnt))
		}
	}
	dp[l][r][k] = res
	return res
}
