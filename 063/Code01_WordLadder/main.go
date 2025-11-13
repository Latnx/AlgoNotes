package main

import "fmt"

// 单词接龙
// 测试链接 : https://leetcode.cn/problems/word-ladder/
var dict map[string]struct{}
var visited map[string]bool
var curLevel map[string]struct{}
var nextLevel map[string]struct{}

func build(wordList []string) {
	dict = map[string]struct{}{}
	for _, word := range wordList {
		dict[word] = struct{}{}
	}
	visited = map[string]bool{}
	curLevel = map[string]struct{}{}
	nextLevel = map[string]struct{}{}
}

func next(smallLevel *map[string]struct{}, bigLevel map[string]struct{}) bool {
	for word := range *smallLevel {
		visited[word] = true
	}
	for word := range *smallLevel {
		ch := []byte(word)
		for i := range ch {
			for j := byte('a'); j <= 'z'; j++ {
				old := ch[i]
				ch[i] = j
				str := string(ch)
				if _, ok := dict[str]; ok && !visited[str] && str != word {
					if _, ok := bigLevel[str]; ok {
						return false
					}
					if _, ok := nextLevel[str]; !ok {
						nextLevel[str] = struct{}{}
					}
				}
				ch[i] = old
			}
		}
	}
	*smallLevel, nextLevel = nextLevel, *smallLevel
	clear(nextLevel)
	return true
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	build(wordList)
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	smallLevel := map[string]struct{}{beginWord: {}}
	bigLevel := map[string]struct{}{endWord: {}}
	res := 2
	for next(&smallLevel, bigLevel) {
		if len(smallLevel) > len(bigLevel) {
			smallLevel, bigLevel = bigLevel, smallLevel
		}
		if len(smallLevel) == 0 {
			return 0
		}
		res++
	}
	return res
}

func main() {
	wordList1 := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord1 := "hit"
	endWord1 := "cog"
	fmt.Printf("Example 1: %v\n", ladderLength(beginWord1, endWord1, wordList1))
}
