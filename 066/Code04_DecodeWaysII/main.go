package main

// 解码方法
// 测试链接 : https://leetcode.cn/problems/decode-ways/

// 基础dfs
func f1(s []byte, index int) int {
	if index >= len(s) {
		return 1
	}
	c := s[index]
	ans := 0
	// 单字符
	if c == '0' {
		return 0
	}
	if c == '*' {
		ans += 9 * f1(s, index+1)
	} else {
		ans += f1(s, index+1)
	}

	// 双字符
	if index+1 < len(s) {
		c2 := s[index+1]

		if c == '*' && c2 == '*' {
			// 11–19 (9) + 21–26 (6) = 15
			ans += 15 * f1(s, index+2)

		} else if c == '*' {
			// *X
			if c2 >= '0' && c2 <= '9' {
				// 10〜19
				ans += 1 * f1(s, index+2)
				if c2 <= '6' {
					// 20〜26
					ans += 1 * f1(s, index+2)
				}
			}

		} else if c2 == '*' {
			// X*
			if c == '1' {
				ans += 9 * f1(s, index+2)
			} else if c == '2' {
				ans += 6 * f1(s, index+2)
			}

		} else {
			// 普通数字
			num := (c-'0')*10 + (c2 - '0')
			if num >= 10 && num <= 26 {
				ans += f1(s, index+2)
			}
		}
	}
	return ans
}

// 增加缓存
var arr = make([]int, 100000)

func f2(s []byte, index int) int {
	if index >= len(s) {
		return 1
	}
	if arr[index] != -1 {
		return arr[index]
	}
	ans := 0
	c := s[index]
	// 单字符
	if c == '0' {
		return 0
	}
	if c == '*' {
		ans += 9 * f2(s, index+1)
	} else {
		ans += f2(s, index+1)
	}

	// 双字符
	if index+1 < len(s) {
		c2 := s[index+1]

		if c == '*' && c2 == '*' {
			// 11–19 (9) + 21–26 (6) = 15
			ans += 15 * f2(s, index+2)

		} else if c == '*' {
			// *X
			if c2 >= '0' && c2 <= '9' {
				// 10〜19
				ans += 1 * f2(s, index+2)
				if c2 <= '6' {
					// 20〜26
					ans += 1 * f2(s, index+2)
				}
			}

		} else if c2 == '*' {
			// X*
			if c == '1' {
				ans += 9 * f2(s, index+2)
			} else if c == '2' {
				ans += 6 * f2(s, index+2)
			}
		} else {
			// 普通数字
			num := (c-'0')*10 + (c2 - '0')
			if num >= 10 && num <= 26 {
				ans += f2(s, index+2)
			}
		}
	}
	arr[index] = ans % 1000000007
	return ans
}

func numDecodings(s string) int {
	for i := range arr {
		arr[i] = -1
	}
	return f2([]byte(s), 0) % 1000000007
}
