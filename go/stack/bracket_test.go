package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorrectnessOfBrackets_Success(t *testing.T) {
	assert.True(t, CorrectnessOfBrackets("(()((())()))"))
	assert.True(t, CorrectnessOfBrackets("(()()(()))"))
}

func TestCorrectnessOfBrackets_Failed(t *testing.T) {
	assert.False(t, CorrectnessOfBrackets("())("))
	assert.False(t, CorrectnessOfBrackets("))(("))
	assert.False(t, CorrectnessOfBrackets("((())"))
}
