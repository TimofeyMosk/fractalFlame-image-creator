package nonlinear_transformamtions

import "math"

type Handkerchief struct {
}

func (Handkerchief) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	aTan := math.Atan(x / y)
	newX = r * math.Sin(aTan+r)
	newY = r * math.Cos(aTan-r)
	return newX, newY
}
