package main

import "errors"

type Node interface {
	Next() Node
	Prev() Node
	SetNext(node Node)
	SetPrev(node Node)
	Value() int
}

type DummyNode struct {
	prev Node
	next Node
}

func (n *DummyNode) Next() Node {
	return n.next
}

func (n *DummyNode) Prev() Node {
	return n.prev
}

func (n *DummyNode) SetNext(node Node) {
	n.next = node
}

func (n *DummyNode) SetPrev(node Node) {
	n.prev = node
}

func (n *DummyNode) Value() int {
	return 0
}

type RealNode struct {
	prev  Node
	next  Node
	value int
}

func (n *RealNode) Next() Node {
	return n.next
}

func (n *RealNode) Prev() Node {
	return n.prev
}

func (n *RealNode) SetNext(node Node) {
	n.next = node
}

func (n *RealNode) SetPrev(node Node) {
	n.prev = node
}

func (n *RealNode) Value() int {
	return n.value
}

type DummyLinkedList2 struct {
	head Node
	tail Node
}

func NewDummyList() *DummyLinkedList2 {
	n := &DummyNode{}
	l := &DummyLinkedList2{
		head: n,
		tail: n,
	}

	l.Clean()
	return l
}

func isDummyNode(n Node) bool {
	_, ok := n.(*DummyNode)
	return ok
}

func (l *DummyLinkedList2) Head() Node {
	if isDummyNode(l.head.Next()) {
		return nil
	}

	return l.head.Next()
}

func (l *DummyLinkedList2) Tail() Node {
	if isDummyNode(l.tail.Prev()) {
		return nil
	}

	return l.tail.Prev()
}

func (l *DummyLinkedList2) AddInTail(item Node) {
	addAfter(l.tail.Prev(), item)
}

func (l *DummyLinkedList2) Count() int {
	count := 0

	for node := l.head.Next(); !isDummyNode(node); node = node.Next() {
		count++
	}

	return count
}

func (l *DummyLinkedList2) Find(n int) (Node, error) {
	for node := l.head.Next(); !isDummyNode(node); node = node.Next() {
		if node.Value() == n {
			return node, nil
		}
	}

	return nil, errors.New("node not found")
}

func (l *DummyLinkedList2) FindAll(n int) []Node {
	var nodes []Node

	for node := l.head.Next(); !isDummyNode(node); node = node.Next() {
		if node.Value() == n {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

func (l *DummyLinkedList2) Delete(n int, all bool) {
	deleted := false

	for node := l.head.Next(); !isDummyNode(node); node = (node).Next() {
		if node.Value() == n {
			node.Prev().SetNext(node.Next())
			node.Next().SetPrev(node.Prev())
			deleted = true
		}

		if !all && deleted {
			return
		}
	}
}

func (l *DummyLinkedList2) Insert(after Node, add Node) {
	addAfter(after, add)
}

func (l *DummyLinkedList2) InsertFirst(first Node) {
	addAfter(l.head, first)
}

func (l *DummyLinkedList2) Clean() {
	l.head.SetNext(l.head)
	l.head.SetPrev(l.head)
}

func addAfter(after Node, add Node) {
	add.SetNext(after.Next())
	add.SetPrev(after)

	after.Next().SetPrev(add)
	after.SetNext(add)
}
