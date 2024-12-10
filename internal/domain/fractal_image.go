package domain

import (
	"image"
	"image/color"
	"sync"
)

type Image interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x int, y int) color.Color
}

type FractalImage struct {
	height int
	width  int
	Img    [][]Pixel
}

func NewFractalImage(height, width int) *FractalImage {
	fractalImg := FractalImage{
		height: height,
		width:  width,
		Img:    make([][]Pixel, height),
	}

	for i := range fractalImg.Img {
		fractalImg.Img[i] = make([]Pixel, width)
		for j := range fractalImg.Img[i] {
			fractalImg.Img[i][j] = *NewPixel()
		}
	}

	return &fractalImg
}
func (f FractalImage) GetHeight() int {
	return f.height
}
func (f FractalImage) GetWidth() int {
	return f.width
}

func (f FractalImage) ColorModel() color.Model {
	return color.RGBAModel
}
func (f FractalImage) At(x, y int) color.Color {
	return f.Img[y][x].Color
}
func (f FractalImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, f.GetWidth(), f.GetHeight())
}

type Pixel struct {
	Color  color.RGBA
	Normal float64
	Count  uint64
	M      sync.Mutex
}

func NewPixel() *Pixel {
	return &Pixel{
		Color:  color.RGBA{0, 0, 0, 255},
		Normal: 0,
		Count:  0,
		M:      sync.Mutex{},
	}
}
