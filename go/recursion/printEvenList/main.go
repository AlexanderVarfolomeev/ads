package main

import (
	"asd/deque"
)

func PrintEvenValues(list deque.Deque[int]) {
	if list.Size() == 0 {
		return
	}

	val, _ := list.RemoveFront()
	if val%2 == 0 {
		println(val)
	}

	PrintEvenValues(list)
}

func main() {
	l := deque.Deque[int]{}
	l.AddTail(1)
	l.AddTail(5)
	l.AddTail(8)
	l.AddTail(5)
	l.AddTail(6)
	l.AddTail(6)
	PrintEvenValues(l)
}
