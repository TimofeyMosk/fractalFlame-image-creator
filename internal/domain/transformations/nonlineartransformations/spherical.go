package nonlineartransformations

type Spherical struct {
}

func (Spherical) Transform(x, y float64) (newX, newY float64) {
	temp := 1 / (x*x + y*y)
	newX = x * temp
	newY = y * temp

	return newX, newY
}
