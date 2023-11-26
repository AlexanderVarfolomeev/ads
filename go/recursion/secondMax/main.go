package main

import (
	"asd/deque"
	"errors"
)

func SecondMax(list deque.Deque[int]) (int, error) {
	if list.Size() < 2 {
		return 0, errors.New("list contains less than 2 items")
	}
	val1, _ := list.RemoveFront()
	val2, _ := list.RemoveFront()
	m1 := Max(val1, val2)
	m2 := Min(val1, val2)

	return secondMax(list, m1, m2), nil
}

func secondMax(list deque.Deque[int], m1, m2 int) int {
	if list.Size() == 0 {
		return m2
	}
	val, _ := list.RemoveFront()

	if val > m1 {
		m2 = m1
		m1 = m2
	} else if val > m2 {
		m2 = val
	}

	return secondMax(list, m1, m2)
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
