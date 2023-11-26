package main

func PrintEventIndexValue(list []int) {
	printEventIndexValue(list, 0)
}

func printEventIndexValue(list []int, index int) {
	if len(list) == 0 {
		return
	}

	if index%2 == 0 {
		println(list[0])
	}

	printEventIndexValue(list[1:], index+1)
}

func main() {
	l := []int{1, 2, 5, 53, 44, 46}
	PrintEventIndexValue(l)
}
