package main

import (
	"container/list"
)

func ListLen(list *list.List) int {
	if list.Len() == 0 {
		return 0
	}
	list.Remove(list.Front()) // в стандартной библиотеке нет pop()
	return 1 + ListLen(list)
}

func main() {
	l := list.New()
	l.PushBack(5)
	l.PushBack(8)

	println(ListLen(l))
}
