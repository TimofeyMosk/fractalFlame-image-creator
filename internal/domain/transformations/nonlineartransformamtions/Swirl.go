package nonlineartransformamtions

import "math"

type Swirl struct {
}

func (Swirl) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	newX = x*math.Sin(r) - y*math.Cos(r)
	newY = x*math.Cos(r) + y*math.Sin(r)
	return newX, newY
}
