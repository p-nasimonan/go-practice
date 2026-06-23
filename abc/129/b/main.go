package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var W = make([]int, n)
	for i := range W {
		fmt.Scan(&W[i])
	}
	var min int
	min = 10000
	for t := 1; t <= n-1; t++ {
		s1 := 0
		s2 := 0
		for _, w := range W[:t] {
			s1 += w
		}
		for _, w := range W[t:] {
			s2 += w
		}
		distance := int(math.Abs(float64(s1 - s2)))

		if distance < min {
			min = distance
		}
	}
	fmt.Println(min)
}
