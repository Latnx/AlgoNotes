package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range n {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}
	for i := range n {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteByte('\n')
	}
}

// scanner 14ms
// Scan 437ms
// Fscan 77ms
// reader 14ms
