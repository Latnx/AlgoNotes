package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 设计有setAll功能的哈希表, O(1)
// https://www.nowcoder.com/practice/7c4559f138e74ceb9ba57d76fd169967

type pair struct {
	value     int
	timestamp int
}

// 时间戳
var count int = 0
var all int

var m map[int]pair = make(map[int]pair)

func put(x, y int) {
	m[x] = pair{value: y, timestamp: count + 1}
}

func get(x int) {
	v, ok := m[x]
	if !ok {
		fmt.Println(-1)
		return
	}
	if v.timestamp <= count {
		fmt.Println(all)
	} else {
		fmt.Println(v.value)
	}
}

func setAll(allNum int) {
	count++
	all = allNum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	opNum := 0
	fmt.Scan(&opNum)
	for i := 0; i < opNum; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		numStrs := strings.Split(input, " ")
		nums := make([]int, len(numStrs))
		for i, str := range numStrs {
			nums[i], _ = strconv.Atoi(str)
		}
		switch nums[0] {
		case 1:
			put(nums[1], nums[2])
		case 2:
			get(nums[1])
		case 3:
			setAll(nums[1])
		}
	}
}
