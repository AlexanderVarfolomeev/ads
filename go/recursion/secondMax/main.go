package main

import "errors"

func SecondMax(list []int) (int, error) {
	if len(list) < 2 {
		return 0, errors.New("list contains less than 2 items")
	}
	m1 := Max(list[0], list[1])
	m2 := Min(list[0], list[1])

	return secondMax(list[2:], m1, m2), nil
}

func secondMax(list []int, m1, m2 int) int {
	if len(list) == 0 {
		return m2
	}
	val := list[0]

	if val > m1 {
		m2 = m1
		m1 = m2
	} else if val > m2 {
		m2 = val
	}

	return secondMax(list[1:], m1, m2)
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func main() {

}
