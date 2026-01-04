package main

import (
	"bufio"
	"os"
	"strconv"
)

// 有依赖的背包(模版)
// 物品分为两大类：主件和附件
// 主件的购买没有限制，钱够就可以；附件的购买有限制，该附件所归属的主件先购买，才能购买这个附件
// 例如，若想买打印机或扫描仪这样的附件，必须先购买电脑这个主件
// 以下是一些主件及其附件的展示：
// 电脑：打印机，扫描仪 | 书柜：图书 | 书桌：台灯，文具 | 工作椅：无附件
// 每个主件最多有2个附件，并且附件不会再有附件，主件购买后，怎么去选择归属附件完全随意，钱够就可以
// 所有的物品编号都在1~m之间，每个物品有三个信息：价格v、重要度p、归属q
// 价格就是花费，价格 * 重要度 就是收益，归属就是该商品是依附于哪个编号的主件
// 比如一件商品信息为[300,2,6]，花费300，收益600，该商品是6号主件商品的附件
// 再比如一件商品信息[100,4,0]，花费100，收益400，该商品自身是主件(q==0)
// 给定m件商品的信息，给定总钱数n，返回在不违反购买规则的情况下最大的收益
// 测试链接 : https://www.luogu.com.cn/problem/P1064
// 测试链接 : https://www.nowcoder.com/practice/f9c6f980eeec43ef85be20755ddbeaf4
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	obj := make([][3]int, m+1)
	// 是否是主商品 有几个附件 附件是什么
	king := make([]bool, m+1)
	// fans := make([]int, m+1)
	follows := make([][]int, m+1)
	for i := range follows {
		follows[i] = make([]int, 0)
	}
	for i := 1; i <= m; i++ {
		// 价格、重要度、主件编号
		scanner.Scan()
		obj[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		obj[i][1], _ = strconv.Atoi(scanner.Text())
		obj[i][1] *= obj[i][0]
		scanner.Scan()
		obj[i][2], _ = strconv.Atoi(scanner.Text())

		if obj[i][2] == 0 {
			king[i] = true
		} else {
			follows[obj[i][2]] = append(follows[obj[i][2]], i)
		}
	}
	writer.WriteString(strconv.Itoa(compute1(n, m, obj, king, follows)))

}
func compute1(n, m int, obj [][3]int, king []bool, follows [][]int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 上一个展开的主商品的编号
	p := 0
	for i := 1; i <= m; i++ {
		if king[i] {
			for j := 0; j <= n; j++ {
				// 不选当前主商品
				dp[i][j] = dp[p][j]
				// 选当前主商品
				if j-obj[i][0] >= 0 {
					dp[i][j] = max(dp[p][j-obj[i][0]]+obj[i][1], dp[i][j])
					// 选择一个附件
					for _, ob := range follows[i] {
						if j-obj[i][0]-obj[ob][0] >= 0 {
							dp[i][j] = max(dp[p][j-obj[i][0]-obj[ob][0]]+obj[i][1]+obj[ob][1], dp[i][j])
						}
					}
					// 选择俩附件
					if len(follows[i]) == 2 && j-obj[i][0]-obj[follows[i][0]][0]-obj[follows[i][1]][0] >= 0 {
						dp[i][j] = max(dp[p][j-obj[i][0]-obj[follows[i][0]][0]-obj[follows[i][1]][0]]+obj[i][1]+obj[follows[i][0]][1]+obj[follows[i][1]][1], dp[i][j])
					}
				}
			}
			p = i
		}
	}
	return dp[p][n]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
