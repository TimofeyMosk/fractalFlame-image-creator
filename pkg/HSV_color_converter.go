package pkg

func HSVToRGB(h, s, v float64) (r, g, b uint8) {
	c := v * s
	x := c * (1 - absMod(h/60.0, 2) - 1)
	m := v - c

	var r1, g1, b1 float64

	switch {
	case h >= 0 && h < 60:
		r1, g1, b1 = c, x, 0
	case h >= 60 && h < 120:
		r1, g1, b1 = x, c, 0
	case h >= 120 && h < 180:
		r1, g1, b1 = 0, c, x
	case h >= 180 && h < 240:
		r1, g1, b1 = 0, x, c
	case h >= 240 && h < 300:
		r1, g1, b1 = x, 0, c
	case h >= 300 && h < 360:
		r1, g1, b1 = c, 0, x
	}

	r = uint8((r1 + m) * 255)
	g = uint8((g1 + m) * 255)
	b = uint8((b1 + m) * 255)

	return
}

// absMod вычисляет положительный остаток от деления.
func absMod(a, b float64) float64 {
	mod := a - b*float64(int(a/b))
	if mod < 0 {
		return mod + b
	}

	return mod
}
