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
		fmt.Printf("Ошибка парсинга флагов: %v", err)
		return
	}

	fmt.Printf("Конфигурация: %s\n", cfg)

	fractalGenerator := application.NewFractalFlameImageGenerator(cfg)
	fractalImage := fractalGenerator.Start()

	err = infrastructure.SaveImage(cfg.Filename, fractalImage)
	if err != nil {
		fmt.Printf("Don`t save fractal.png: %v", err)
		return
	}

	fmt.Printf("Generation time: %.3fs\n", time.Since(startApplication).Seconds())
}
