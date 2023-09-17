package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome_Success(t *testing.T) {
	assert.True(t, isPalindrome("123321"))

	assert.True(t, isPalindrome("123 34 43 321"))
	assert.True(t, isPalindrome("1 3 1"))

	assert.True(t, isPalindrome("а роза упала на лапу Азора"))
	assert.True(t, isPalindrome("Sum summus mus "))
}

func TestIsPalindrome_Failed(t *testing.T) {
	assert.False(t, isPalindrome("1223321"))

	assert.False(t, isPalindrome("1_23 34 43 321"))
	assert.False(t, isPalindrome("1 3 1 1"))

	assert.False(t, isPalindrome("а роза не упала на лапу Азора"))
	assert.False(t, isPalindrome(" 12"))
}
