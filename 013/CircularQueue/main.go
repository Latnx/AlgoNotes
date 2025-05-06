package main

// https://leetcode.cn/problems/design-circular-queue
type MyCircularQueue struct {
	list   []int
	frount int
	rear   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		list: make([]int, k+1),
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.list[q.rear] = value
	q.rear = (q.rear + 1) % len(q.list)
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.frount = (q.frount + 1) % len(q.list)
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.list[q.frount]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.list[(q.rear+len(q.list)-1)%len(q.list)]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.frount == q.rear
}

func (q *MyCircularQueue) IsFull() bool {
	k := len(q.list)
	return (q.rear+1)%k == q.frount
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
