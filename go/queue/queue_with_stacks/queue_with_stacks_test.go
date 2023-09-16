package queue_with_stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Enqueue(t *testing.T) {
	s := Queue[int]{}

	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)

	assert.Equal(t, 3, s.Size())

	for i := 1; i <= 3; i++ {
		val, _ := s.Dequeue()
		assert.Equal(t, i, val)
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
