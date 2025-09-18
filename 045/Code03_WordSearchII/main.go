package main

import (
	"fmt"
)

// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words
// 返回所有二维网格上的单词。单词必须按照字母顺序，通过 相邻的单元格 内的字母构成
// 其中“相邻”单元格是那些水平相邻或垂直相邻的单元格
// 同一个单元格内的字母在一个单词中不允许被重复使用
// 1 <= m, n <= 12
// 1 <= words.length <= 3 * 10^4
// 1 <= words[i].length <= 10
// 测试链接 : https://leetcode.cn/problems/word-search-ii/

// 当前做法为遍历每一个字符串，因此对于每个字符串都要对棋盘进行遍历，
// 使用前缀树进行剪枝，先将所有字符串存入前缀树中，再对棋盘进行遍历，去前缀树中查找
// path 中存放的是路径
// trieWord 存放字典树上的单词，在遍历到即可收集

var trieWord []string = make([]string, 10000)
var tree [][26]int = make([][26]int, 10000)
var cnt = 1
var res []string = make([]string, 0)

func Insert(word string) {
	next := 1
	for _, val := range word {
		path := val - 'a'
		if tree[next][path] == 0 {
			cnt++
			tree[next][path] = cnt
		}
		next = tree[next][path]
	}
	trieWord[next] = word
}

// 更新记录当前前缀是否存在下一个字母，记录后删除，防止重复记录
func StartWith(pre byte, index int) (bool, int) {
	path := pre - 'a'
	if tree[index][path] == 0 {
		return false, 0
	}
	index = tree[index][path]
	if trieWord[index] != "" {
		res = append(res, trieWord[index])
	}
	trieWord[index] = ""
	return true, index
}

// 现在index表示第一个节点
func findWord(board [][]byte, i, j int, index int) {
	if (i < 0 || i >= len(board)) || (j < 0 || j >= len(board[0])) || board[i][j] == '#' {
		return
	}
	ok, index := StartWith(board[i][j], index)
	if !ok {
		return
	}
	tmp := board[i][j]
	board[i][j] = '#'
	findWord(board, i-1, j, index)
	findWord(board, i, j-1, index)
	findWord(board, i+1, j, index)
	findWord(board, i, j+1, index)
	board[i][j] = tmp
}

func clear() {
	cnt = 1
	res = make([]string, 0)
	for i := range trieWord {
		trieWord[i] = ""
	}
	for i := range tree {
		for j := range tree[i] {
			tree[i][j] = 0
		}
	}
}

func findWords(board [][]byte, words []string) []string {
	clear()
	// 建立一个树,方便查询
	for _, word := range words {
		Insert(word)
	}
	// 遍历棋盘
	for i := range board {
		for j := range board[i] {
			findWord(board, i, j, 1)
		}
	}
	return res
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	words := []string{"oath", "pea", "eat", "rain"}

	fmt.Print(findWords(board, words))
}
