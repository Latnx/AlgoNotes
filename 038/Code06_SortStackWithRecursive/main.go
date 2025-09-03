package main

// 使用递归排序一个栈
// 排序过程中只能用：
// 1) 栈提供的push、pop、isEmpty三个方法(append, [:], len()==0)
// 2) 递归函数，并且返回值最多为单个整数
import (
	"fmt"
)

func sort(stack *[]int) {
	deep := deep(stack)
	for deep > 0 {
		m := maxF(stack, deep)
		k := times(stack, deep, m)
		down(stack, deep, m, k)
		deep -= k
	}
}

// 栈深度
func deep(stack *[]int) int {
	if len(*stack) == 0 {
		return 0
	}
	num := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	deep := deep(stack) + 1
	*stack = append(*stack, num)
	return deep
}

// 从栈当前的顶部开始，往下数deep层
// 返回这deep层里的最大值
func maxF(stack *[]int, deep int) int {
	if deep == 0 {
		return 0
	}
	num := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	maxNums := max(num, maxF(stack, deep-1))
	*stack = append(*stack, num)
	return maxNums
}

// 从栈当前的顶部开始，往下数deep层，已知最大值是max了
// 返回，max出现了几次，不改变栈的数据状况
func times(stack *[]int, deep, max int) int {
	if deep == 0 {
		return 0
	}
	num := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	restTimes := times(stack, deep-1, max)
	if num == max {
		restTimes++
	}
	*stack = append(*stack, num)
	return restTimes
}

// 从栈当前的顶部开始，往下数deep层，已知最大值是max，出现了k次
// 请把这k个最大值沉底，剩下的数据状况不变
func down(stack *[]int, deep, max, k int) {
	if deep == 0 {
		for range k {
			*stack = append(*stack, max)
		}
		return
	}
	num := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	down(stack, deep-1, max, k)
	if num != max {
		*stack = append(*stack, num)
	}
}

func main() {
	test := []int{}
	test = append(test, 1)
	test = append(test, 5)
	test = append(test, 4)
	test = append(test, 5)
	test = append(test, 3)
	test = append(test, 2)
	test = append(test, 3)
	test = append(test, 1)
	test = append(test, 4)
	test = append(test, 2)
	sort(&test)
	for _, val := range test {
		fmt.Println(val)
	}
}
