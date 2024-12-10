package nonlineartransformations

import "math"

type Spiral struct {
}

func (Spiral) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	rInv := 1.0 / r
	aTan := math.Atan(x / y)

	newX = rInv * (math.Cos(aTan) + math.Sin(r))
	newY = rInv * (math.Sin(aTan) - math.Cos(r))

	return newX, newY
}
