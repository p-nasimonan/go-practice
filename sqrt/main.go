package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	low, high := 0.0, x
	for i := 0; i < 100; i++ {
		mid := (low + high) / 2
		result := mid * mid
		if result < x {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	fmt.Printf("math:%f my:%f\n", math.Sqrt(32), Sqrt(32))
}
