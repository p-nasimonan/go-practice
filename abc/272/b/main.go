package main

import (
	"fmt"
)

func main() {
	var n int //人数
	var m int //舞踏会の回数
	fmt.Scan(&n, &m)
	var k int //その舞踏会に「何人」来たか

	matrix := make([][]bool, n)
	for i := range matrix {
		matrix[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&k)
		x := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&x[j])
		}

		for a := 0; a < k; a++ {
			for b := a + 1; b < k; b++ {
				// ここで -1 して、両方向を true にする
				p1, p2 := x[a]-1, x[b]-1
				matrix[p1][p2] = true
				matrix[p2][p1] = true
			}
		}
	}
	ans := "Yes"
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if !matrix[i][j] {
				ans = "No"
			}
		}
	}
	fmt.Println(ans)
}
