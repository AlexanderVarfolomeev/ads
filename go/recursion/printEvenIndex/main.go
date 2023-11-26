package main

func PrintEvenIndexValue(list []int) {
	printEvenIndexValue(list, 0)
}

func printEvenIndexValue(list []int, index int) {
	if len(list) == 0 {
		return
	}

	if index%2 == 0 {
		println(list[0])
	}

	printEvenIndexValue(list[1:], index+1)
}

func main() {
	l := []int{1, 2, 5, 53, 44, 46}
	PrintEvenIndexValue(l)
}
