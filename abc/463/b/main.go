package main

import (
	"fmt"
)

func main() {
	var n int
	var x string
	fmt.Scan(&n, &x)
	ques := make(map[string]struct{})
	var S string

	var que int
	switch x {
	case "A":
		que = 0
	case "B":
		que = 1
	case "C":
		que = 2
	case "D":
		que = 3
	case "E":
		que = 4
	default:
		que = 0
	}
	ok := false
	for i := 0; i < n; i++ {
		fmt.Scan(&S)
		for j, s := range S {
			if j == que {
				ques[string(s)] = struct{}{}
			}
		}
	}
	_, ok = ques["o"]
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
