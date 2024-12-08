package lineartransformation

import "image/color"

type Affine struct {
	A, B, C, D, E, F float64
	Color            color.RGBA
}

func (t Affine) Transform(x, y float64) (newX, newY float64) {
	newX = t.A*x + t.B*y + t.C
	newY = t.D*x + t.E*y + t.F

	return newX, newY
}

func (t Affine) GetColor() color.RGBA {
	return t.Color
}
