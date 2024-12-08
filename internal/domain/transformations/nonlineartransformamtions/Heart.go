package nonlinear_transformamtions

import "math"

type Heart struct {
	ScaleX, ScaleY float64
	ShiftUpY       float64
}

func (t Heart) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	aTan := math.Atan(y / x)
	newX = r * math.Sin(aTan*r)
	newY = -r * math.Cos(aTan*r)
	newX *= t.ScaleX
	newY *= t.ScaleY
	newY -= t.ShiftUpY
	return newX, newY
}
