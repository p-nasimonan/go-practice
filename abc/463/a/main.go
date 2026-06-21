package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	if x/y == 16.0/9.0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
