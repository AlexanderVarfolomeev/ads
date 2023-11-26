package main

func PrintEventValues(list []int) {
	if len(list) == 0 {
		return
	}

	if list[0]%2 == 0 {
		println(list[0])
	}

	PrintEventValues(list[1:])
}

func main() {
	l := []int{1, 2, 5, 53, 44, 46}
	PrintEventValues(l)
}
