package main

// 每个元音包含偶数次的最长子字符串
// 给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度
// 每个元音字母，即 'a'，'e'，'i'，'o'，'u'
// 在子字符串中都恰好出现了偶数次。
// 测试链接 : https://leetcode.cn/problems/find-the-longest-substring-containing-vowels-in-even-counts/
// aeiou
const a = 0b00001
const e = 0b00010
const i = 0b00100
const o = 0b01000
const u = 0b10000

func findTheLongestSubstring(s string) int {
	preSumMap := make(map[int]int)
	preSumMap[0] = -1
	sum := 0
	maxLenth := 0
	for k, v := range s {
		switch v {
		case 'a':
			sum ^= a
		case 'e':
			sum ^= e
		case 'i':
			sum ^= i
		case 'o':
			sum ^= o
		case 'u':
			sum ^= u
		}
		if index, ok := preSumMap[sum]; ok {
			maxLenth = max(maxLenth, k-index)
		} else {
			preSumMap[sum] = k
		}
	}
	return maxLenth
}

func main() {
	findTheLongestSubstring("eleetminicoworoep")
}
