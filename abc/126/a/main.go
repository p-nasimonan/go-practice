package main

import (
	"fmt"
	"unicode"
)

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k)
	fmt.Scan(&s)

	rs := []rune(s)
	if k >= 1 && k <= len(rs) {
		rs[k-1] = unicode.ToLower(rs[k-1])
	}

	fmt.Println(string(rs))

}
