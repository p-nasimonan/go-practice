/*
あなたはカフェの注文を処理するシステムを作っています。
1日の営業が終わった（あるいは途中の）タイミングで、テーブルごとの合計請求金額を計算する関数を作成してください。
最終的な「各テーブル番号」と「そのテーブルの合計請求金額」の対応表を返してください。
商品メニュー:　商品名 値段
注文の履歴ログ: テーブルid 商品名

出力
各テーブルの合計 テーブルid 合計金額
*/

package main

import "fmt"

type Log struct {
	tableId  int
	itemName string
}

func sumLogs(logs []Log, menu map[string]int) map[int]int {
	var sumlogs = make(map[int]int, len(logs))
	for _, log := range logs {
		price, exists := menu[log.itemName]
		if !exists {
			continue // なければスキップ
		}
		sumlogs[log.tableId] += price
	}
	return sumlogs
}

func main() {
	menu := map[string]int{
		"tea":    130,
		"coffee": 140,
		"rabbit": 500,
		"chino":  600,
	}
	var N int
	fmt.Scan(&N)
	logs := make([]Log, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&logs[i].tableId, &logs[i].itemName)
	}
	fmt.Println(sumLogs(logs, menu))
}
