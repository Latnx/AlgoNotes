package main

import (
	"fmt"
	"math"
)

// 现在有一个打怪类型的游戏，这个游戏是这样的，你有n个技能
// 每一个技能会有一个伤害，
// 同时若怪物小于等于一定的血量，则该技能可能造成双倍伤害
// 每一个技能最多只能释放一次，已知怪物有m点血量
// 现在想问你最少用几个技能能消灭掉他(血量小于等于0)
// 技能的数量是n，怪物的血量是m
// i号技能的伤害是x[i]，i号技能触发双倍伤害的血量最小值是y[i]
// 1 <= n <= 10
// 1 <= m、x[i]、y[i] <= 10^6
// 测试链接 : https://www.nowcoder.com/practice/d88ef50f8dab4850be8cd4b95514bbbd
var kill []int
var blood []int

func f(index int, r int) int {
	if r <= 0 {
		return index
	}
	if index == len(kill) {
		return math.MaxInt32
	}
	res := math.MaxInt32
	for i := index; i < len(kill); i++ {
		kill[i], kill[index] = kill[index], kill[i]
		blood[i], blood[index] = blood[index], blood[i]
		a, b := math.MaxInt32, math.MaxInt32
		if r <= blood[index] {
			a = f(index+1, r-kill[index]*2)
		} else {
			b = f(index+1, r-kill[index])
		}
		res = min(res, a, b)
		kill[i], kill[index] = kill[index], kill[i]
		blood[i], blood[index] = blood[index], blood[i]
	}
	return res
}
func main() {
	n := 0
	fmt.Scan(&n)

	for range n {
		skillN := 0
		rest := 0
		kill = make([]int, n)
		blood = make([]int, n)
		fmt.Scan(&skillN, &rest)
		for i := range skillN {
			fmt.Scan(&kill[i], &blood[i])
		}
		fmt.Println("答案：", f(0, rest))
	}
	/*
		10 20
		45 89
		5 40
	*/
	// kill = []int{10, 45, 5}
	// blood = []int{20, 89, 40}
	// fmt.Println(f(0, 100))
}
