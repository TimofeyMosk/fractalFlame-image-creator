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
	//	Iterations:           100_000_000,
	//	LinearTransformCount: 10,
	//	Symmetry:             false,
	//	LogarithmicGamma:     true,
	//	ThreadCount:          8,
	//	NonLinearTransforms: []application.NonLinearTransformConfig{
	//		{Name: "sinusoidal", Probability: 0.2},
	//		{Name: "polar", Probability: 0.2},
	//		{Name: "spherical", Probability: 0.3},
	//		{Name: "disk", Probability: 0.3}},
	//	Gamma:                     1.5,
	//	StretchingCompressionCoef: 3,
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

// GenerateBrightColor генерирует яркий цвет
// ./main -height=1080 -width=1920 -iterations=10000000 -linear-transform-count=10 -symmetry -log-gamma -threads=8 -nonlinear-transforms="Sinusoidal:0.2,Polar:0.3,Spherical:0.4"
//./main -height=1080 -width=1920 -iter=500000000 -linear-transform-count=10 -symmetry -threads 8  -log-gamma -nonlinear-transforms="sinusoidal:0.1,polar:0.1,disk:0.4,Handkerchief: 0.4" -scc 5
