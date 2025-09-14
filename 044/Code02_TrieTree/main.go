package main

import "fmt"

// 用固定数组实现前缀树，空间使用是静态的。推荐！
// 测试链接 : https://www.nowcoder.com/practice/7f8a8553ddbf4eaab749ec988726702b

const MAXN = 150001

type Trie struct {
	tree [][26]int
	pass []int
	end  []int
}

var cnt int

// 从1开始计数，0表示无数据
func Constructor() Trie {
	cnt = 1
	return Trie{
		make([][26]int, MAXN),
		make([]int, MAXN),
		make([]int, MAXN),
	}
}

func (t *Trie) Insert(word string) {
	next := 1
	for _, val := range word {
		p := val - 'a'
		if t.tree[next][p] == 0 {
			cnt++
			t.tree[next][p] = cnt
		}
		next = t.tree[next][p]
		t.pass[next]++
	}
	t.end[next]++
}

func (t *Trie) Delete(word string) {
	next := 1
	for _, val := range word {
		p := val - 'a'
		next = t.tree[next][p]
		t.pass[next]--
	}
	t.end[next]--
}
func (t *Trie) Search(word string) bool {
	next := 1
	for _, val := range word {
		p := val - 'a'
		next = t.tree[next][p]
		if next == 0 {
			return false
		}
	}
	return t.end[next] != 0
}
func (t *Trie) StartsWith(pre string) int {
	next := 1
	for _, val := range pre {
		p := val - 'a'
		next = t.tree[next][p]
		if next == 0 {
			return 0
		}
	}
	return t.pass[next]
}
func main() {
	n := 0
	t := Constructor()
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		op, str := 0, ""
		fmt.Scan(&op, &str)
		switch op {
		case 1:
			t.Insert(str)
		case 2:
			t.Delete(str)
		case 3:
			if t.Search(str) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case 4:
			fmt.Println(t.StartsWith(str))
		}
	}
}
