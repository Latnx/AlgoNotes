package main

// 895. 最大频率栈
// https://leetcode.cn/problems/maximum-frequency-stack/

type FreqStack struct {
	freqMap  map[int]int
	freqList [][]int
}

func Constructor() FreqStack {
	return FreqStack{
		freqMap:  map[int]int{},
		freqList: make([][]int, 0),
	}
}

func (this *FreqStack) Push(val int) {
	index := this.freqMap[val]
	if len(this.freqList) == index {
		this.freqList = append(this.freqList, make([]int, 0))
	}
	this.freqList[index] = append(this.freqList[index], val)
	this.freqMap[val]++
}

func (this *FreqStack) Pop() int {
	last_index := len(this.freqList[len(this.freqList)-1]) - 1
	x := this.freqList[len(this.freqList)-1][last_index]
	this.freqList[len(this.freqList)-1] = this.freqList[len(this.freqList)-1][:last_index]
	if len(this.freqList[len(this.freqList)-1]) == 0 {
		this.freqList = this.freqList[:len(this.freqList)-1]
	}
	this.freqMap[x]--
	return x
}

func main() {
	obj := Constructor()
	obj.Push(5)
	obj.Push(7)
	obj.Push(5)
	obj.Push(7)
	obj.Push(4)
	obj.Push(5)
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
}
