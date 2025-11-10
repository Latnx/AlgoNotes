package main

import "slices"

// 贴纸拼词
// 测试链接 : https://leetcode.cn/problems/stickers-to-spell-word/
func minStickers(stickers []string, target string) int {
	// 排序
	targetB := []byte(target)
	stickersB := make([][]byte, len(stickers))
	slices.Sort(targetB)
	visited := make(map[string]bool)
	for i := range stickers {
		stickersB[i] = []byte(stickers[i])
		slices.Sort(stickersB[i])
	}
	adj := make([][][]byte, 26)
	for i := range adj {
		adj[i] = make([][]byte, 0)
	}
	for i := range stickersB {
		for j := range stickersB[i] {
			adj[stickersB[i][j]-'a'] = append(adj[stickersB[i][j]-'a'], stickersB[i])

		}
	}
	// dfs
	queue := make([][]byte, 400)
	l, r := 0, 0
	queue[r] = targetB
	r++
	level := 1
	for l < r {
		for levelNum := r - l; levelNum != 0; levelNum-- {
			for _, stricker := range adj[queue[l][0]-'a'] {
				subRes := sub(queue[l], stricker)
				if visited[string(subRes)] {
					continue
				} else {
					visited[string(subRes)] = true
				}
				if len(subRes) == 0 {
					return level
				}
				queue[r] = subRes
				r++
			}
			l++
		}
		level++
	}
	return -1
}

func sub(s []byte, stricker []byte) []byte {
	resMap := make([]int, 26)
	for _, val := range s {
		resMap[val-'a']++
	}
	for _, val := range stricker {
		resMap[val-'a']--
	}
	res := []byte{}
	for i := range resMap {
		for j := 0; j < resMap[i]; j++ {
			res = append(res, byte(i+'a'))
		}
	}
	return res
}
func main() {
	minStickers([]string{"these", "guess", "about", "garden", "him"}, "atomher")
}
