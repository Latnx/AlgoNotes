package main

import "slices"

// 单词接龙 II
// 测试链接 : https://leetcode.cn/problems/word-ladder-ii/

// 找最短的路径
// 最短（bfs），路径（dfs）

// var dict map[string]struct{}
var graph map[string][]string
var visited map[string]bool
var curLevel map[string]struct{}
var nextLevel map[string]struct{}

var path []string
var res [][]string

func build(wordList []string) {
	graph = make(map[string][]string)
	for _, word := range wordList {
		graph[word] = make([]string, 0)
	}
	visited = map[string]bool{}
	path = make([]string, 0)
	res = make([][]string, 0)
	curLevel = make(map[string]struct{})
	nextLevel = map[string]struct{}{}
}

// 1.begin->end层层bfs
// 1.1 找到开始begin，并添加与begin相差一个字母的，bfs
// 1.2 直到找到end为止
// 2.建图，dfs建立path

// 剪枝
// 1. 在字典中删除（防重复）
// 2. 反向建图，减少dfs展开

func bfs(begin, end string) bool {
	find := false
	curLevel[begin] = struct{}{}
	for len(curLevel) != 0 {
		for word := range curLevel {
			visited[word] = true
		}
		for word := range curLevel {
			ch := []byte(word)
			for i := range ch {
				for j := byte('a'); j <= 'z'; j++ {
					old := ch[i]
					ch[i] = j
					str := string(ch)
					if _, ok := graph[str]; ok && !visited[str] && str != word {
						if str == end {
							find = true
						}
						graph[str] = append(graph[str], word)
						if _, ok := nextLevel[str]; !ok {
							nextLevel[str] = struct{}{}
						}
					}
					ch[i] = old
				}
			}
		}
		if find {
			return find
		}
		curLevel = nextLevel
		nextLevel = map[string]struct{}{}
	}
	return find
}

func dfs(word, aim string) {
	path = append(path, word)
	if word == aim {
		res = append(res, make([]string, len(path)))
		copy(res[len(res)-1], path)
		slices.Reverse(res[len(res)-1])
	} else {
		for _, next := range graph[word] {
			dfs(next, aim)
		}
	}
	path = path[:len(path)-1]
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	build(wordList)
	if _, ok := graph[endWord]; !ok {
		return res
	}
	if bfs(beginWord, endWord) {
		dfs(endWord, beginWord)
	}
	return res
}

func main() {
	findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
}
