package application

import (
	"image/color"
	"math"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

func CompressionFractalImage(coef int, fractalImg *domain.FractalImage) domain.FractalImage {
	newHeight, newWidth := fractalImg.GetHeight()/coef, fractalImg.GetWidth()/coef
	newFractalImg := domain.NewFractalImage(newHeight, newWidth)

	for i := 0; i < newHeight; i++ {
		for j := 0; j < newWidth; j++ {
			var (
				sumSquareRed   uint
				sumSquareGreen uint
				sumSquareBlue  uint
				sumCount       uint64
			)

			for y := 0; y < coef; y++ {
				for x := 0; x < coef; x++ {
					sumSquareRed += uint(fractalImg.Img[i*coef+y][j*coef+x].Color.R) * uint(fractalImg.Img[i*coef+y][j*coef+x].Color.R)
					sumSquareGreen += uint(fractalImg.Img[i*coef+y][j*coef+x].Color.G) * uint(fractalImg.Img[i*coef+y][j*coef+x].Color.G)
					sumSquareBlue += uint(fractalImg.Img[i*coef+y][j*coef+x].Color.B) * uint(fractalImg.Img[i*coef+y][j*coef+x].Color.B)
					sumCount += fractalImg.Img[i*coef+y][j*coef+x].Count
				}
			}

			avgRed := uint8(math.Sqrt(float64(sumSquareRed / uint(coef*coef))))
			avgGreen := uint8(math.Sqrt(float64(sumSquareGreen / uint(coef*coef))))
			avgBlue := uint8(math.Sqrt(float64(sumSquareBlue / uint(coef*coef))))
			avgCount := sumCount / uint64(coef*coef)

			newFractalImg.Img[i][j].Color = color.RGBA{R: avgRed, G: avgGreen, B: avgBlue, A: 255}
			newFractalImg.Img[i][j].Count = avgCount
		}
	}

	return *newFractalImg
}
