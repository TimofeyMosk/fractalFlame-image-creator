package application

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"runtime"
	"strings"
	"time"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
)

// ParseFlags парсит флаги из командной строки и возвращает Config.
func ParseFlags() (*domain.Config, error) {
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
	nonLinearTransforms := flag.String("nonlinear-transforms",
		"",
		"Список нелинейных преобразований с вероятностями, формат: имя:вероятность,...")
	gammaStr := flag.String("gamma", "2.2", "gamma factor float  value")

	flag.Parse()

	if *height <= 0 {
		return nil, fmt.Errorf("параметр -height обязателен и должен быть больше 0")
	}

	if *width <= 0 {
		return nil, fmt.Errorf("параметр -width обязателен и должен быть больше 0")
	}

	if *iterations <= 0 {
		return nil, fmt.Errorf("параметр -iter обязателен и должен быть больше 0")
	}

	// Проверка необязательных параметров
	if *StretchingCompressionCoef <= 0 {
		return nil, fmt.Errorf("параметр -scc должен быть больше 0")
	}

	gamma, err := parseFloat(*gammaStr)
	if err != nil {
		return nil, fmt.Errorf("параметр -gamma должен быть вещественным числом")
	}

	if *threadCount <= 0 {
		*threadCount = runtime.NumCPU()
	}

	if *filename == "" {
		return nil, fmt.Errorf("параметр -filename не должен быть пуст")
	}

	transforms, err := ParseNonLinearTransformations(nonLinearTransforms)
	if err != nil {
		return nil, err
	}

	return &domain.Config{
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

func parseFloat(value string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(value, "%f", &f)

	return f, err
}

func ParseNonLinearTransformations(nonLinearTransforms *string) ([]domain.NonLinearTransformConfig, error) {
	var transforms []domain.NonLinearTransformConfig

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

			transforms = append(transforms, domain.NonLinearTransformConfig{
				Name:        strings.ToLower(parts[0]),
				Probability: probability,
			})
		}
	} else {
		nameTransforms := []string{"bubble", "cosine", "cross", "diamond", "exponential", "eyefish",
			"fisheye", "disk", "handkerchief", "heart", "hyperbolic", "spiral", "tangent",
			"horseshoe", "polar", "sinusoidal", "spherical", "swirl"}
		rand.Shuffle(len(nameTransforms), func(i, j int) {
			nameTransforms[i], nameTransforms[j] = nameTransforms[j], nameTransforms[i]
		})

		// Генерация случайных вероятностей, сумма которых равна 1
		countTr := 4
		probabilities := make([]float64, countTr)
		total := 1.0

		for i := 0; i < countTr-1; i++ {
			probabilities[i] = rand.Float64() * total //nolint // No need to use cryptographic randomiser(it is slower)
			total -= probabilities[i]
		}

		probabilities[countTr-1] = total

		for i := 0; i < countTr; i++ {
			transforms = append(transforms, domain.NonLinearTransformConfig{
				Name:        nameTransforms[i],
				Probability: probabilities[i],
			})
		}
	}

	sumP := 0.0

	for _, transform := range transforms {
		sumP += transform.Probability
	}

	if sumP > 1.01 {
		return nil, fmt.Errorf("cумма вероятностей не может быть больше 1")
	}

	return transforms, nil
}
