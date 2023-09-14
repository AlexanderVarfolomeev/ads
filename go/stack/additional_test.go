package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePostfix_Success(t *testing.T) {
	exp1 := "8 2 + 5 * 9 + = "
	val, err := CalculatePostfix(exp1)

	assert.NoError(t, err)
	assert.Equal(t, 59, val)

	exp2 := "1 3 * 5 * 10 + = "
	val, err = CalculatePostfix(exp2)

	assert.NoError(t, err)
	assert.Equal(t, 25, val)
}

func TestCalculatePostfix_Failed(t *testing.T) {
	exp1 := "8 2sdf + 5 * 9 + = "
	_, err := CalculatePostfix(exp1)

	assert.Error(t, err)

	exp2 := "1 3 3 * 5 * 10 + = "
	_, err = CalculatePostfix(exp2)

	assert.Error(t, err)
}
