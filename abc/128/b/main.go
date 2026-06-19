package main

import (
	"fmt"
	"maps"
	"slices"
)

func makelanking(N int, S []string, P []int) []int {
	type CityRank struct {
		Name  string
		Ranks []int
	}
	tmpcitys := map[string]struct{}{}
	for i := 0; i < N; i++ {
		tmpcitys[S[i]] = struct{}{}
	}
	citys := slices.Collect(maps.Keys(tmpcitys))
	cityranks := make([]CityRank, len(citys))
	for i := 0; i < N; i++ {
		if S[i] 
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	var S = make([]string, N)
	var P = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&S[i], &P[i])
	}
	fmt.Println(makelanking(N, S, P))
}
