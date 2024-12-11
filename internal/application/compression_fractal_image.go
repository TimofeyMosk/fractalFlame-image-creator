package application

import (
	"image/color"
	"math"
	"sync"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

// Используется среднеквадратичное значение яркости для каждого цветового канала,
// что б яркие цвета влияли на конечный результат сильнее чем пустые черные пиксели (с 0/0/0/255)

func CompressionFractalImage(coef, threads int, fractalImg *domain.FractalImage) *domain.FractalImage {
	newHeight, newWidth := fractalImg.GetHeight()/coef, fractalImg.GetWidth()/coef
	newFractalImg := domain.NewFractalImage(newHeight, newWidth)
	threads = min(threads, newHeight)
	pixelByThread := math.Floor(float64(newHeight) / float64(threads))
	wg := sync.WaitGroup{}

	for i := 0; i < threads-1; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			startHeight, finishHeight := i*int(pixelByThread), i*int(pixelByThread)+int(pixelByThread)
			compressPartImage(startHeight, finishHeight, coef, newFractalImg, fractalImg)
		}()
	}

	compressPartImage((threads-1)*int(pixelByThread), newFractalImg.GetHeight(), coef, newFractalImg, fractalImg)
	wg.Wait()

	return newFractalImg
}

func compressPartImage(heightStart, heightFinish, coef int, newFractal, fractal *domain.FractalImage) {
	for i := heightStart; i < heightFinish; i++ {
		for j := 0; j < newFractal.GetWidth(); j++ {
			var (
				sumSquareRed   uint
				sumSquareGreen uint
				sumSquareBlue  uint
				sumCount       uint64
			)

			for y := 0; y < coef; y++ {
				for x := 0; x < coef; x++ {
					curX, curY := j*coef+x, i*coef+y
					sumSquareRed += uint(fractal.Img[curY][curX].Color.R) * uint(fractal.Img[curY][curX].Color.R)
					sumSquareGreen += uint(fractal.Img[curY][curX].Color.G) * uint(fractal.Img[curY][curX].Color.G)
					sumSquareBlue += uint(fractal.Img[curY][curX].Color.B) * uint(fractal.Img[curY][curX].Color.B)
					sumCount += fractal.Img[curY][curX].Count
				}
			}

			avgRed := uint8(math.Sqrt(float64(sumSquareRed / uint(coef*coef))))
			avgGreen := uint8(math.Sqrt(float64(sumSquareGreen / uint(coef*coef))))
			avgBlue := uint8(math.Sqrt(float64(sumSquareBlue / uint(coef*coef))))
			avgCount := sumCount / uint64(coef*coef)

			newFractal.Img[i][j].Color = color.RGBA{R: avgRed, G: avgGreen, B: avgBlue, A: 255}
			newFractal.Img[i][j].Count = avgCount
		}
	}
}
