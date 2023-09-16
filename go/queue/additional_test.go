package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CircleQueue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q = circle(q, 2)

	v, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 3, v)

	v, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, v)

	v, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 2, v)
}

func Test_CircleEmptyQueue(t *testing.T) {
	q := Queue[int]{}

	q = circle(q, 2)

	assert.Equal(t, 0, q.Size())
}
