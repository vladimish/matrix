package mat64

// Equal iterates over matrices and
// return false if it finds different elements.
func Equal(a, b [][]float64) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

// AddMatrices returns the sum of a and b matrices.
func AddMatrices(a, b [][]float64) [][]float64 {
	res := make([][]float64, len(a))
	e := make([]float64, len(a)*len(b[0]))
	for i := range res {
		res[i] = e[i*len(b[0]) : (i+1)*len(b[0])]
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			res[i][j] = a[i][j] + b[i][j]
		}
	}

	return res
}

// AddScalar adds b to all elements of a.
func AddScalar(a [][]float64, b float64) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] += b
		}
	}
}

// MultiplyScalar multiplies all elements of a by b.
func MultiplyScalar(a [][]float64, b float64) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] *= b
		}
	}
}

// MultiplyMatrices multiples a on b.
func MultiplyMatrices(a [][]float64, b [][]float64) [][]float64 {
	res := make([][]float64, len(a))
	e := make([]float64, len(a)*len(b[0]))
	for i := range res {
		res[i] = e[i*len(b[0]) : (i+1)*len(b[0])]
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res
}

// Transpose returns transposed matrix of a.
func Transpose(a [][]float64) [][]float64 {
	res := make([][]float64, len(a[0]))
	for i := 0; i < len(a[0]); i++ {
		res[i] = make([]float64, len(a))
		for k := 0; k < len(a); k++ {
			res[i][k] = a[k][i]
		}
	}

	return res
}
