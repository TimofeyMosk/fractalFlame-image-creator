package application

import (
	"math/rand/v2"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

func Render(img *domain.FractalImage, linTrasforms []LinearTransformation, nonLinearTranforms []NonLinTransWithProbability, iterations uint64) {
	x, y := rand.Float64(), rand.Float64()
	if rand.Int()%2 == 0 {
		x *= -1
	}
	if rand.Int()%2 == 0 {
		y *= -1
	}

	var xMax, xMin, yMax, yMin float64

	if img.GetWidth() > img.GetHeight() {
		xMax = float64(img.GetWidth()) / float64(img.GetHeight())
		xMin = -1.0 * xMax
		yMax = 1.0
		yMin = -1.0
	} else {
		xMax = 1.0
		xMin = -1.0
		yMax = float64(img.GetWidth()) / float64(img.GetHeight())
		yMin = -1.0 * yMax
	}

	var i uint64
	for i = 0; i < iterations; i++ {
		linT := linTrasforms[rand.IntN(len(linTrasforms))]
		// Аффинное преобразование
		newX, newY := linT.Transform(x, y)

		p := rand.Float64()
		ch := 0

		for index := range nonLinearTranforms {
			if p < nonLinearTranforms[index].Probability {
				ch = index
				break
			}
		}

		trans := nonLinearTranforms[ch].Transformation
		x, y = trans.Transform(newX, newY)

		// Преобразование координат в пространство изображения
		if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
			imgX := int((x - xMin) / (xMax - xMin) * float64(img.GetWidth()))
			imgY := int((y - yMin) / (yMax - yMin) * float64(img.GetHeight()))

			if imgX >= 0 && imgX < img.GetWidth() && imgY >= 0 && imgY < img.GetHeight() {
				// Увеличение яркости пикселя
				img.Img[imgY][imgX].M.Lock()
				if img.Img[imgY][imgX].Count != 0 {
					original := img.Img[imgY][imgX].Color
					transformationColor := linT.GetColor()
					img.Img[imgY][imgX].Color.R = uint8((uint16(original.R) + uint16(transformationColor.R) + 1) >> 1)
					img.Img[imgY][imgX].Color.G = uint8((uint16(original.G) + uint16(transformationColor.G) + 1) >> 1)
					img.Img[imgY][imgX].Color.B = uint8((uint16(original.B) + uint16(transformationColor.B) + 1) >> 1)
					img.Img[imgY][imgX].Count++
				} else {
					transformationColor := linT.GetColor()
					img.Img[imgY][imgX].Color.R = transformationColor.R
					img.Img[imgY][imgX].Color.G = transformationColor.G
					img.Img[imgY][imgX].Color.B = transformationColor.B
					img.Img[imgY][imgX].Count++
				}
				img.Img[imgY][imgX].M.Unlock()
			}
		}
	}
}
