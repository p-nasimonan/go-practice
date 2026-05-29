package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	width := 300
	height := 300
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 色塗る
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			val := (float32(x)/float32(width) + float32(y)/float32(height)) / 2 * 255
			img.Set(x, y, color.RGBA{0, uint8(val), 0, 255})
		}
	}

	f, err := os.Create("green_image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// PNG形式で書き出し
	png.Encode(f, img)

}
