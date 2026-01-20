package main

import (
	"bufio"
	"os"
	"strconv"
)

// 选课，树上01背包的普通解法
// 在大学里每个学生，为了达到一定的学分，必须从很多课程里选择一些课程来学习
// 在课程里有些课程必须在某些课程之前学习，如高等数学总是在其它课程之前学习
// 现在有 N 门功课，每门课有个学分，每门课有一门或没有直接先修课
// 若课程 a 是课程 b 的先修课即只有学完了课程 a，才能学习课程 b
// 一个学生要从这些课程里选择 M 门课程学习
// 问他能获得的最大学分是多少
// 测试链接 : https://www.luogu.com.cn/problem/P2014
// 请同学们务必参考如下代码中关于输入、输出的处理
// 这是输入输出处理效率很高的写法
// 提交以下的code，提交时请把类名改成"Main"，可以直接通过
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	graph := make([][]int, N+1)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	scores := make([]int, N+1)
	for i := range N {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		graph[k] = append(graph[k], i+1)
		scanner.Scan()
		s, _ := strconv.Atoi(scanner.Text())
		scores[i+1] = s
	}
	dp = make([][][]int, N+1)
	for i := range dp {
		dp[i] = make([][]int, len(graph[i])+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, M+2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	writer.WriteString(strconv.Itoa(dfs(graph, scores, 0, len(graph[0]), M+1)))
}

// dp[i][j][k]: 以i为头，在前j棵子树里选k门
var dp [][][]int

func dfs(graph [][]int, scores []int, i, j, k int) int {
	if k == 0 {
		return 0
	}
	if j == 0 || k == 1 {
		return scores[i]
	}
	if dp[i][j][k] != -1 {
		return dp[i][j][k]
	}
	// 不选j
	ans := dfs(graph, scores, i, j-1, k)
	// 选j,v是第j棵子树的头节点
	v := graph[i][j-1]
	for s := 1; s < k; s++ {
		ans = max(ans, dfs(graph, scores, i, j-1, k-s)+dfs(graph, scores, v, len(graph[v]), s))
	}
	dp[i][j][k] = ans
	return ans
}
