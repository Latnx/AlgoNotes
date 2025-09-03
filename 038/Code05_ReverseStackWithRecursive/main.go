package main

import "fmt"

// 递归逆序一个栈
func reverse(stack *[]int) {
	if len(*stack) == 0 {
		return
	}
	ans := bottomOut(stack)
	reverse(stack)
	*stack = append(*stack, ans)
}

// 移除并返回栈底元素
func bottomOut(stack *[]int) int {
	ans := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	if len(*stack) == 0 {
		return ans
	}
	last := bottomOut(stack)
	*stack = append(*stack, ans)
	return last
}
func main() {
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	stack = append(stack, 5)
	reverse(&stack)
	for _, val := range stack {
		fmt.Println(val)
	}
}
