package main

import (
	"fmt"
)

func main() {
	var num int64 = -12345
	fmt.Printf("Binary: %d\n", uint64(num)) // 强制转换为 uint64 以显示补码形式
}
