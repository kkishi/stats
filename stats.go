package stats

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
