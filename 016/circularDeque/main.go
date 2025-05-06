package main

// https://leetcode.cn/problems/design-circular-deque/description/

type MyCircularDeque struct {
	list  []int
	front int
	rear  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		list:  make([]int, k+1),
		front: 0,
		rear:  0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front - 1 + len(this.list)) % len(this.list)
	this.list[this.front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.list[this.rear] = value
	this.rear = (this.rear + 1) % len(this.list)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % len(this.list)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + len(this.list)) % len(this.list)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[(this.rear-1+len(this.list))%len(this.list)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.rear == this.front
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.rear+1)%len(this.list) == this.front
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
