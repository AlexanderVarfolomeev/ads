package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilter_Add(t *testing.T) {
	values := []string{"0123456789", "1234567890", "2345678901",
		"3456789012", "4567890123", "5678901234", "6789012345", "7890123456", "8901234567", "9012345678",
	}

	f := BloomFilter{
		filter:     0,
		filter_len: 32,
	}

	for _, value := range values {
		f.Add(value)
	}

	for _, value := range values {
		isValue := f.IsValue(value)
		assert.True(t, isValue)
	}
}

func TestBloomFilter_IsValue_Empty(t *testing.T) {
	values := []string{"0123456789", "1234567890", "2345678901",
		"3456789012", "4567890123", "5678901234", "6789012345", "7890123456", "8901234567", "9012345678",
	}

	f := BloomFilter{
		filter:     0,
		filter_len: 32,
	}

	for _, value := range values {
		isValue := f.IsValue(value)
		assert.False(t, isValue)
	}
}
