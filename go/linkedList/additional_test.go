package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumLinkedLists_EmptyLists(t *testing.T) {
	l1 := LinkedList{}
	l2 := LinkedList{}

	res, err := Sum(l1, l2)

	assert.NoError(t, err)
	assert.Equal(t, LinkedList{}, res)
}

func TestSumLinkedLists_DifSizes(t *testing.T) {
	l1 := LinkedList{}
	l2 := LinkedList{}

	l1.AddInTail(Node{value: 1})
	l1.AddInTail(Node{value: 2})
	l1.AddInTail(Node{value: 3})

	l2.AddInTail(Node{value: 1})
	l2.AddInTail(Node{value: 3})
	res, err := Sum(l1, l2)

	assert.Error(t, err)
	assert.Equal(t, LinkedList{}, res)
}

func TestSumLinkedLists_Success(t *testing.T) {
	l1 := LinkedList{}
	l2 := LinkedList{}

	l1.AddInTail(Node{value: 1})
	l1.AddInTail(Node{value: 2})
	l1.AddInTail(Node{value: 3})

	l2.AddInTail(Node{value: 1})
	l2.AddInTail(Node{value: 3})
	l2.AddInTail(Node{value: 3})

	res, err := Sum(l1, l2)
	assert.NoError(t, err)

	assert.Equal(t, res.head.value, 2)
	assert.Equal(t, res.head.next.value, 5)
	assert.Equal(t, res.tail.value, 6)
}
