package model

type Matrix []float32

func Dot(a, b Matrix) (sum float32) {
	for i := 0; i < len(b); i++ {
		sum += a[i] * b[i]
	}
	return
}
