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

	// ピクセルごとにフィルター処理
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			var rSum, gSum, bSum float32
			var count float32

			// 自分の周り 5x5 のピクセルを探索する
			for dy := -2; dy <= 2; dy++ {
				for dx := -2; dx <= 2; dx++ {
					px := x + dx
					py := y + dy

					// 画像の端っこで外側に飛び出さないようにチェック
					if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
						r, g, b, _ := img.At(px, py).RGBA()
						rSum += float32(r >> 8)
						gSum += float32(g >> 8)
						bSum += float32(b >> 8)
						count++
					}
				}
			}

			// 合計割って平均を出す
			newColor := color.RGBA{
				uint8(rSum / count),
				uint8(gSum / count),
				uint8(bSum / count),
				255,
			}
			dest.Set(x, y, newColor)
		}
	}

	// 5. 保存する
	outFile, _ := os.Create("output.png")
	defer outFile.Close()
	png.Encode(outFile, dest)

	fmt.Println("フィルター処理が完了しました！ output.png を確認してみてね。")
}
