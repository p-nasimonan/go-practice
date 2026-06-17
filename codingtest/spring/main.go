// あるお店で商品を買い、お金を支払いました。このとき、お釣りとして返す硬貨の合計枚数が最も少なくなるような、各硬貨の枚数の組み合わせを求めるGoの関数を実装してください。
package main

import "fmt"

type Item struct {
	name  string
	price int
}

func (item Item) Oturi(paid int) map[int]int {
	COINS := []int{500, 100, 50, 10, 5, 1}
	oturi := paid - item.price
	var return_coins = make(map[int]int, 6)
	for _, coin := range COINS {
		if coin <= oturi {
			return_coins[coin] = oturi / coin
			oturi = oturi % coin
		}
	}
	return return_coins
}

func main() {
	otya := Item{name: "otya", price: 341}
	fmt.Println(otya.Oturi(700))
}
