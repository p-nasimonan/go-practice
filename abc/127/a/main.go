package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if 6 <= a && a <= 12 {
		b = int(b / 2)
	} else if a <= 5 {
		b = 0
	}
	fmt.Println(b)
}

/* より良い条件分岐

var a, b int
fmt.Scan(&a, &b)
if a <= 5 {
	fmt.Println(0); return
}
if a <= 12 {
	fmt.Println(b / 2); return
}
fmt.Println(b)
*/
