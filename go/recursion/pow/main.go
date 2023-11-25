package main

func Pow(x, n float64) float64 {
	if n == 0 || n == 1 {
		return 1
	}

	if n < 0 {
		return 1. / x * Pow(x, n+1)
	}

	return x * Pow(x, n-1)
}

func main() {
	println(Pow(2, 0))
}
