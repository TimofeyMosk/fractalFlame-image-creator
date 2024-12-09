package application

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"runtime"
	"strings"
	"time"
)

// NonLinearTransformConfig описывает нелинейное преобразование и вероятность его срабатывания.
type NonLinearTransformConfig struct {
	Name        string  // Название нелинейного преобразования
	Probability float64 // Вероятность применения
}

// Config структура для хранения параметров приложения.
type Config struct {
	Height                    int    // Высота изображения
	Width                     int    // Ширина изображения
	Iterations                uint64 // Количество итераций
	LinearTransformCount      int    // Количество линейных трансформаций
	Symmetry                  bool   // Наличие симметрии
	LogarithmicGamma          bool   // Логарифмическая гамма-коррекция
	Gamma                     float64
	ThreadCount               int                        // Количество потоков
	StretchingCompressionCoef int                        // Коэффициент растяжения (и последующего сжатия) изображения
	NonLinearTransforms       []NonLinearTransformConfig // Нелинейные преобразования с вероятностями
	Filename                  string
}

// ParseFlags парсит флаги из командной строки и возвращает Config.
func ParseFlags() (*Config, error) {
	// Обязательные параметры
	height := flag.Int("height", 0, "Высота изображения (обязательно)")
	width := flag.Int("width", 0, "Ширина изображения (обязательно)")
	iterations := flag.Uint64("iter", 0, "Количество итераций (обязательно)")

	// Необязательные параметры
	linearTransformCount := flag.Int("linear-transform-count", 10, "Количество линейных трансформаций (по умолчанию 10)")
	symmetry := flag.Bool("symmetry", false, "Включить симметрию")
	logGamma := flag.Bool("log-gamma", false, "Включить логарифмическую гамма-коррекцию")
	threadCount := flag.Int("threads", runtime.NumCPU(), "Количество потоков (по умолчанию все доступные)")
	StretchingCompressionCoef := flag.Int("scc", 1, "Коэффициент расстяжения и последующего сжатия изображения(убирает шумы)")
	filename := flag.String("filename", "fractal_image"+"_"+time.Now().Format("D_02_01_2006_T_15_04_05.png"), "Название файла при сохранении")
	// Список нелинейных преобразований
	nonLinearTransforms := flag.String("nonlinear-transforms",
		"",
		"Список нелинейных преобразований с вероятностями, формат: имя:вероятность,...")

	gammaStr := flag.String("gamma", "2.2", "gamma factor float  value")
	flag.Parse()

	// Проверка обязательных параметров
	if *height <= 0 {
		return nil, fmt.Errorf("параметр -height обязателен и должен быть больше 0")
	}

	if *width <= 0 {
		return nil, fmt.Errorf("параметр -width обязателен и должен быть больше 0")
	}

	if *iterations <= 0 {
		return nil, fmt.Errorf("параметр -iterations обязателен и должен быть больше 0")
	}

	// Проверка необязательных параметров
	if *StretchingCompressionCoef <= 0 {
		return nil, fmt.Errorf("параметр -StretchingCompressionCoef должен быть больше 0")
	}

	gamma, err := parseFloat(*gammaStr)
	if err != nil {
		return nil, fmt.Errorf("параметр -StretchingCompressionCoef должен быть вещественным числом")
	}

	if *threadCount <= 0 {
		*threadCount = runtime.NumCPU()
	}

	if *filename == "" {
		return nil, fmt.Errorf("параметр -filename не должен быть пуст")
	}

	// Парсинг нелинейных преобразований
	var transforms []NonLinearTransformConfig

	if *nonLinearTransforms != "" {
		for _, transform := range strings.Split(*nonLinearTransforms, ",") {
			parts := strings.Split(transform, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("некорректный формат нелинейного преобразования: %s", transform)
			}

			probability, err := parseFloat(parts[1])
			if err != nil {
				return nil, fmt.Errorf("некорректная вероятность в преобразовании: %s", transform)
			}

			transforms = append(transforms, NonLinearTransformConfig{
				Name:        strings.ToLower(parts[0]),
				Probability: probability,
			})
		}
	} else {
		nameTransforms := []string{"disk", "handkerchief", "heart",
			"horseshoe", "polar", "sinusoidal", "spherical", "swirl"}
		rand.Shuffle(len(nameTransforms), func(i, j int) {
			nameTransforms[i], nameTransforms[j] = nameTransforms[j], nameTransforms[i]
		})
		countTr := 3
		// Генерация случайных вероятностей, сумма которых равна 1
		probabilities := make([]float64, countTr)
		total := 1.0
		for i := 0; i < 2; i++ {
			probabilities[i] = rand.Float64() * total
			total -= probabilities[i]
		}
		probabilities[2] = total

		for i, _ := range nameTransforms {
			transforms = append(transforms, NonLinearTransformConfig{
				Name:        nameTransforms[i],
				Probability: probabilities[i],
			})
		}

	}

	return &Config{
		Height:                    *height,
		Width:                     *width,
		Iterations:                *iterations,
		LinearTransformCount:      *linearTransformCount,
		Symmetry:                  *symmetry,
		LogarithmicGamma:          *logGamma,
		ThreadCount:               *threadCount,
		NonLinearTransforms:       transforms,
		Gamma:                     gamma,
		StretchingCompressionCoef: *StretchingCompressionCoef,
		Filename:                  *filename,
	}, nil
}

// parseFloat — вспомогательная функция для парсинга float64.
func parseFloat(value string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(value, "%f", &f)

	return f, err
}
