package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	res := SumOfDigits(1)
	assert.Equal(t, 1, res)

	res = SumOfDigits(5)
	assert.Equal(t, 5, res)

	res = SumOfDigits(0)
	assert.Equal(t, 0, res)

	res = SumOfDigits(15)
	assert.Equal(t, 6, res)

	res = SumOfDigits(999)
	assert.Equal(t, 27, res)
}
