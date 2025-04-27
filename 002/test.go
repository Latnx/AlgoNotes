package main

import (
	"fmt"
	"math"
	"math/rand"
)

func exchange(person []int) {
	for i := range person {
		var a int
		if person[i] == 0 {
			break
		}
		for {
			a = rand.Intn(len(person))
			if a != i {
				break
			}
		}
		person[i] -= 1
		person[a] += 1

	}
}

func cal(a []int) float64 {
	var sum float64 = 0
	var sumcai int = 0
	for i := range a {
		for j := range a {
			sum += math.Abs(float64(a[i]) - float64(a[j]))
		}
		sumcai += a[i]
	}
	fmt.Println(sumcai)
	return sum / float64(2*len(a)*sumcai)
}

func main() {
	var n int = 100
	var person []int = make([]int, n)
	for i := range person {
		person[i] += 100
	}
	for range 100000000000000 {
		exchange(person)
	}
	fmt.Print(person)
	fmt.Println(cal(person))
}
