package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var baseRes int = 256

type Image struct {
	res int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	res := baseRes * m.res
	return image.Rect(0, 0, res, res)
}

func (m Image) At(x, y int) color.Color {
	v := uint8((x*x + y*y) / m.res)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{4}
	err := png.Encode(os.Stdout, m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Dumped to Stdout\n")
}
