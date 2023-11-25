package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowPositive(t *testing.T) {
	res := Pow(1, 5)
	assert.Equal(t, 1., res)

	res = Pow(2, 5)
	assert.Equal(t, 32., res)

	res = Pow(3, 3)
	assert.Equal(t, 27., res)

	res = Pow(2, 0)
	assert.Equal(t, 1., res)
}

func TestPowNegative(t *testing.T) {
	res := Pow(1, -5)
	assert.Equal(t, 1., res)

	res = Pow(2, -5)
	assert.Equal(t, 1./32, res)

	res = Pow(3, -3)
	assert.Equal(t, 1./27, res)
}
