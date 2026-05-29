package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw" // JPEGを扱うために必要
	"image/png"  // PNGを扱うために必要
	"os"
)

func main() {
	file, err := os.Open("image.png")
	if err != nil {
		fmt.Println("画像ひらけなかった")
	}
	defer file.Close()

	// 2. 画像をデコード（解析）する
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("画像のデコードに失敗しました:", err)
		return
	}

	bounds := img.Bounds()
	dest := image.NewRGBA(bounds)
	draw.Draw(dest, bounds, img, bounds.Min, draw.Src)

	// 4. ピクセルごとにフィルター処理
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 元の色を取得
			oldColor := dest.At(x, y)
			r, g, b, a := oldColor.RGBA()

			// GoのRGBA()は0〜65535の範囲なので、0〜255に変換
			R := float32(r >> 8)
			G := float32(g >> 8)
			B := float32(b >> 8)

			// --- フィルターアルゴリズム：モノクロ（輝度計算） ---
			// 人間の目に自然に見える輝度の比率：R:0.299, G:0.587, B:0.114
			gray := uint8(R*0.299 + G*0.587 + B*0.114)
			newColor := color.RGBA{gray, gray, gray, uint8(a >> 8)}
			// ----------------------------------------------

			dest.Set(x, y, newColor)
		}
	}

	// 5. 保存する
	outFile, _ := os.Create("output.png")
	defer outFile.Close()
	png.Encode(outFile, dest)

	fmt.Println("フィルター処理が完了しました！ output.png を確認してみてね。")
}
