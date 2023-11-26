package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondMax(t *testing.T) {
	secondMax, err := SecondMax([]int{2, 5, 4, 3, 5})
	assert.Equal(t, 5, secondMax)
	assert.NoError(t, err)

	secondMax, err = SecondMax([]int{2, 5, 4, 3})
	assert.Equal(t, 4, secondMax)
	assert.NoError(t, err)

	secondMax, err = SecondMax([]int{2})
	assert.Equal(t, 0, secondMax)
	assert.Error(t, err)

	secondMax, err = SecondMax([]int{})
	assert.Equal(t, 0, secondMax)
	assert.Error(t, err)
}
