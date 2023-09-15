package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	s := Queue[int]{}

	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)

	assert.Equal(t, 3, s.Size())

	n := s.head
	for i := 1; i <= 3; i++ {
		assert.Equal(t, n.value, i)
		n = n.next
	}
}

func TestQueue_Dequeue(t *testing.T) {
	s := Queue[int]{}

	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)

	assert.Equal(t, 3, s.Size())

	for i := 1; i <= 3; i++ {
		val, err := s.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, i, val)
		assert.Equal(t, 3-i, s.Size())
	}

	_, err := s.Dequeue()
	assert.Error(t, err)
	assert.Equal(t, 0, s.Size())
}

func TestNilTailAfterDequeue(t *testing.T) {
	// Create a new queue
	q := &Queue[int]{}

	// 1. Enqueue one item.
	q.Enqueue(1)
	assert.Equal(t, 1, q.Size())
	assert.NotNil(t, q.head)
	assert.NotNil(t, q.tail)
	assert.Equal(t, 1, q.head.value)
	assert.Equal(t, 1, q.tail.value)

	// 2. Dequeue that item.
	val, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	assert.Nil(t, q.head)
	assert.Nil(t, q.tail)
	assert.Equal(t, 0, q.Size())

	// 3. Enqueue another item.
	q.Enqueue(2)
	assert.Equal(t, 1, q.Size())
	assert.NotNil(t, q.head)
	assert.NotNil(t, q.tail)
	assert.Equal(t, 2, q.head.value)
	assert.Equal(t, 2, q.tail.value)
	assert.Nil(t, q.tail.next)
}
