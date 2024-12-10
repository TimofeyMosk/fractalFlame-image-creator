package nonlineartransformations

import "math"

type Cross struct {
}

func (Cross) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(((x*x - y*y) * (x*x - y*y)))
	newX = r * x * 0.2
	newY = r * y * 0.2

	return newX, newY
}
