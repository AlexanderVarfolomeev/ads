package main

func Pow(n float64, m int) float64 {
	if m == 0 {
		return 1
	}

	if m < 0 {
		return 1. / n * Pow(n, m+1)
	}

	return n * Pow(n, m-1)
}

func main() {
	println(Pow(2, 0))
}
