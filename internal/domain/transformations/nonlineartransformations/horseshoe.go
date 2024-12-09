package nonlineartransformations

import "math"

type Horseshoe struct {
}

func (Horseshoe) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	newX = (1 / r) * ((x - y) * (x + y))
	newY = (1 / r) * 2 * x * y

	return newX, newY
}
