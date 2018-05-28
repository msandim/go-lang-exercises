package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type MyImage struct {
	width, height int
}

func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width-1, i.height-1)
}

func (i MyImage) At(x, y int) color.Color {
	value := uint8(x ^ y)
	return color.RGBA{value, value, 255, 255}
}

func main() {
	m := MyImage{width: 256, height: 256}
	pic.ShowImage(m)
}
