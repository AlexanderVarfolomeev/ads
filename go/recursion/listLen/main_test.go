package main

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenList(t *testing.T) {
	l := list.New()
	l.PushBack(5)
	l.PushBack(8)
	l.PushBack(8)

	assert.Equal(t, 3, ListLen(l))

	l2 := list.New()

	assert.Equal(t, 0, ListLen(l2))
}
