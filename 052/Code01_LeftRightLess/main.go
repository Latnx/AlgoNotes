package main

// 单调栈求每个位置左右两侧，离当前位置最近、且值严格小于的位置
// 给定一个可能含有重复值的数组 arr
// 找到每一个 i 位置左边和右边离 i 位置最近且值比 arr[i] 小的位置
// 返回所有位置相应的信息。

// 单调栈可以获得所有最近且小于的某一位置 O(n)
// 准备一个栈, 栈里放下标, 从栈底到栈顶数字由小到大(大压小)

// 测试链接 : https://www.nowcoder.com/practice/2a2c00e7a88a498693568cef63a4b7bb
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 处理输入输出
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	// 准备空间
	stack := make([]int, len(arr))
	r := 0
	leftAns := make([]int, len(arr))
	rightAns := make([]int, len(arr))

	// 遍历
	for index, val := range arr {
		for r > 0 && val <= arr[stack[r-1]] {
			r--
			cur := stack[r]
			if r == 0 {
				leftAns[cur] = -1
			} else {
				leftAns[cur] = stack[r-1]
			}
			rightAns[cur] = index
		}
		stack[r] = index
		r++
	}
	// 清算stack
	for r > 0 {
		r--
		cur := stack[r]
		if r == 0 {
			leftAns[cur] = -1
		} else {
			leftAns[cur] = stack[r-1]
		}
		rightAns[cur] = -1
	}

	// 修正,如果记录的最近最小的右边与自己相等的话,是由下一坐标
	for i := n - 2; i >= 0; i-- {
		if rightAns[i] != -1 && arr[i] == arr[rightAns[i]] {
			rightAns[i] = rightAns[rightAns[i]]
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, leftAns[i], rightAns[i])
	}
}
