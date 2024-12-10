package application

import (
	"math"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

func LogGammaCorrection(fractalImage *domain.FractalImage, gamma float64) {
	maximum := initNormalAndGetMax(fractalImage)

	for i := 0; i < fractalImage.GetHeight(); i++ {
		for j := 0; j < fractalImage.GetWidth(); j++ {
			fractalImage.Img[i][j].Normal /= maximum

			r, g, b := float64(fractalImage.Img[i][j].Color.R)*math.Pow(fractalImage.Img[i][j].Normal, 1.0/gamma),
				float64(fractalImage.Img[i][j].Color.G)*math.Pow(fractalImage.Img[i][j].Normal, 1.0/gamma),
				float64(fractalImage.Img[i][j].Color.B)*math.Pow(fractalImage.Img[i][j].Normal, 1.0/gamma)

			fractalImage.Img[i][j].Color.R = uint8(r)
			fractalImage.Img[i][j].Color.G = uint8(g)
			fractalImage.Img[i][j].Color.B = uint8(b)
		}
	}
}

func initNormalAndGetMax(img *domain.FractalImage) float64 {
	maximum := 0.0

	for i := 0; i < img.GetHeight(); i++ {
		for j := 0; j < img.GetWidth(); j++ {
			if img.Img[i][j].Count == 0 {
				continue
			}

			img.Img[i][j].Normal = math.Log10(float64(img.Img[i][j].Count))
			if img.Img[i][j].Normal > maximum {
				maximum = img.Img[i][j].Normal
			}
		}
	}

	return maximum
}
