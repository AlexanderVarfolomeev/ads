package main

import "errors"

type NativeCache[T any] struct {
	size   int
	slots  []string
	values []T
	hits   []int
}

func Init[T any](sz int) NativeCache[T] {
	ht := NativeCache[T]{size: sz}
	ht.slots = make([]string, sz)
	ht.values = make([]T, sz)
	ht.hits = make([]int, sz)
	return ht
}

func (c *NativeCache[T]) Add(key string, value T) {
	index := c.seekSlot(key)

	if index == -1 {
		index = c.freeSlot()
	}

	c.values[index] = value
	c.slots[index] = key
	c.hits[index] = 0
}

func (c *NativeCache[T]) Get(key string) (T, error) {
	var result T
	index := c.seekSlot(key)

	if index == -1 {
		return result, errors.New("not found")
	}

	if c.slots[index] == key {
		c.hits[index]++
		return c.values[index], nil
	}

	return result, errors.New("not found")
}

func (c *NativeCache[T]) hashFun(value string) int {
	hash := 1
	for i, c := range value {
		hash += int(c) * i
	}

	return hash % c.size
}

// free slot with min hits
func (c *NativeCache[T]) freeSlot() int {
	min := c.hits[0]
	index := 0
	for i, hit := range c.hits {
		if hit <= min {
			min = hit
			index = i
		}
	}

	return index
}

// find free slot
func (c *NativeCache[T]) seekSlot(value string) int {
	key := c.hashFun(value)
	fromStart := false
	start := key
	for c.slots[key] != value && c.slots[key] != "" {
		key++

		if fromStart && key >= start {
			return -1
		}

		if key == c.size && !fromStart {
			fromStart = true
			key = 0
		}
	}

	return key
}
