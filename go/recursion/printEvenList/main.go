package main

import "container/list"

func PrintEventValues(list *list.List) {
	if list.Len() == 0 {
		return
	}

	val := list.Front()
	if val.Value.(int)%2 == 0 {
		println(val.Value.(int))
	}
	list.Remove(val)
	PrintEventValues(list)
}

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(5)
	l.PushBack(54)
	PrintEventValues(l)
}
