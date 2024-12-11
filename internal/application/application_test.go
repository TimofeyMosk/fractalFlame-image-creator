package application_test

import (
	"strconv"
	"testing"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/application"
	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

// Бенчмарк для FractalFlameImageGenerator с разным количеством горутин.
func BenchmarkFractalFlameImageGenerator_Start(b *testing.B) {
	// Конфигурация для генератора фракталов
	cfg := &domain.Config{
		Height:               1080,
		Width:                1920,
		LinearTransformCount: 100,
		NonLinearTransforms: []domain.NonLinearTransformConfig{
			{Name: "sinusoidal", Probability: 0.3},
			{Name: "polar", Probability: 0.1},
			{Name: "spherical", Probability: 0.1},
			{Name: "disk", Probability: 0.1},
			{Name: "heart", Probability: 0.4},
		},
		Iterations:                2000000,
		Gamma:                     2.2,
		StretchingCompressionCoef: 4,
		ThreadCount:               1, // Это будет изменяться
		Symmetry:                  false,
		LogarithmicGamma:          true,
	}

	threadCounts := []int{1, 2, 4, 8, 12, 16, 64, 128, 1280} // Разные значения горутин для тестирования

	for _, threads := range threadCounts {
		b.Run("ThreadCount="+strconv.Itoa(threads), func(b *testing.B) {
			cfg.ThreadCount = threads // Устанавливаем количество горутин
			generator := application.NewFractalFlameImageGenerator(cfg)

			b.ResetTimer() // Сбрасываем таймер перед запуском теста

			for i := 0; i < b.N; i++ {
				generator.Start() // Запуск функции
			}
		})
	}
}
