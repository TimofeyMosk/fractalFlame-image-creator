package infrastructure

import (
	"image"
	"image/png"
	"os"
)

func SaveImage(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	// Кодируем изображение в формат PNG
	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
