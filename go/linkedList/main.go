package main

import (
	"errors"
	"os"
	"reflect"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

func (l *LinkedList) Count() int {
	cnt := 0

	node := l.head
	for node != nil {
		cnt++
		node = node.next
	}

	return cnt
}

func (l *LinkedList) Find(n int) (Node, error) {
	node := l.head
	for node != nil {
		if node.value == n {
			return *node, nil
		}
		node = node.next
	}

	return Node{}, errors.New("node not found")
}

func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node

	node := l.head
	for node != nil {
		if node.value == n {
			nodes = append(nodes, *node)
		}
		node = node.next
	}

	return nodes
}

func (l *LinkedList) Delete(n int, all bool) {
	deleted := false

	for l.head != nil && l.head.value == n {
		l.head = l.head.next
		if !all {
			deleted = true
			break
		}
	}

	if l.head == nil {
		l.tail = nil
		return
	}

	if deleted {
		return
	}

	for node := l.head; node != nil && node.next != nil; {
		if node.next.value == n {
			if node.next == l.tail {
				l.tail = node
			}

			node.next = node.next.next
			if !all {
				return
			}
		} else {
			node = node.next
		}
	}

}

func (l *LinkedList) deleteHead(n int) {
	for l.head != nil && l.head.value == n {
		l.head = l.head.next
	}

	if l.head == l.tail && l.head.value == n || l.head == nil {
		l.head = nil
		l.tail = nil
	}
}

func (l *LinkedList) Insert(after *Node, add Node) {
	node := l.head
	for node != nil {
		if node == after {
			add.next = node.next
			node.next = &add

			if node == l.tail {
				l.tail = &add
			}
			return
		}
		node = node.next
	}
}

func (l *LinkedList) InsertFirst(first Node) {
	first.next = l.head
	l.head = &first

	if l.head.next == nil {
		l.tail = &first
	}
}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}
