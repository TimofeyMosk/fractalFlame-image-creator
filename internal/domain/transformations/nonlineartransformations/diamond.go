package nonlineartransformations

import "math"

type Diamond struct {
}

func (Diamond) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	Atan := math.Atan(x / y)
	newX = math.Sin(Atan) * math.Cos(r)
	newY = math.Cos(Atan) * math.Sin(r)

	return newX, newY
}
