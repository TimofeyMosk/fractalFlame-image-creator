package application

import (
	"image/color"
	"runtime"
	"sync"

	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain"
	"github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain/transformations/lineartransformation"
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
	linTransf                 []LinearTransformation
	noLinTransf               []NonLinTransWithProbability
	iteration                 uint64
	logGammaCorrection        bool
	symmetry                  bool
	gamma                     float64
	coefStretchingCompression int
	threadCount               int
}

func NewFractalFlameImageGenerator(cfg *domain.Config) *FractalFlameImageGenerator {
	return &FractalFlameImageGenerator{
		fractal:                   domain.NewFractalImage(cfg.Height*cfg.StretchingCompressionCoef, cfg.Width*cfg.StretchingCompressionCoef),
		linTransf:                 initLinTransform(cfg.LinearTransformCount),
		noLinTransf:               initNoLinTransoformation(cfg.NonLinearTransforms, cfg.Height, cfg.Width),
		iteration:                 cfg.Iterations,
		gamma:                     cfg.Gamma,
		coefStretchingCompression: cfg.StretchingCompressionCoef,
		threadCount:               cfg.ThreadCount,
		symmetry:                  cfg.Symmetry,
		logGammaCorrection:        cfg.LogarithmicGamma,
	}
}

func (f *FractalFlameImageGenerator) Start() *domain.FractalImage {
	runtime.GOMAXPROCS(f.threadCount)
	iterationsByGorutine := f.iteration / uint64(f.threadCount)

	wg := &sync.WaitGroup{}
	for i := 0; i < f.threadCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			Render(f, iterationsByGorutine)
		}()
	}

	wg.Wait()

	if f.logGammaCorrection {
		MultiThreadLogGamma(f.threadCount, f.fractal, f.gamma)
	}

	if f.coefStretchingCompression > 1 {
		f.fractal = CompressionFractalImage(f.coefStretchingCompression, f.threadCount, f.fractal)
	}

	return f.fractal
}

func initLinTransform(countTr int) []LinearTransformation {
	curCount := 0
	result := make([]LinearTransformation, 0, countTr)

	for curCount < countTr {
		affineTrans := lineartransformation.NewAffine()

		if validAffine(affineTrans) {
			red, green, blue := GenerateBrightColor()
			affineTrans.Color = color.RGBA{
				R: red,
				G: green,
				B: blue,
				A: 255,
			}

			result = append(result, affineTrans)
			curCount++
		}
	}

	return result
}

func validAffine(t *lineartransformation.Affine) bool {
	if (t.A*t.A+t.D*t.D < 1) && (t.B*t.B+t.E*t.E < 1) &&
		(t.A*t.A+t.B*t.B+t.D*t.D+t.E*t.E < 1+(t.A*t.E-t.B*t.D)*(t.A*t.E-t.B*t.D)) {
		return true
	}

	return false
}
