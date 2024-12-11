package domain

import "fmt"

// NonLinearTransformConfig описывает нелинейное преобразование и вероятность его срабатывания.
type NonLinearTransformConfig struct {
	Name        string  // Название нелинейного преобразования
	Probability float64 // Вероятность применения
}

// Config структура для хранения параметров приложения.
type Config struct {
	Height                    int                        // Высота изображения
	Width                     int                        // Ширина изображения
	Iterations                uint64                     // Количество итераций
	LinearTransformCount      int                        // Количество линейных трансформаций
	Symmetry                  bool                       // Наличие симметрии
	LogarithmicGamma          bool                       // Логарифмическая гамма-коррекция
	Gamma                     float64                    // Параметр гамма для Логарифмической гамма-коррекции
	ThreadCount               int                        // Количество потоков
	StretchingCompressionCoef int                        // Коэффициент растяжения (и последующего сжатия) изображения
	NonLinearTransforms       []NonLinearTransformConfig // Нелинейные преобразования с вероятностями
	Filename                  string
}

func (c Config) String() string {
	res := fmt.Sprintf("Height: %v\n", c.Height)
	res += fmt.Sprintf("Width: %v\n", c.Width)
	res += fmt.Sprintf("Iterations: %v\n", c.Iterations)
	res += fmt.Sprintf("LinearTransformCount: %v\n", c.LinearTransformCount)
	res += fmt.Sprintf("Symmetry: %v\n", c.Symmetry)
	res += fmt.Sprintf("Logarithm gamma: %v\n", c.LogarithmicGamma)

	if c.LogarithmicGamma {
		res += fmt.Sprintf("Gamma: %v\n", c.Gamma)
	}

	res += fmt.Sprintf("Threads: %v\n", c.ThreadCount)
	res += fmt.Sprintf("Stretch compression coef: %v\n", c.StretchingCompressionCoef)

	for i := range c.NonLinearTransforms {
		res += fmt.Sprintf("Name: %v, Probabilities; %.3f \n", c.NonLinearTransforms[i].Name, c.NonLinearTransforms[i].Probability)
	}

	res += fmt.Sprintf("Filename: %v\n", c.Filename)

	return res
}
