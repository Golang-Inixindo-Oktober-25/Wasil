package math

func Tambah(a, b float64) float64 {
	return a + b
}

func Kurang(a, b float64) float64 {
	return a - b
}

func Kali(a, b float64) float64 {
	return a * b
}

func Bagi(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
