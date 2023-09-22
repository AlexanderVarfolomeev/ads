package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestHashTable_Put(t *testing.T) {
	h := Init(18, 2)

	puts := make([]int, 0)
	for i := 0; i < 19; i++ {
		puts = append(puts, h.Put(strconv.Itoa(i)))
	}

	for i := 0; i < 19; i++ {
		assert.Equal(t, puts[i], h.Find(strconv.Itoa(i)))
	}
}

func TestHashTable_Put_OverHead(t *testing.T) {
	h := Init(19, 3)

	puts := make([]int, 0)
	for i := 0; i < 200; i++ {
		puts = append(puts, h.Put(strconv.Itoa(i)))
	}

	for i := 0; i < 19; i++ {
		assert.Equal(t, puts[i], h.Find(strconv.Itoa(i)))
	}
}

func TestHashTable_Put_Collisions(t *testing.T) {
	h := Init(19, 3)

	puts := make([]int, 0)
	puts = append(puts, h.Put("acca"))
	puts = append(puts, h.Put("ccaa"))
	puts = append(puts, h.Put("caca"))
	puts = append(puts, h.Put("acac"))

	accKey := h.Find("acca")

	assert.Equal(t, puts[0], accKey)
	assert.Equal(t, puts[1], accKey+3)
	assert.Equal(t, puts[2], (accKey+6)%h.size)
	assert.Equal(t, puts[3], (accKey+9)%h.size)
}

func TestHashTable_Put_Cycle(t *testing.T) {
	h := Init(3, 3)
	puts := make([]int, 0)
	puts = append(puts, h.Put("acca"))
	puts = append(puts, h.Put("ccaa"))
	puts = append(puts, h.Put("caca"))
	puts = append(puts, h.Put("acac"))

	println()
}
