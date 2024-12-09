package nonlineartransformations

import "math"

type Polar struct {
	Width, Height int
}

func (t Polar) Transform(x, y float64) (newX, newY float64) {
	scaleX := 1.0
	if t.Width > t.Height {
		scaleX = float64(t.Width) / float64(t.Height)
	}

	newX = (math.Atan(x/y) / math.Pi) * scaleX * 2
	newY = math.Sqrt(x*x+y*y) - 1

	return newX, newY
}
