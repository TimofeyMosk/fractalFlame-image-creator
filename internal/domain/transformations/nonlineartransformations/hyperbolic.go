package nonlineartransformations

import "math"

type Hyperbolic struct {
}

func (Hyperbolic) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	Atan := math.Atan(x / y)
	newX = math.Sin(Atan) / r
	newY = math.Cos(Atan) * r

	return newX, newY
}
