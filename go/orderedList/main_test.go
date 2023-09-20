package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedList_Add_Ascending_Int(t *testing.T) {
	l := OrderedList[int]{_ascending: true}

	for i := 0; i < 100; i++ {
		l.Add(i)
	}

	n := l.head
	for i := 0; i < 100; i++ {
		assert.NotNil(t, n)
		assert.Equal(t, i, n.value)

		n = n.next
	}
	assert.Equal(t, 100, l.Count())

	l2 := OrderedList[int]{_ascending: true}

	for i := 100; i >= 0; i-- {
		l2.Add(i)
	}

	n2 := l2.head
	for i := 0; i < 100; i++ {
		assert.NotNil(t, n2)
		assert.Equal(t, i, n2.value)

		n2 = n2.next
	}
}

func TestOrderedList_Add_Descending_Int(t *testing.T) {
	l := OrderedList[int]{_ascending: false}

	for i := 0; i < 100; i++ {
		l.Add(i)
	}

	n := l.head
	for i := 99; i >= 0; i-- {
		assert.NotNil(t, n)
		assert.Equal(t, i, n.value)

		n = n.next
	}

	l2 := OrderedList[int]{_ascending: false}

	for i := 100; i >= 0; i-- {
		l2.Add(i)
	}

	n2 := l2.head
	for i := 100; i >= 0; i-- {
		assert.NotNil(t, n2)
		assert.Equal(t, i, n2.value)

		n2 = n2.next
	}
}

func TestOrderedList_Find_Ascending_Int(t *testing.T) {
	l := OrderedList[int]{_ascending: true}

	for i := 0; i < 100; i++ {
		l.Add(i)
	}

	for i := 50; i < 100; i++ {
		find, err := l.Find(i)
		assert.NoError(t, err)
		assert.Equal(t, i, find.value)
	}
}

func TestOrderedList_Find_Ascending_ExitEarly(t *testing.T) {
	l := OrderedList[int]{_ascending: true}

	for i := 0; i < 100; i += 2 {
		l.Add(i)
	}

	_, err := l.Find(13)
	assert.Error(t, err)
}

func TestOrderedList_Find_Descending_Int(t *testing.T) {
	l := OrderedList[int]{_ascending: false}

	for i := 0; i < 100; i++ {
		l.Add(i)
	}

	for i := 50; i < 100; i++ {
		find, err := l.Find(i)
		assert.NoError(t, err)
		assert.Equal(t, i, find.value)
	}
}

func TestOrderedList_Find_Descending_ExitEarly(t *testing.T) {
	l := OrderedList[int]{_ascending: false}

	for i := 0; i < 100; i += 2 {
		l.Add(i)
	}

	_, err := l.Find(97)
	assert.Error(t, err)
}

func TestOrderedList_Delete_Ascending_Int(t *testing.T) {
	l := OrderedList[int]{_ascending: true}

	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	l.Delete(0)
	assert.Equal(t, 99, l.Count())

	_, err := l.Find(0)
	assert.Error(t, err)

	assert.Equal(t, 1, l.head.value)

	for i := 1; i < 100; i++ {
		l.Delete(i)
		assert.Equal(t, 99-i, l.Count())

		_, err := l.Find(i)
		assert.Error(t, err)

		if l.head != nil {
			assert.Equal(t, i+1, l.head.value)
		}
	}

	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}

func TestOrderedList_Delete_Descending_Int(t *testing.T) {
	l := OrderedList[int]{_ascending: false}

	for i := 0; i < 100; i++ {
		l.Add(i)
	}
	l.Delete(0)
	assert.Equal(t, 99, l.Count())

	_, err := l.Find(0)
	assert.Error(t, err)

	assert.Equal(t, 99, l.head.value)

	for i := 1; i < 100; i++ {
		l.Delete(i)
		assert.Equal(t, 99-i, l.Count())

		_, err := l.Find(i)
		assert.Error(t, err)

		if l.head != nil {
			assert.Equal(t, i+1, l.tail.value)
		}
	}

	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}
