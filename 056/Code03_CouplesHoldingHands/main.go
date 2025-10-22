package main

// 情侣牵手
// n对情侣坐在连续排列的 2n 个座位上，想要牵到对方的手
// 人和座位由一个整数数组 row 表示，其中 row[i] 是坐在第 i 个座位上的人的ID
// 情侣们按顺序编号，第一对是 (0, 1)，第二对是 (2, 3)，以此类推，最后一对是 (2n-2, 2n-1)
// 返回 最少交换座位的次数，以便每对情侣可以并肩坐在一起
// 每次交换可选择任意两人，让他们站起来交换座位
// 测试链接 : https://leetcode.cn/problems/couples-holding-hands/

// 如果一个集合里有k对情侣,要交换k-1此
var arr []int
var size []int

func find(i int) int {
	if arr[i] != i {
		arr[i] = find(arr[i])
	}
	return arr[i]
}
func union(x, y int) {
	findX, findY := find(x), find(y)
	if findX == findY {
		return
	}
	arr[findX] = findY
	size[findY] += size[findX]
	size[findX] = 1
}
func minSwapsCouples(row []int) int {
	arr = make([]int, len(row)/2)
	size = make([]int, len(row)/2)
	for i := range arr {
		arr[i] = i
		size[i] = 1
	}
	for i := 0; i < len(row)/2; i++ {
		union(row[2*i]/2, row[2*i+1]/2)
	}
	res := 0
	for _, val := range size {
		res += val - 1
	}
	return res
}

func main() {
	minSwapsCouples([]int{0, 2, 1, 3})
}
