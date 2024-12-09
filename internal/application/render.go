package application

import (
	rand2 "math/rand"
	"math/rand/v2"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

func Render(ffg *FractalFlameImageGenerator, iterations uint64) {
	x, y := initStartPoint()
	xMax, xMin, yMax, yMin := initScreenRatio(ffg.fractal)

	var i uint64
	for i = 0; i < iterations; i++ {
		linT := ffg.linTransf[rand.IntN(len(ffg.linTransf))]
		newX, newY := linT.Transform(x, y)
		trans := choiceTransform(ffg.noLinTransf)

		if trans != nil {
			x, y = trans.Transform(newX, newY)
		} else {
			x, y = newX, newY
		}

		if ffg.symmetry {
			if rand2.Int()%2 == 0 {
				x *= -1
				y *= -1
			}
		}

		// Преобразование координат в пространство изображения
		if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
			imgX := int((x - xMin) / (xMax - xMin) * float64(ffg.fractal.GetWidth()))
			imgY := int((y - yMin) / (yMax - yMin) * float64(ffg.fractal.GetHeight()))

			if imgX >= 0 && imgX < ffg.fractal.GetWidth() && imgY >= 0 && imgY < ffg.fractal.GetHeight() {
				ffg.fractal.Img[imgY][imgX].M.Lock()
				if ffg.fractal.Img[imgY][imgX].Count != 0 {
					original := ffg.fractal.Img[imgY][imgX].Color
					transformationColor := linT.GetColor()
					ffg.fractal.Img[imgY][imgX].Color.R = uint8((uint16(original.R) + uint16(transformationColor.R) + 1) >> 1)
					ffg.fractal.Img[imgY][imgX].Color.G = uint8((uint16(original.G) + uint16(transformationColor.G) + 1) >> 1)
					ffg.fractal.Img[imgY][imgX].Color.B = uint8((uint16(original.B) + uint16(transformationColor.B) + 1) >> 1)
					ffg.fractal.Img[imgY][imgX].Count++
				} else {
					transformationColor := linT.GetColor()
					ffg.fractal.Img[imgY][imgX].Color.R = transformationColor.R
					ffg.fractal.Img[imgY][imgX].Color.G = transformationColor.G
					ffg.fractal.Img[imgY][imgX].Color.B = transformationColor.B
					ffg.fractal.Img[imgY][imgX].Count++
				}
				ffg.fractal.Img[imgY][imgX].M.Unlock()
			}
		}
	}
}

func initStartPoint() (x, y float64) {
	x, y = rand.Float64(), rand.Float64()

	if rand.Int()%2 == 0 {
		x *= -1
	}

	if rand.Int()%2 == 0 {
		y *= -1
	}

	return x, y
}

func initScreenRatio(image *domain.FractalImage) (xMax, xMin, yMax, yMin float64) {
	if image.GetWidth() > image.GetHeight() {
		xMax = float64(image.GetWidth()) / float64(image.GetHeight())
		xMin = -1.0 * xMax
		yMax = 1.0
		yMin = -1.0
	} else {
		xMax = 1.0
		xMin = -1.0
		yMax = float64(image.GetWidth()) / float64(image.GetHeight())
		yMin = -1.0 * yMax
	}

	return xMax, xMin, yMax, yMin
}

func choiceTransform(arrTr []NonLinTransWithProbability) NonLinearTransoformation {
	if len(arrTr) == 0 {
		return nil
	}
	p := rand.Float64()
	ch := 0

	for index := range arrTr {
		if p < arrTr[index].Probability {
			ch = index
			break
		}
	}

	return arrTr[ch].Transformation
}
