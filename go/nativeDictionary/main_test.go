package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNativeDictionary_Put(t *testing.T) {
	nd := Init[string](30)
	for i := 0; i < 500; i++ {
		nd.Put(strconv.Itoa(i), strconv.Itoa(i))

		get, err := nd.Get(strconv.Itoa(i))
		assert.NoError(t, err)
		assert.Equal(t, strconv.Itoa(i), get)
	}
}

func TestNativeDictionary_Get_NotFound(t *testing.T) {
	nd := Init[string](1)
	nd.Put("1", "1")
	nd.Put("2", "2")
	nd.Put("3", "3")

	_, err := nd.Get("5")
	assert.Error(t, err)
}

func TestNativeDictionary_IsKey(t *testing.T) {
	nd := Init[string](30)
	for i := 0; i < 500; i++ {
		nd.Put(strconv.Itoa(i), strconv.Itoa(i))

		key := nd.IsKey(strconv.Itoa(i))
		assert.True(t, key)
	}
}

func TestNativeDictionary_IsKey_NotFound(t *testing.T) {
	nd := Init[string](1)
	nd.Put("1", "1")
	nd.Put("2", "2")
	nd.Put("3", "3")

	key := nd.IsKey("5")
	assert.False(t, key)
}
