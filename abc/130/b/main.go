/*
数直線上を N+1 回跳ねるボールがあり、1 回目は 座標 D 1 =0, i 回目は 座標 D i =D i−1 +L i−1 (2≤i≤N+1) で跳ねます。
数直線の座標が
X 以下の領域でボールが跳ねる回数は何回でしょうか。
n x
3 6
l0 l1 l2
3  4  5
d0 d1  d2 d3
0  3+0 4+3 5+7
要はN回は寝るボールが0回目0、1回目移行からL0,L1,,,Lnが始まる
*/

package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	var L = make([]int, n)
	for i := range L {
		fmt.Scan(&L[i])
	}
	var BoalXs = make([]int, n+1)
	BoalXs[0] = 0
	ans := 1

	for i := 1; i <= n; i++ {
		BoalXs[i] = BoalXs[i-1] + L[i-1]

		if BoalXs[i] <= x {
			ans++
		}
	}

	fmt.Println(ans)
}
