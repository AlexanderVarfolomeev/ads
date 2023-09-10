package main

import "errors"

type DummyNode struct {
	prev  *DummyNode
	next  *DummyNode
	value int
	dummy bool
}

type DummyLinkedList2 struct {
	head *DummyNode
	tail *DummyNode
}

func NewDummyList() *DummyLinkedList2 {
	l := &DummyLinkedList2{
		head: &DummyNode{
			dummy: true,
		},
		tail: &DummyNode{
			dummy: true,
		},
	}

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *DummyLinkedList2) Head() *DummyNode {
	if l.head.next.dummy {
		return nil
	}

	return l.head.next
}

func (l *DummyLinkedList2) Tail() *DummyNode {
	if l.tail.prev.dummy {
		return nil
	}

	return l.tail.prev
}

func (l *DummyLinkedList2) AddInTail(item DummyNode) {
	addAfter(l.tail.prev, &item)
}

func (l *DummyLinkedList2) Count() int {
	count := 0

	for node := l.head.next; !node.dummy; node = node.next {
		count++
	}

	return count
}

func (l *DummyLinkedList2) Find(n int) (DummyNode, error) {
	for node := l.head.next; !node.dummy; node = node.next {
		if node.value == n {
			return *node, nil
		}
	}

	return DummyNode{value: 0, next: nil}, errors.New("node not found")
}

func (l *DummyLinkedList2) FindAll(n int) []DummyNode {
	var nodes []DummyNode

	for node := l.head.next; !node.dummy; node = node.next {
		if node.value == n {
			nodes = append(nodes, *node)
		}
	}

	return nodes
}

func (l *DummyLinkedList2) Delete(n int, all bool) {
	deleted := false

	for node := l.head.next; !node.dummy; node = node.next {
		if node.value == n {
			node.prev.next = node.next
			node.next.prev = node.prev

			deleted = true
		}

		if !all && deleted {
			return
		}
	}
}

func (l *DummyLinkedList2) Insert(after *DummyNode, add DummyNode) {
	for node := l.head.next; !node.dummy; node = node.next {
		if node == after {
			addAfter(after, &add)
			return
		}
	}
}

func (l *DummyLinkedList2) InsertFirst(first DummyNode) {
	addAfter(l.head, &first)
}

func (l *DummyLinkedList2) Clean() {
	l.head.next = l.tail
	l.tail.prev = l.head
}

func addAfter(after *DummyNode, add *DummyNode) {
	add.next = after.next
	add.prev = after

	after.next.prev = add
	after.next = add
}
