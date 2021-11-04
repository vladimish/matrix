package mat64

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestEqual(t *testing.T) {
	tests := map[string]struct {
		a   [][]float64
		b   [][]float64
		exp bool
	}{
		"big_float": {
			[][]float64{{1, 2.3333333333, 3, 4, math.MaxFloat64}, {5, math.Inf(1), 3, 2, 1}},
			[][]float64{{1, 2.3333333333, 3, 4, math.MaxFloat64}, {5, math.Inf(1), 3, 2, 1}},
			true,
		},
		"reversed": {
			[][]float64{{5, 4, 3, 2, 1}},
			[][]float64{{1, 2, 3, 4, 5}},
			false,
		},
	}

	for i := range tests {
		got := Equal(tests[i].a, tests[i].b)
		if got != tests[i].exp {
			t.Errorf("Case %s: expected %v, got: %v", i, tests[i].exp, got)
		}
		fmt.Printf("Case %s: pass\n", i)
	}
}

func TestAddMatricesWithSingleElements(t *testing.T) {
	a := [][]float64{{1}}
	b := [][]float64{{1}}
	exp := [][]float64{{2}}

	got := AddMatrices(a, b)
	if !Equal(exp, got) {
		t.Errorf("Expected %v, got: %v", exp, got)
	}
}

func TestAddMatricesWithMatrices(t *testing.T) {
	a := [][]float64{{1, 3, 1}, {1, 0, 0}}
	b := [][]float64{{0, 0, 5}, {7, 5, 0}}
	exp := [][]float64{{1, 3, 6}, {8, 5, 0}}

	got := AddMatrices(a, b)
	if !Equal(exp, got) {
		t.Errorf("Expected %v, got: %v", exp, got)
	}
}

func TestAddScalar(t *testing.T) {
	a := [][]float64{{1, 3, 1}, {1, 0, 0}}
	b := float64(5)
	exp := [][]float64{{6, 8, 6}, {6, 5, 5}}

	AddScalar(a, b)
	if !Equal(a, exp) {
		t.Errorf("Expected %v, got: %v", exp, a)
	}
}

func TestMultiplyScalar(t *testing.T) {
	a := [][]float64{{1, 8, -3}, {4, -2, 5}}
	b := float64(2)
	exp := [][]float64{{2, 16, -6}, {8, -4, 10}}

	MultiplyScalar(a, b)
	if !Equal(a, exp) {
		t.Errorf("Expected %v, got: %v", exp, a)
	}
}

func TestTranspose(t *testing.T) {
	tests := map[string]struct {
		a   [][]float64
		exp [][]float64
	}{
		"two_by_three": {[][]float64{{1, 8, -3}, {4, -2, 5}}, [][]float64{{1, 4}, {8, -2}, {-3, 5}}},
		"one_by_one":   {[][]float64{{1}}, [][]float64{{1}}},
		"one_by_two":   {[][]float64{{1, 2}}, [][]float64{{1}, {2}}},
		"two_by_one":   {[][]float64{{1}, {2}}, [][]float64{{1, 2}}},
	}

	for i := range tests {
		got := Transpose(tests[i].a)
		if !Equal(got, tests[i].exp) {
			t.Errorf("Case %s: expected %v, got: %v", i, tests[i].exp, got)
		}
		fmt.Printf("Case %s: pass\n", i)
	}
}

func TestMultiplyMatrices(t *testing.T) {
	tests := map[string]struct {
		a   [][]float64
		b   [][]float64
		exp [][]float64
	}{
		"one_by_one": {
			[][]float64{{4}},
			[][]float64{{3}},
			[][]float64{{12}},
		},
		"small": {
			[][]float64{{1}, {2}},
			[][]float64{{3, 4}},
			[][]float64{{3, 4}, {6, 8}},
		},
		"big": {
			[][]float64{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24}},
			[][]float64{{1, 2, 3, 4, 5, 6, 7, 8}, {9, 10, 11, 12, 13, 14, 15, 16}, {17, 18, 19, 20, 21, 22, 23, 24}, {25, 26, 27, 28, 29, 30, 31, 32}, {33, 34, 35, 36, 37, 38, 39, 40}, {41, 42, 43, 44, 45, 46, 47, 48}},
			[][]float64{{581, 602, 623, 644, 665, 686, 707, 728}, {1337, 1394, 1451, 1508, 1565, 1622, 1679, 1736}, {2093, 2186, 2279, 2372, 2465, 2558, 2651, 2744}, {2849, 2978, 3107, 3236, 3365, 3494, 3623, 3752}},
		},
	}

	for i := range tests {
		got := MultiplyMatrices(tests[i].a, tests[i].b)
		if !Equal(got, tests[i].exp) {
			t.Errorf("Case %s: expected %v, got: %v", i, tests[i].exp, got)
		}
		fmt.Printf("Case %s: pass\n", i)
	}
}

func BenchmarkMultiplyMatrices(b *testing.B) {
	am := make([][]float64, 1000)
	bm := make([][]float64, 1000)

	for i := range am {
		am[i] = make([]float64, 1000)
		bm[i] = make([]float64, 1000)
		for k := range am {
			am[i][k] = rand.NormFloat64()
			bm[i][k] = rand.NormFloat64()
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MultiplyMatrices(am, bm)
	}
}
