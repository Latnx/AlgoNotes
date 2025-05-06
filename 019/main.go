package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // 读取整行输入
	input = strings.TrimSpace(input)
	// 去除换行符和空格
	// 解析整数数组
	numStrs := strings.Split(input, " ")
	nums := make([]int, len(numStrs))
	for i, str := range numStrs {
		nums[i], _ = strconv.Atoi(str)
	}
	fmt.Println("Parsed numbers:", nums)
}
