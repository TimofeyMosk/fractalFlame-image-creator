package main

import (
	"fmt"
	"time"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/application"
	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/infrastructure"
)

func main() {
	startApplication := time.Now()

	cfg, err := application.ParseFlags()
	if err != nil {
		fmt.Errorf("Ошибка парсинга флагов: %w", err)
		return
	}

	fmt.Printf("Конфигурация: %+v\n", cfg)

	//cfg := &application.Config{
	//	Height:               1080,
	//	Width:                1920,
	//	Iterations:           100_000_00,
	//	LinearTransformCount: 10,
	//	Symmetry:             false,
	//	LogarithmicGamma:     false,
	//	ThreadCount:          8,
	//	NonLinearTransforms:  []application.NonLinearTransformConfig{
	//		//{Name: "sinusoidal", Probability: 0.2},
	//		//{Name: "polar", Probability: 0.3},
	//		//{Name: "tangent", Probability: 1.0},
	//		//{Name: "disk", Probability: 0.3},
	//		//{Name: "heart", Probability: 0.4},
	//	},
	//	Gamma:                     2.2,
	//	StretchingCompressionCoef: 1,
	//	Filename:                  "Fractal.png",
	//}

	fractalGenerator := application.NewFractalFlameImageGenerator(cfg)
	fractalImage := fractalGenerator.Start()

	err = infrastructure.SaveImage(cfg.Filename, fractalImage)
	if err != nil {
		fmt.Errorf("Don`t save fractal.png: %v", err)
		return
	}

	fmt.Println(time.Since(startApplication).Seconds())
}
