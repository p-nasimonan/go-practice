package main

import (
	"fmt"
)

func main() {
	var r, d, x2000 int
	const n int = 10
	fmt.Scan(&r, &d, &x2000)

	x := make([]int, n+1)
	x[0] = x2000
	for i := 0; i < n; i++ {
		x[i+1] = r*x[i] - d
	}

	for i := 1; i <= n; i++ {
		fmt.Println(x[i])
	}
}
