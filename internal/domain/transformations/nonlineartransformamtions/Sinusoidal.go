package nonlinear_transformamtions

import "math"

type Sinusoidal struct {
	Width, Height int
}

func (t Sinusoidal) Transform(x, y float64) (newX, newY float64) {
	var scaleX, scaleY float64
	if t.Width > t.Height {
		scaleX = float64(t.Width) / float64(t.Height)
		scaleY = 1.0
	} else {
		scaleX = 1.0
		scaleY = float64(t.Height) / float64(t.Width)
	}

	newX = math.Sin(x) * scaleX
	newY = math.Sin(y) * scaleY

	return newX, newY
}
