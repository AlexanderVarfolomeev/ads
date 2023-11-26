package main

import (
	"asd/deque"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondMax(t *testing.T) {
	d := deque.Deque[int]{}
	d.AddTail(2)
	d.AddTail(5)
	d.AddTail(4)
	d.AddTail(3)
	d.AddTail(5)
	secondMax, err := SecondMax(d)
	assert.Equal(t, 5, secondMax)
	assert.NoError(t, err)

	d = deque.Deque[int]{}
	d.AddTail(2)
	d.AddTail(5)
	d.AddTail(4)
	d.AddTail(3)
	secondMax, err = SecondMax(d)
	assert.Equal(t, 4, secondMax)
	assert.NoError(t, err)

	d = deque.Deque[int]{}
	d.AddTail(2)
	secondMax, err = SecondMax(d)
	assert.Equal(t, 0, secondMax)
	assert.Error(t, err)

	d = deque.Deque[int]{}
	secondMax, err = SecondMax(d)
	assert.Equal(t, 0, secondMax)
	assert.Error(t, err)
}
