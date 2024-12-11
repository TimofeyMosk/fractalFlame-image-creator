package application

import (
	"math"
	"sync"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

func MultiThreadLogGamma(threads int, fractalImage *domain.FractalImage, gamma float64) {
	threads = min(threads, fractalImage.GetHeight())
	pixelByThread := math.Floor(float64(fractalImage.GetHeight()) / float64(threads))
	wg := sync.WaitGroup{}
	maximums := make(chan float64, threads)

	for i := 0; i < threads-1; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			startHeight, finishHeight := i*int(pixelByThread), i*int(pixelByThread)+int(pixelByThread)
			maximums <- initNormalAndGetMax(startHeight, finishHeight, fractalImage)
		}()
	}

	maximums <- initNormalAndGetMax((threads-1)*int(pixelByThread), fractalImage.GetHeight(), fractalImage)

	wg.Wait()
	close(maximums)

	maximum := 0.0
	for value := range maximums {
		if value > maximum {
			maximum = value
		}
	}

	for i := 0; i < threads-1; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			startHeight, finishHeight := i*int(pixelByThread), i*int(pixelByThread)+int(pixelByThread)
			LogGammaCorrection(startHeight, finishHeight, fractalImage, maximum, gamma)
		}()
	}

	LogGammaCorrection((threads-1)*int(pixelByThread), fractalImage.GetHeight(), fractalImage, maximum, gamma)
	wg.Wait()
}

func LogGammaCorrection(startHeight, finishHeight int, fractalImage *domain.FractalImage, maximum, gamma float64) {
	for i := startHeight; i < finishHeight; i++ {
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

func initNormalAndGetMax(startHeight, finishHeight int, img *domain.FractalImage) float64 {
	maximum := 0.0

	for i := startHeight; i < finishHeight; i++ {
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
