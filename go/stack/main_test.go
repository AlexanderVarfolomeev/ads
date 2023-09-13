package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := Stack[int]{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3, s.Size())

	val, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, val)

	val, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = s.Pop()
	assert.Error(t, err)
	assert.Equal(t, 0, val)
}

func TestStack_Pop_Empty(t *testing.T) {
	s := Stack[int]{}

	assert.Equal(t, 0, s.Size())

	val, err := s.Pop()
	assert.Error(t, err)
	assert.Equal(t, 0, val)
}

func TestStack_Pop_Not_Empty(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, s.Size())

	val, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
}

func TestStack_Peek_Empty(t *testing.T) {
	s := Stack[int]{}

	assert.Equal(t, 0, s.Size())

	val, err := s.Peek()
	assert.Error(t, err)
	assert.Equal(t, 0, val)
}

func TestStack_Peek_Not_Empty(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, s.Size())

	val, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	assert.Equal(t, 2, s.Size())
}

func TestStack_Size_Empty(t *testing.T) {
	s := Stack[int]{}

	val := s.Size()
	assert.Equal(t, 0, val)
}

func TestStack_Size_Not_Empty(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, s.Size())
}

func TestStack_Size_Not_Empty_AfterPop(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, s.Size())

	s.Pop()

	assert.Equal(t, 1, s.Size())
}
