package stats

import (
	"testing"
)

func TestMin(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1.1, 2, 3, 4, 5}, 1.1},
		{[]float64{10.534, 3, 5, 7, 9}, 3.0},
		{[]float64{-5, 1, 5}, -5.0},
		{[]float64{5}, 5},
	} {
		got, err := Min(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		if got != c.out {
			t.Errorf("Min(%.1f) => %.1f != %.1f", c.in, c.out, got)
		}
	}
	_, err := Min([]float64{})
	if err == nil {
		t.Errorf("Empty slice didn't return an error")
	}
}

func BenchmarkMinSmallFloatSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(makeFloatSlice(5))
	}
}

func BenchmarkMinLargeFloatSlice(b *testing.B) {
	lf := makeFloatSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(lf)
	}
}
