package application

import (
	"math/rand/v2"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/pkg"
)

//nolint:gosec // No need to use cryptographic randomiser(it is slower)
func GenerateBrightColor() (r, g, b uint8) {
	// Генерация случайного H (оттенка) от 0 до 360
	h := rand.Float64() * 360

	// Устанавливаем насыщенность (S) и яркость (V) от 0.9 до 1.0
	s := 0.9 + rand.Float64()*0.1
	v := 0.9 + rand.Float64()*0.1

	// Конвертируем HSV в RGB
	return pkg.HSVToRGB(h, s, v)
}
