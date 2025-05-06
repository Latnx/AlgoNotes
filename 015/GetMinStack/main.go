package main

// https://leetcode.cn/problems/min-stack
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minStack) > 0 {
		min := s.minStack[len(s.minStack)-1]
		if val < min {
			min = val
		}
		s.minStack = append(s.minStack, min)
	} else {
		s.minStack = append(s.minStack, val)
	}

}

func (s *MinStack) Pop() {
	s.stack = s.stack[0 : len(s.stack)-1]
	s.minStack = s.minStack[0 : len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	// fmt.Print(this.minStack)
	return s.minStack[len(s.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
