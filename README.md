# MATRIX
Maybe a bit slow, but minimalistic library for basic matrix operations. 

## Available functions
|Function signature|Description|
|---|---|
|Equal(a, b [][]float64) boolEqual(a, b [][]float64) bool|Equal iterates over matrices and<br>return false if it finds different elements.
|AddMatrices(a, b [][]float64) [][]float64|AddMatrices returns the sum of a and b matrices.|
|AddScalar(a [][]float64, b float64)|AddScalar adds b to all elements of a.|
|MultiplyScalar(a [][]float64, b float64)|MultiplyScalar multiplies all elements of a by b.|
|MultiplyMatrices(a [][]float64, b [][]float64) [][]float64|MultiplyMatrices multiples a on b.|
|Transpose(a [][]float64) [][]float64|Transpose returns transposed matrix of a.|