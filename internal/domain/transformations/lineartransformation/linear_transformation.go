package lineartransformation

import (
	"image/color"
	"math/rand/v2"
)

type Affine struct {
	A, B, C, D, E, F float64
	Color            color.RGBA
}

func NewAffine() *Affine {
	a := createCoefficient()
	b := createCoefficient()
	c := createCoefficient()
	d := createCoefficient()
	e := createCoefficient() * 1.5
	f := createCoefficient() * 1.5

	return &Affine{A: a, B: b, C: c, D: d, E: e, F: f, Color: color.RGBA{0, 0, 0, 255}}
}

func createCoefficient() float64 {
	x := rand.Float64()

	if rand.Int()%2 == 0 {
		x *= -1
	}

	return x
}

func (t Affine) Transform(x, y float64) (newX, newY float64) {
	newX = t.A*x + t.B*y + t.C
	newY = t.D*x + t.E*y + t.F

	return newX, newY
}

func (t Affine) GetColor() color.RGBA {
	return t.Color
}
