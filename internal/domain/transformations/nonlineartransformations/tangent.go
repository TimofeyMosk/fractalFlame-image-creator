package nonlineartransformations

import "math"

type Tangent struct {
}

func (Tangent) Transform(x, y float64) (newX, newY float64) {
	newX = math.Sin(x) / math.Cos(y)
	newY = math.Tan(y)

	return newX, newY
}
