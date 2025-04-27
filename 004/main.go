package main

import "fmt"

func selectionSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	var n = len(arr)
	for i := range len(arr) - 1 {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func bubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	var n = len(arr)
	for end := n - 1; end > 0; end-- {
		for j := 0; j < end; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func insertSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
func main() {
	// arr := selectionSort([]int{10, 9, 321312, 321312, 49})
	// arr := bubbleSort([]int{10, 9, 321312, 321312, 49})
	arr := insertSort([]int{10, 9, 321312, 321312, 49})
	fmt.Print(arr)
}
