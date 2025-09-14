package main

import "fmt"

// 用类描述实现前缀树。不推荐！
// 测试链接 : https://leetcode.cn/problems/implement-trie-ii-prefix-tree/

// void insert(String word)：添加word，可重复添加；
// void delete(String word)：删除word，如果word添加过多次，仅删除一次；
// boolean search(String word)：查询word是否在字典树中出现过(完整的出现过，前缀式不算)；
// int prefixNumber(String pre)：返回以字符串pre作为前缀的单词数量。
// 现在给定一个m，表示有m次操作，每次操作都为以上四种操作之一。
// 每次操作会给定一个整数op和一个字符串word，op代表一个操作码，
// 如果op为1，则代表添加word，op为2则代表删除word，
// op为3则代表查询word是否在字典树中，op为4代表返回以word为前缀的单词数量。

type trie struct {
	tree [26](*trie)
	pass int
	end  int
}

func Constructor() trie {
	return trie{}
}

func (t *trie) insert(word string) {
	for i := 0; i < len(word); i++ {
		p := word[i] - 'a'
		if t.tree[p] == nil {
			t.tree[p] = &trie{}
		}
		t.tree[p].pass++
		t = t.tree[p]
	}
	t.end++
}
func (t *trie) delete(word string) {
	for i := 0; i < len(word); i++ {
		p := word[i] - 'a'
		t.tree[p].pass--
		if t.tree[p].pass == 0 {
			t.tree[p] = nil
			return
		}
		t = t.tree[p]
	}
	t.end--
}
func (t *trie) search(word string) bool {
	for i := 0; i < len(word); i++ {
		p := word[i] - 'a'
		if t.tree[p] == nil {
			return false
		}
		t = t.tree[p]
	}
	return t.end != 0
}
func (t *trie) prefixNumber(pre string) int {
	for i := 0; i < len(pre); i++ {
		p := pre[i] - 'a'
		if t.tree[p] == nil {
			return 0
		}
		t = t.tree[p]
	}
	return t.pass
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
			t.insert(str)
		case 2:
			t.delete(str)
		case 3:
			if t.search(str) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case 4:
			fmt.Println(t.prefixNumber(str))
		}
	}
}
