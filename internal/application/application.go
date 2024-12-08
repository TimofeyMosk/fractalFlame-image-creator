package application

import (
	"image/color"
	"math/rand/v2"
	"runtime"
	"sync"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain/transformations/lineartransformation"
	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain/transformations/nonlineartransformations"
)

type NonLinearTransoformation interface {
	Transform(x, y float64) (newX, newY float64)
}

type LinearTransformation interface {
	Transform(x, y float64) (newX, newY float64)
	GetColor() color.RGBA
}

type NonLinTransWithProbability struct {
	Transformation NonLinearTransoformation
	Probability    float64
}

type FractalFlameImageGenerator struct {
	fractal                   *domain.FractalImage
	LinTransf                 []LinearTransformation
	NoLinTransf               []NonLinTransWithProbability
	Iteration                 uint64
	logGammaCorrection        bool
	symmetry                  bool
	gamma                     float64
	coefStretchingCompression int
	threadCount               int
}

func NewFractalFlameImageGenerator(cfg *Config) *FractalFlameImageGenerator {
	return &FractalFlameImageGenerator{
		fractal:                   domain.NewFractalImage(cfg.Height*cfg.StretchingCompressionCoef, cfg.Width*cfg.StretchingCompressionCoef),
		LinTransf:                 initLinTransform(cfg.LinearTransformCount),
		NoLinTransf:               initNoLinTransoformation(cfg.NonLinearTransforms, cfg.Height, cfg.Width),
		Iteration:                 cfg.Iterations,
		gamma:                     cfg.Gamma,
		coefStretchingCompression: cfg.StretchingCompressionCoef,
		threadCount:               cfg.ThreadCount,
		symmetry:                  cfg.Symmetry,
		logGammaCorrection:        cfg.LogarithmicGamma,
	}
}

func (f *FractalFlameImageGenerator) Start() *domain.FractalImage {
	runtime.GOMAXPROCS(f.threadCount)
	iterationsByGorutine := f.Iteration / uint64(f.threadCount)

	wg := &sync.WaitGroup{}
	for i := 0; i < f.threadCount; i++ {
		wg.Add(1)

		go func() {
			Render(f, iterationsByGorutine)
			wg.Done()
		}()
	}

	wg.Wait()

	if f.logGammaCorrection {
		LogGammaCorrection(f.fractal, f.gamma)
	}

	if f.coefStretchingCompression > 1 {
		f.fractal = CompressionFractalImage(f.coefStretchingCompression, f.fractal)
	}

	return f.fractal
}

func initNoLinTransoformation(nonlinConfig []NonLinearTransformConfig, height, width int) []NonLinTransWithProbability {
	arr := []NonLinTransWithProbability{}

	for i := range nonlinConfig {
		switch nonlinConfig[i].Name {
		case "heart":
			lastP := 0.0
			if len(arr) != 0 {
				lastP = arr[len(arr)-1].Probability
			}

			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Heart{ScaleX: 0.5, ScaleY: 0.3, ShiftUpY: 0.25},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "sinusoidal":
			lastP := 0.0
			if len(arr) != 0 {
				lastP = arr[len(arr)-1].Probability
			}

			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Sinusoidal{Width: width, Height: height},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "polar":
			lastP := 0.0
			if len(arr) != 0 {
				lastP = arr[len(arr)-1].Probability
			}

			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Polar{Width: width, Height: height},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "spherical":
			lastP := 0.0
			if len(arr) != 0 {
				lastP = arr[len(arr)-1].Probability
			}

			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Spherical{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "disk":
			lastP := 0.0
			if len(arr) != 0 {
				lastP = arr[len(arr)-1].Probability
			}

			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Disk{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "swirl":
			lastP := 0.0
			if len(arr) != 0 {
				lastP = arr[len(arr)-1].Probability
			}

			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Swirl{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "horseshoe":
			lastP := 0.0
			if len(arr) != 0 {
				lastP = arr[len(arr)-1].Probability
			}

			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Horseshoe{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "handkerchief":
			lastP := 0.0
			if len(arr) != 0 {
				lastP = arr[len(arr)-1].Probability
			}

			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Handkerchief{},
				Probability:    nonlinConfig[i].Probability + lastP})
		}
	}

	if len(arr) != 0 && arr[len(arr)-1].Probability != 1.0 {
		arr[len(arr)-1].Probability = 1.0
	}

	return arr
}

func initLinTransform(countTr int) []LinearTransformation {
	curCount := 0
	result := make([]LinearTransformation, 0, countTr)

	for curCount < countTr {
		a := createCoefficient()
		b := createCoefficient()
		c := createCoefficient()
		d := createCoefficient()
		e := createCoefficient() * 1.7
		f := createCoefficient() * 1.7

		if (a*a+d*d < 1) && (b*b+e*e < 1) && (a*a+b*b+d*d+e*e < 1+(a*e-b*d)*(a*e-b*d)) {
			red, green, blue := GenerateBrightColor()
			col := color.RGBA{
				R: red,
				G: green,
				B: blue,
				A: 255,
			}
			result = append(result, lineartransformation.Affine{A: a, B: b, C: c, D: d, E: e, F: f, Color: col})
			curCount++
		}
	}

	return result
}

func createCoefficient() float64 {
	x := rand.Float64()

	if rand.Int()%2 == 0 {
		x *= -1
	}

	return x
}
