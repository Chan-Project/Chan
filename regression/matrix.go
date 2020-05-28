package regression

type Matrix []float64

func Dot(a, b Matrix) (sum float64) {
	for i := 0; i < len(b); i++ {
		sum += a[i] * b[i]
	}
	return
}
