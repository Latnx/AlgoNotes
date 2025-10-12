package main

import (
	"fmt"
	"math"
)

// 机器人跳跃问题
// 机器人正在玩一个古老的基于DOS的游戏
// 游戏中有N+1座建筑，从0到N编号，从左到右排列
// 编号为0的建筑高度为0个单位，编号为i的建筑的高度为H(i)个单位
// 起初机器人在编号为0的建筑处
// 每一步，它跳到下一个（右边）建筑。假设机器人在第k个建筑，且它现在的能量值是E
// 下一步它将跳到第个k+1建筑
// 它将会得到或者失去正比于与H(k+1)与E之差的能量
// 如果 H(k+1) > E 那么机器人就失去H(k+1)-E的能量值，否则它将得到E-H(k+1)的能量值
// 游戏目标是到达第个N建筑，在这个过程中，能量值不能为负数个单位
// 现在的问题是机器人以多少能量值开始游戏，才可以保证成功完成游戏
// 测试链接 : https://www.nowcoder.com/practice/7037a3d57bbd4336856b8e16a9cafd71
func cal(builds []int, E int) bool {
	for _, build := range builds {
		E = 2*E - build
		if E < 0 {
			return false
		} else if E > 100000 {
			return false
		}
	}
	return true
}

func main() {
	N := 0
	fmt.Scan(&N)
	builds := make([]int, N)
	m := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&builds[i])
		m = int(math.Max(float64(builds[i]), float64(m)))
	}
	res := 0
	l, r := 0, m
	for l <= r {
		mid := (l + r) / 2
		if cal(builds, mid) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Print(res)
}
