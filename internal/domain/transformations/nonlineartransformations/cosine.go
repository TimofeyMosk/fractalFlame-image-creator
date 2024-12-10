package nonlineartransformations

import "math"

type Cosine struct {
}

func (Cosine) Transform(x, y float64) (newX, newY float64) {
	newX = math.Cos(math.Pi*x) * math.Cosh(y) * 0.4
	newY = -1 * math.Sin(math.Pi*x) * math.Sinh(y) * 0.4

	return newX, newY
}
