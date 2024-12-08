package application

import (
	"math"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

func LogGammaCorrection(img *domain.FractalImage, gamma float64) {
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

	for i := 0; i < img.GetHeight(); i++ {
		for j := 0; j < img.GetWidth(); j++ {
			img.Img[i][j].Normal /= maximum

			r, g, b := float64(img.Img[i][j].Color.R)*math.Pow(img.Img[i][j].Normal, 1.0/gamma),
				float64(img.Img[i][j].Color.G)*math.Pow(img.Img[i][j].Normal, 1.0/gamma),
				float64(img.Img[i][j].Color.B)*math.Pow(img.Img[i][j].Normal, 1.0/gamma)

			img.Img[i][j].Color.R = uint8(r)
			img.Img[i][j].Color.G = uint8(g)
			img.Img[i][j].Color.B = uint8(b)
		}
	}
}
