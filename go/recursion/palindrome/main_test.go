package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	res := IsPalindrome("123321")
	assert.True(t, res)

	res = IsPalindrome("")
	assert.True(t, res)

	res = IsPalindrome("12321")
	assert.True(t, res)

	res = IsPalindrome("1")
	assert.True(t, res)

	res = IsPalindrome("1231")
	assert.False(t, res)

	res = IsPalindrome("212321")
	assert.False(t, res)
}
