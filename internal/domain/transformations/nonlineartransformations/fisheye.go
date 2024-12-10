package nonlineartransformations

import "math"

// Отличается от EyeFish порядком swap (newX,newY в конце)

type Fisheye struct {
}

func (Fisheye) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x+y*y) + 1
	r = 2.0 / r

	newX = r * y
	newY = r * x

	return newX, newY
}
