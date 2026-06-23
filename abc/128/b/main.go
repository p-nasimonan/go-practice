package main

import (
	"fmt"
	"slices"
	"strings"
)

type book struct {
	i0    int
	city  string
	point int
}

func main() {
	var n int
	fmt.Scan(&n)
	var books = make([]book, n)
	for i := range books {
		books[i].i0 = i + 1
		fmt.Scan(&books[i].city, &books[i].point)
	}
	slices.SortFunc(books, func(a, b book) int {
		ans := strings.Compare(a.city, b.city)
		if ans == 0 {
			return b.point - a.point
		}
		return ans
	})

	for i := range books {
		fmt.Println(books[i].i0)
	}
}
