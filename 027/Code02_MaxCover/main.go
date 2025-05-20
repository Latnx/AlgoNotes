// https://www.nowcoder.com/practice/1ae8d0b6bb4e4bcdbf64ec491f63fc37
// 线段的最多重合
// 判断左边界在当前pair的左边界前且右边界在当前边界后的pair有几个
// 因此排序左边界，对右边界使用最小堆，从前向后依次加入pair，小于等于左边界就Pop这样就可以保证左边在左边，右边在右边
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type pair [2]int
type pairList []pair
type pairHeap []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p pairHeap) Len() int            { return len(p) }
func (p pairHeap) Less(i, j int) bool  { return p[i][1] < p[j][1] }
func (p pairHeap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *pairHeap) Push(x interface{}) { (*p) = append(*p, x.(pair)) }
func (p *pairHeap) Pop() interface{}   { x := (*p)[len(*p)-1]; (*p) = (*p)[:len(*p)-1]; return x }

func main() {
	a := 0
	b := 0
	n := 0
	fmt.Scan(&n)
	var arr pairList = make([]pair, n)
	for i := range n {
		fmt.Scan(&a, &b)
		arr[i] = pair{a, b}
	}
	// 按照开始位置排序
	sort.Sort(arr)
	var hp pairHeap
	var res []int = make([]int, n)
	for i, a := range arr {
		for len(hp) > 0 && hp[0][1] <= a[0] {
			heap.Pop(&hp)
		}
		heap.Push(&hp, a)
		res[i] = len(hp)
	}
	var max = res[0]
	for _, r := range res {
		if r > max {
			max = r
		}
	}
	fmt.Print(max)
}
