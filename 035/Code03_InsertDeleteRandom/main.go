package main

import "math/rand/v2"

// 插入、删除和获取随机元素O(1)时间的结构
// https://leetcode.cn/problems/insert-delete-getrandom-o1/description/
// 集合则需要set，随机需要list
type RandomizedSet struct {
	m    map[int]int
	list []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    make(map[int]int),
		list: make([]int, 0),
	}
}

func (randomizedSet *RandomizedSet) Insert(val int) bool {
	_, ok := randomizedSet.m[val]
	if ok {
		return false
	}
	// 始终插入在最后一个位置
	randomizedSet.list = append(randomizedSet.list, val)
	// 获取最后一个位置的下标
	randomizedSet.m[val] = len(randomizedSet.m)
	return true
}

func (randomizedSet *RandomizedSet) Remove(val int) bool {
	index, ok := randomizedSet.m[val]
	if !ok {
		return false
	}
	delete(randomizedSet.m, val)
	del_ele := randomizedSet.list[len(randomizedSet.list)-1]
	randomizedSet.list = randomizedSet.list[:len(randomizedSet.list)-1]
	if index != len(randomizedSet.list) {
		randomizedSet.list[index] = del_ele
		randomizedSet.m[randomizedSet.list[index]] = index
	}

	return true
}

func (randomizedSet *RandomizedSet) GetRandom() int {
	return randomizedSet.list[rand.IntN(len(randomizedSet.list))]
}

func main() {
	obj := Constructor()
	obj.Insert(11)
	obj.Remove(11)
	obj.GetRandom()
}
