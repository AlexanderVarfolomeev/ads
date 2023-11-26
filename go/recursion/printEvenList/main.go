package main

func PrintEventValues(list []int) {
	if len(list) == 0 {
		return
	}

	val := list[0]
	if val%2 == 0 {
		println(val)
	}

	PrintEventValues(list[1:])
}

func main() {
	l := []int{1, 2, 5, 53, 44, 46}
	PrintEventValues(l)
}
