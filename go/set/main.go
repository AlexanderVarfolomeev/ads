package main

import (
        "constraints"
//      "fmt"
	"os"
        "strconv"
)
type PowerSet[T constraints.Ordered] struct {
	dict map[T]struct{}
}

func Init[T constraints.Ordered]() PowerSet[T] {
	return PowerSet[T]{
		dict: make(map[T]struct{}, 20000),
	}
}

func (s *PowerSet[T]) Size() int {
	return len(s.dict)
}

func (s *PowerSet[T]) Put(value T) {
	s.dict[value] = struct{}{}
}

func (s *PowerSet[T]) Get(value T) bool {
	_, ok := s.dict[value]
	return ok
}

func (s *PowerSet[T]) Remove(value T) bool {
	if _, ok := s.dict[value]; ok {
		delete(s.dict, value)
		return true
	}
	return false
}

func (s *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	result := Init[T]()

	for key := range set2.dict {
		if _, ok := s.dict[key]; ok {
			result.Put(key)
		}
	}

	return result
}

func (s *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	result := Init[T]()
	for key := range set2.dict {
		result.Put(key)
	}

	for key := range s.dict {
		result.Put(key)
	}
	return result
}

func (s *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	result := Init[T]()
	for key := range s.dict {
		if _, ok := set2.dict[key]; !ok {
			result.Put(key)
		}
	}
	return result
}

func (s *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
	for key := range set2.dict {
		if _, ok := s.dict[key]; !ok {
			return false
		}
	}
	return true
}
