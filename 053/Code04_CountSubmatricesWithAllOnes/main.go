package main

// 统计全1子矩形的数量
// 给你一个 m * n 的矩阵 mat，其中只有0和1两种值
// 请你返回有多少个 子矩形 的元素全部都是1
// 测试链接 : https://leetcode.cn/problems/count-submatrices-with-all-ones/

// 联想，使用高度
// 找到左右两边小于的，计算 (height[i] - max(leftHeight, rightHeight)) * (weight * (weight + 1)) / 2
func cal(height []int) int {
	stack := make([]int, len(height))
	r := 0
	leftMin := make([]int, len(height))
	rightMin := make([]int, len(height))
	for i, val := range height {
		for r > 0 && val <= height[stack[r-1]] {
			r--
			if r == 0 {
				leftMin[stack[r]] = -1
			} else {
				leftMin[stack[r]] = stack[r-1]
			}
			rightMin[stack[r]] = i
		}
		stack[r] = i
		r++
	}
	for r > 0 {
		r--
		if r == 0 {
			leftMin[stack[r]] = -1
		} else {
			leftMin[stack[r]] = stack[r-1]
		}
		rightMin[stack[r]] = len(height)
	}
	res := 0
	for i := range height {
		left, right := leftMin[i], rightMin[i]
		leftHeight, rightHeight := 0, 0
		weight := (right - left - 1)
		if left != -1 {
			leftHeight = height[left]
		}
		if right != len(height) {
			rightHeight = height[right]
		}
		res += (height[i] - max(leftHeight, rightHeight)) * (weight * (weight + 1)) / 2
	}
	return res
}

func numSubmat(mat [][]int) int {
	height := make([]int, len(mat[0]))
	res := 0
	for i := range mat {
		for j := range mat[0] {
			if mat[i][j] == 0 {
				height[j] = 0
			} else {
				height[j]++
			}
		}
		res += cal(height)
	}
	return res
}

func main() {
	numSubmat([][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}})
}
