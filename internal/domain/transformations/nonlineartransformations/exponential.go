package nonlineartransformations

import "math"

type Exponential struct {
}

func (Exponential) Transform(x, y float64) (newX, newY float64) {
	exp := math.Exp(x - 1)
	newX = exp * math.Cos(math.Pi*y)
	newY = exp * math.Sin(math.Pi*y)

	return newX, newY
}
