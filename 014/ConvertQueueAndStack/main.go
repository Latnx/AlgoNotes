package main

// https://leetcode.cn/problems/implement-queue-using-stacks/submissions/

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		for range len(q.inStack) {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}
	num := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return num
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		for range len(q.inStack) {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}
	num := q.outStack[len(q.outStack)-1]
	// q.outStack = q.outStack[:len(q.outStack)-1]
	return num
}

func (q *MyQueue) Empty() bool {
	if len(q.outStack) == 0 && len(q.inStack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
