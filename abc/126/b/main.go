package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	if len(s) < 4 {
		return
	}
	a := s[:2]
	b := s[2:4]

	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)

	isMonthA := 1 <= ai && ai <= 12
	isMonthB := 1 <= bi && bi <= 12

	if isMonthA && isMonthB {
		fmt.Println("AMBIGUOUS")
		return
	}
	if isMonthA {
		fmt.Println("MMYY")
		return
	}
	if isMonthB {
		fmt.Println("YYMM")
		return
	}
	fmt.Println("NA")
}
