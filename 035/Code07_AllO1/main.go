package main

import (
	"container/list"
)

// 设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。
// 双向链表加哈希表
// https://leetcode.cn/problems/all-oone-data-structure/
// 双向链表作词频桶。
// 因为返回最大和最小，但又因为O(1)不能使用堆结构
// 因为对频率的更改每次增加1,或者减少1,因此可以使用双向链表。
// 双向链表连接频率桶，每个桶中存放对应频率的字符串,这样最大和最小在两端,而增加和减少只需要左移和右移节点
// 为了快速查找到字符串对应的节点创建一个哈希表作为映射

type node struct {
	count int
	keys  map[string]struct{}
}

type AllOne struct {
	*list.List
	nodes map[string]*list.Element
}

// 双向链表添加头尾
func Constructor() AllOne {
	l := list.New()
	return AllOne{List: l, nodes: map[string]*list.Element{}}
}

func (this *AllOne) Inc(key string) {
	index, ok := this.nodes[key]
	// key 不存在：要放到 count=1 的桶里（新建或复用）
	// key 已存在：要从原桶移到 count+1 的桶里（新建或复用）
	// 在必要时删掉旧桶
	if !ok {
		first := this.List.Front()
		// 新建桶
		if first == nil || first.Value.(*node).count != 1 {
			keys := map[string]struct{}{}
			keys[key] = struct{}{}
			this.nodes[key] = this.List.PushFront(&node{count: 1, keys: keys})
		} else {
			this.List.Front().Value.(*node).keys[key] = struct{}{}
			this.nodes[key] = this.Front()
		}
	} else {
		// 新建桶
		if index.Next() == nil || index.Next().Value.(*node).count != index.Value.(*node).count+1 {
			keys := map[string]struct{}{}
			keys[key] = struct{}{}
			this.nodes[key] = this.List.InsertAfter(&node{count: index.Value.(*node).count + 1, keys: keys}, index)
		} else {
			index.Next().Value.(*node).keys[key] = struct{}{}
			this.nodes[key] = index.Next()
		}
		delete(index.Value.(*node).keys, key)
		if len(index.Value.(*node).keys) == 0 {
			this.List.Remove(index)
		}
	}
}

func (this *AllOne) Dec(key string) {
	// 测试用例保证存在
	// 如果已经在桶1中,则直接删除
	// 要从原桶移到 count+1 的桶里（新建或复用）,在必要时删掉旧桶
	index := this.nodes[key]
	if index.Value.(*node).count == 1 {
		delete(this.nodes, key)
	} else if index.Prev() == nil || index.Prev().Value.(*node).count != index.Value.(*node).count-1 {
		keys := map[string]struct{}{}
		keys[key] = struct{}{}
		this.nodes[key] = this.List.InsertBefore(&node{count: index.Value.(*node).count - 1, keys: keys}, index)
	} else {
		this.nodes[key] = index.Prev()
		index.Prev().Value.(*node).keys[key] = struct{}{}
	}
	delete(index.Value.(*node).keys, key)
	if len(index.Value.(*node).keys) == 0 {
		this.List.Remove(index)
	}
}

func (this *AllOne) GetMinKey() string {
	front := this.List.Front()
	if front == nil {
		return ""
	}
	for k := range front.Value.(*node).keys {
		return k
	}
	return ""
}

func (this *AllOne) GetMaxKey() string {
	back := this.List.Back()
	if back == nil {
		return ""
	}
	for k := range back.Value.(*node).keys {
		return k
	}
	return ""
}

func main() {
	obj := Constructor()
	obj.Inc("a")
	obj.Inc("b")
	obj.Dec("b")
	obj.GetMaxKey()
	obj.GetMinKey()
}
