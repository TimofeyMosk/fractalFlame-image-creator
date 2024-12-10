package application

import "github.com/es-debug/backend_academy_2024_project_4-go-TimofeyMosk/internal/domain/transformations/nonlineartransformations"

func initNoLinTransoformation(nonlinConfig []NonLinearTransformConfig, height, width int) []NonLinTransWithProbability {
	arr := []NonLinTransWithProbability{}

	for i := range nonlinConfig {
		lastP := getLastProbability(arr)

		switch nonlinConfig[i].Name {
		case "heart":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Heart{ScaleX: 0.5, ScaleY: 0.3, ShiftUpY: 0.25},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "sinusoidal":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Sinusoidal{Width: width, Height: height},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "polar":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Polar{Width: width, Height: height},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "spherical":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Spherical{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "disk":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Disk{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "swirl":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Swirl{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "horseshoe":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Horseshoe{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "handkerchief":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Handkerchief{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "bubble":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Bubble{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "cosine":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Cosine{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "cross":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Cross{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "diamond":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Diamond{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "exponential":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Exponential{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "eyefish":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Eyefish{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "fisheye":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Fisheye{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "hyperbolic":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Hyperbolic{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "spiral":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Spiral{},
				Probability:    nonlinConfig[i].Probability + lastP})
		case "tangent":
			arr = append(arr, NonLinTransWithProbability{
				Transformation: nonlineartransformations.Tangent{},
				Probability:    nonlinConfig[i].Probability + lastP})

		}
	}

	if len(arr) != 0 && arr[len(arr)-1].Probability != 1.0 {
		arr[len(arr)-1].Probability = 1.0
	}

	return arr
}

func getLastProbability(arr []NonLinTransWithProbability) float64 {
	lastP := 0.0
	if len(arr) != 0 {
		lastP = arr[len(arr)-1].Probability
	}

	return lastP
}
