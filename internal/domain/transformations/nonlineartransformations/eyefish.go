package nonlineartransformations

import "math"

type Eyefish struct {
}

func (Eyefish) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x+y*y) + 1
	r = 2.0 / r

	newX = r * x
	newY = r * y

	return newX, newY
}
