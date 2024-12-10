package nonlineartransformations

import "math"

type Bubble struct {
}

func (Bubble) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	r *= r
	r += 4
	r = 4.0 / r

	newX = r * x
	newY = r * y

	return newX, newY
}
