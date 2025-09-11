package main

import "fmt"

// 打表：输入和输出都是简单类型，

// 可以用r、e、d三种字符拼接字符串，如果拼出来的字符串中
// 有且仅有1个长度>=2的回文子串，那么这个字符串定义为"好串"
// 返回长度为n的所有可能的字符串中，好串有多少个
// 结果对 1000000007 取模， 1 <= n <= 10^9
// 示例：
// n = 1, 输出0
// n = 2, 输出3
// n = 3, 输出18
var cntTotal int

func dfs(path []byte, index int) {
	if index == len(path) {
		var cnt int
		for i := range path {
			for j := i + 1; j < len(path) && cnt <= 1; j++ {
				is := true
				for k := i; k <= j; k++ {
					if path[k] != path[j-k+i] {
						is = false
					}
				}
				if is {
					cnt++
				}
			}
		}
		if cnt == 1 {
			cntTotal++
		}
		return
	}
	for _, val := range []byte{'R', 'E', 'D'} {
		path[index] = val
		dfs(path, index+1)
		path[index] = 0
	}
}

func red(n int) (res int) {
	cntTotal = 0
	path := make([]byte, n)
	dfs(path, 0)
	return
}
func main() {
	for i := range 30 {
		red(i)
		fmt.Println(i, cntTotal)
	}
}
