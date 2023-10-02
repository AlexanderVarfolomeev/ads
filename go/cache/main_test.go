package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNativeCache_Add(t *testing.T) {
	c := Init[int](200)
	c.Add("1", 1)
	c.Add("2", 2)
	c.Add("3", 3)

	for i := 1; i <= 3; i++ {
		get, err := c.Get(strconv.Itoa(i))
		assert.NoError(t, err)
		assert.Equal(t, i, get)
	}
}

func TestNativeCache_Add_Hit(t *testing.T) {
	c := Init[int](5)
	for i := 1; i <= 5; i++ {
		c.Add(strconv.Itoa(i), i)
	}

	c.Get("1")
	c.Get("2")
	c.Get("2")
	c.Get("3")
	c.Get("3")
	c.Get("3")
	c.Get("4")
	c.Get("4")
	c.Get("4")
	c.Get("4")

	for i := 1; i <= 4; i++ {
		assert.Equal(t, i, c.hits[i])
	}
}

func TestNativeCache_Add_Full(t *testing.T) {
	c := Init[int](5)
	for i := 1; i <= 5; i++ {
		c.Add(strconv.Itoa(i), i)
	}

	c.Get("1")
	c.Get("2")
	c.Get("2")
	c.Get("3")
	c.Get("3")
	c.Get("3")
	c.Get("4")
	c.Get("4")
	c.Get("4")
	c.Get("4")

	for i := 6; i <= 10; i++ {
		c.Add(strconv.Itoa(i), i)

		for j := 0; j < i; j++ {
			c.Get(strconv.Itoa(i))
		}
	}

	for i := 6; i <= 10; i++ {
		get, err := c.Get(strconv.Itoa(i))
		assert.NoError(t, err)
		assert.Equal(t, i, get)
	}
}
