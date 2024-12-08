package application

import (
	"image/color"
	"math"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

// Используется среднеквадратичное значение яркости для каждого цветового канала,
// что б яркие цвета влияли на конечный результат сильнее чем пустые черные пиксели (с 0/0/0/255)

func CompressionFractalImage(coef int, fractalImg *domain.FractalImage) *domain.FractalImage {
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
					curX, curY := j*coef+x, i*coef+y
					sumSquareRed += uint(fractalImg.Img[curY][curX].Color.R) * uint(fractalImg.Img[curY][curX].Color.R)
					sumSquareGreen += uint(fractalImg.Img[curY][curX].Color.G) * uint(fractalImg.Img[curY][curX].Color.G)
					sumSquareBlue += uint(fractalImg.Img[curY][curX].Color.B) * uint(fractalImg.Img[curY][curX].Color.B)
					sumCount += fractalImg.Img[curY][curX].Count
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

	return newFractalImg
}
