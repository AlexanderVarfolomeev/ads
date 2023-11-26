package main

import "container/list"

func PrintEventIndexValue(list *list.List) {
	printEventIndexValue(list, 0)
}

func printEventIndexValue(list *list.List, index int) {
	if list.Len() == 0 {
		return
	}

	value := list.Front()
	if index%2 == 0 {
		println(value.Value.(int))
	}

	list.Remove(value)
	printEventIndexValue(list, index+1)
}

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(5)
	l.PushBack(54)
	PrintEventIndexValue(l)
}
