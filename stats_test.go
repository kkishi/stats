package stats

import (
	"math"
	"testing"
)

func close(a, b float64) bool { return math.Abs(a-b) < 1e-14 }

func TestSum(t *testing.T) {
	tests := []struct {
		xs   []float64
		want float64
	}{
		{
			want: 0,
		},
		{
			xs:   []float64{1, 2, 3},
			want: 6,
		},
	}
	for _, test := range tests {
		if got := Sum(test.xs); !close(got, test.want) {
			t.Errorf("Sum(%v): got %f, want %f", test.xs, got, test.want)
		}
	}
}

func TestArithmeticMean(t *testing.T) {
	tests := []struct {
		xs   []float64
		want float64
	}{
		{
			want: 0,
		},
		{
			xs:   []float64{1, 2, 3},
			want: 2,
		},
	}
	for _, test := range tests {
		if got := ArithmeticMean(test.xs); !close(got, test.want) {
			t.Errorf("ArithmeticMean(%v): got %f, want %f", test.xs, got, test.want)
		}
	}
}
