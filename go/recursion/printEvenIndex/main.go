package main

import "asd/deque"

func PrintEvenIndexValue(list deque.Deque[int]) {
	printEvenIndexValue(list, 0)
}

func printEvenIndexValue(list deque.Deque[int], index int) {
	if list.Size() == 0 {
		return
	}

	val, _ := list.RemoveFront()
	if index%2 == 0 {
		println(val)
	}

	printEvenIndexValue(list, index+1)
}

func main() {
	l := deque.Deque[int]{}
	l.AddTail(1)
	l.AddTail(5)
	l.AddTail(8)
	l.AddTail(5)
	l.AddTail(6)
	l.AddTail(6)
	PrintEvenIndexValue(l)
}
