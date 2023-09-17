package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque_AddFront(t *testing.T) {
	d := Deque[int]{}

	d.AddFront(1)
	d.AddFront(2)
	d.AddFront(3)

	n := d.head

	for i := 3; i > 0; i-- {
		assert.Equal(t, i, n.value)
		n = n.next
	}

	assert.Equal(t, 3, d.Size())
}

func TestDeque_AddTail(t *testing.T) {
	d := Deque[int]{}

	d.AddTail(1)
	d.AddTail(2)
	d.AddTail(3)

	n := d.head

	for i := 1; i <= 3; i++ {
		assert.Equal(t, i, n.value)
		n = n.next
	}

	assert.Equal(t, 3, d.Size())
}

func TestDeque_AddTail_AddFront(t *testing.T) {
	d := Deque[int]{}

	d.AddTail(3)
	d.AddTail(4)
	d.AddFront(2)
	d.AddFront(1)

	n := d.head

	for i := 1; i <= 4; i++ {
		assert.Equal(t, i, n.value)
		n = n.next
	}

	assert.Equal(t, 4, d.Size())
}

func TestDeque_RemoveFront(t *testing.T) {
	d := Deque[int]{}

	d.AddFront(1)
	d.AddFront(2)
	d.AddFront(3)

	for i := 3; i > 0; i-- {
		assert.Nil(t, d.head.prev)

		front, err := d.RemoveFront()
		assert.NoError(t, err)
		assert.Equal(t, i, front)
	}

	assert.Nil(t, d.tail)
	assert.Nil(t, d.tail)
	assert.Equal(t, 0, d.Size())
}

func TestDeque_RemoveTail(t *testing.T) {
	d := Deque[int]{}

	d.AddTail(1)
	d.AddTail(2)
	d.AddTail(3)

	for i := 3; i > 0; i-- {
		assert.Nil(t, d.tail.next)

		tail, err := d.RemoveTail()
		assert.NoError(t, err)
		assert.Equal(t, i, tail)
	}

	assert.Nil(t, d.tail)
	assert.Nil(t, d.tail)
	assert.Equal(t, 0, d.Size())
}
