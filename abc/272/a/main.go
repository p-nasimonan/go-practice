package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	var count int
	for _, a := range A {
		count += a
	}
	fmt.Printf("%d\n", count)
}
