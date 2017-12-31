package stats

import "math"

func Sum(xs []float64) float64 {
	var f float64
	for _, x := range xs {
		f += x
	}
	return f
}

func ArithmeticMean(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	return Sum(xs) / float64(len(xs))
}

func square(x float64) float64 {
	return x * x
}

func SquaredDeviationSum(xs []float64) float64 {
	m := ArithmeticMean(xs)
	var s float64
	for _, x := range xs {
		s += square(x - m)
	}
	return s
}

func Variance(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	return SquaredDeviationSum(xs) / float64(len(xs))
}

func StandardDeviation(xs []float64) float64 {
	return math.Sqrt(Variance(xs))
}

func CoefficientOfVariation(xs []float64) float64 {
	m := ArithmeticMean(xs)
	if m == 0 {
		return 0
	}
	return StandardDeviation(xs) / m
}

func CorrelationCoefficient(xs []float64, ys []float64) float64 {
	xm, ym := ArithmeticMean(xs), ArithmeticMean(ys)
	var s float64
	for i, x := range xs {
		s += (x - xm) * (ys[i] - ym)
	}
	return s / math.Sqrt(SquaredDeviationSum(xs)*SquaredDeviationSum(ys))
}
