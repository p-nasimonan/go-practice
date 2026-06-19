package main

import (
	"fmt"
)

func main() {
	var p, q, r int
	min := 9999999999
	fmt.Scan(&p, &q, &r)

	if p+q < min {
		min = p + q
	}
	if p+r < min {
		min = p + r
	}
	if q+r < min {
		min = q + r
	}
	fmt.Println(min)
}
