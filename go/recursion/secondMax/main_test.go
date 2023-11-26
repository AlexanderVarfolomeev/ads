package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondMax(t *testing.T) {
	secondMax := SecondMax([]int{2, 5, 4, 3, 5})
	assert.Equal(t, 5, secondMax)

	secondMax = SecondMax([]int{2, 5, 4, 3})
	assert.Equal(t, 4, secondMax)

	secondMax = SecondMax([]int{2})
	assert.Equal(t, 2, secondMax)

	secondMax = SecondMax([]int{})
	assert.Equal(t, 0, secondMax)
}
