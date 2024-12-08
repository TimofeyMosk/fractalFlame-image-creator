package nonlinear_transformamtions

import "math"

type Disk struct {
}

func (Disk) Transform(x, y float64) (newX, newY float64) {
	aTanModPi := math.Atan(x/y) / math.Pi
	rPi := math.Sqrt(x*x+y*y) * math.Pi
	newX = aTanModPi * math.Sin(rPi)
	newY = aTanModPi * math.Cos(rPi)

	return newX, newY
}
