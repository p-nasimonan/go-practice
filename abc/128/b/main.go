// T1秒後にA待ち、その後ろにT2秒後の客がA秒待ち、、客２の客１を待つ時間は客1の待つ時間-T2でT2が客１の待つ時間より大きいときは待つ時間は0、でそこに+来る時間とA秒まつ
package main

import "fmt"

func main() {
	var N, A int
	fmt.Scan(&N, &A)
	var resultT = make([]int, N)
	var T = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&T[i])
	}
	waitT := 0
	for i := 0; i < N; i++ {
		if i != 0 {
			waitT = resultT[i-1] - T[i]
		}
		if waitT < 0 {
			waitT = 0
		}
		resultT[i] = waitT + T[i] + A
	}

	for _, t := range resultT {
		fmt.Println(t)
	}
}
