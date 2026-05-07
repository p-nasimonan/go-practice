// fmt.Scanfで"1 2 3"のような数を受け取り変数に代入し、Printfで表示する
package main

import (
	"fmt"
)

func main() {
	var d, e, f int
	const ADD_VALUE int = 1
	fmt.Scan(&d, &e, &f)
	fmt.Printf("%d %d %d\n", d+ADD_VALUE, e+ADD_VALUE, f+ADD_VALUE)
}
