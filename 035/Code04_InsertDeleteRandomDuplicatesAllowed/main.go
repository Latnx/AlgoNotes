package main

import "math/rand/v2"

// 381. O(1) 时间插入、删除和获取随机元素 - 允许重复
//  https://leetcode.cn/problems/insert-delete-getrandom-o1-duplicates-allowed/

type RandomizedCollection struct {
	m    map[int](map[int]struct{})
	list []int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		m:    make(map[int](map[int]struct{})),
		list: make([]int, 0),
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	_, ok := this.m[val]
	if !ok {
		this.m[val] = make(map[int]struct{})
	}
	this.m[val][len(this.list)] = struct{}{}
	this.list = append(this.list, val)
	return len(this.m[val]) == 1
}

func (this *RandomizedCollection) Remove(val int) bool {
	indexSet, ok := this.m[val]
	if !ok || len(indexSet) == 0 {
		return false
	}

	var index int
	for key := range indexSet {
		index = key
		break // 取到一个就停
	}
	delete(this.m[val], index)
	this.Fix(index)
	return true
}
func (this *RandomizedCollection) Fix(index int) {
	del_ele := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	delete(this.m[del_ele], len(this.list))
	if index != len(this.list) {
		this.list[index] = del_ele
		this.m[del_ele][index] = struct{}{}
	}
}
func (this *RandomizedCollection) GetRandom() int {
	return this.list[rand.IntN(len(this.list))]
}

func main() {
	obj := Constructor()
	obj.Insert(1)
	obj.Remove(1)
	obj.GetRandom()
}
