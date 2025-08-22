package main

// LRU缓存,函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
// https://leetcode.cn/problems/lru-cache/

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type LRUCache struct {
	m          map[int]*DLinkedNode
	head, tail *DLinkedNode
	n          int
	capacity   int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		m:        make(map[int]*DLinkedNode),
		tail:     &DLinkedNode{},
		head:     &DLinkedNode{},
		capacity: capacity,
		n:        0,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
func (cache *LRUCache) Get(key int) int {
	node, ok := cache.m[key]
	if !ok {
		return -1
	}
	// 删除 并添加
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = cache.head.next
	node.prev = cache.head
	cache.head.next.prev = node
	cache.head.next = node
	return node.value
}

// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value； 如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
func (cache *LRUCache) Put(key int, value int) {
	node, ok := cache.m[key]
	// 存在：更新
	if ok {
		node.value = value
		node.prev.next = node.next
		node.next.prev = node.prev
	} else {
		node = &DLinkedNode{key: key, value: value}
		cache.m[key] = node
		if cache.n >= cache.capacity {
			delete(cache.m, cache.tail.prev.key)
			cache.n--
			cache.tail.prev = cache.tail.prev.prev
			cache.tail.prev.next = cache.tail
		}
		cache.n++

	}
	node.next = cache.head.next
	node.prev = cache.head
	cache.head.next.prev = node
	cache.head.next = node
}
