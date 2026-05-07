// 3つ数字受け取って5以上なら (Big) と出す
package main

import (
	"fmt"
)

func main() {
	var numbers [3]int
	fmt.Scan(&numbers[0], &numbers[1], &numbers[2])
	for _, number := range numbers {
		if number >= 5 {
			fmt.Println(number, ": Big")
		}
	}

}
