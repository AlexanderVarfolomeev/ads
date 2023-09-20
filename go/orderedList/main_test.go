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

func TestOrderedList_Add_Ascending_Int_Same(t *testing.T) {

	l2 := OrderedList[int]{_ascending: true}

	for i := 100; i >= 0; i-- {
		l2.Add(100)
	}

	n2 := l2.head
	for i := 0; i < 100; i++ {
		assert.NotNil(t, n2)
		assert.Equal(t, 100, n2.value)

		n2 = n2.next
	}
}

func TestOrderedList_Add_Ascending_Float(t *testing.T) {
	l := OrderedList[float64]{_ascending: true}

	for i := 0; i < 100; i++ {
		l.Add(float64(i) - 0.5)
	}

	n := l.head
	for i := 0; i < 100; i++ {
		assert.NotNil(t, n)
		assert.Equal(t, float64(i)-0.5, n.value)

		n = n.next
	}
	assert.Equal(t, 100, l.Count())

	l2 := OrderedList[float64]{_ascending: true}

	for i := 100; i >= 0; i-- {
		l2.Add(float64(i) - 0.5)
	}

	n2 := l2.head
	for i := 0; i < 100; i++ {
		assert.NotNil(t, n2)
		assert.Equal(t, float64(i)-0.5, n2.value)

		n2 = n2.next
	}
}

func TestOrderedList_Add_Ascending_Int_ToMiddle(t *testing.T) {
	l := OrderedList[int]{_ascending: true}

	l.Add(1)
	l.Add(2)
	l.Add(4)
	l.Add(5)
	l.Add(3)

	for i := 1; i <= 5; i++ {
		find, err := l.Find(i)
		assert.NoError(t, err)
		assert.Equal(t, i, find.value)
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

func Test_Add_Ordered_String(t *testing.T) {
	list := &OrderedList[string]{_ascending: true}

	list.Add("apple")
	list.Add("banana")
	list.Add("apple")

	assert.Equal(t, 3, list.Count())

	firstNode := list.head
	secondNode := firstNode.next
	thirdNode := secondNode.next

	assert.NotNil(t, firstNode)
	assert.NotNil(t, secondNode)
	assert.NotNil(t, thirdNode)
	assert.Nil(t, thirdNode.next)

	assert.Equal(t, "apple", firstNode.value)
	assert.Equal(t, "apple", secondNode.value)
	assert.Equal(t, "banana", thirdNode.value)
	assert.True(t, firstNode.next == secondNode)
	assert.True(t, secondNode.prev == firstNode)
	assert.True(t, secondNode.next == thirdNode)
	assert.True(t, thirdNode.prev == secondNode)
}

func Test_Add_Desc_String(t *testing.T) {
	list := &OrderedList[string]{_ascending: true}

	list.Add("cherry")
	list.Add("banana")
	list.Add("apple")

	assert.Equal(t, 3, list.Count())

	firstNode := list.head
	secondNode := firstNode.next
	thirdNode := secondNode.next

	assert.NotNil(t, firstNode)
	assert.NotNil(t, secondNode)
	assert.NotNil(t, thirdNode)
	assert.Nil(t, thirdNode.next)

	assert.Equal(t, "apple", firstNode.value)
	assert.Equal(t, "banana", secondNode.value)
	assert.Equal(t, "cherry", thirdNode.value)
}

func Test_Add_WithSpaces(t *testing.T) {
	list := &OrderedList[string]{_ascending: true}
	list.Add(" cherry ")
	list.Add(" banana ")
	list.Add("apple")

	firstNode := list.head
	secondNode := firstNode.next
	thirdNode := secondNode.next

	assert.NotNil(t, firstNode)
	assert.NotNil(t, secondNode)
	assert.NotNil(t, thirdNode)
	assert.Nil(t, thirdNode.next)

	assert.Equal(t, "apple", firstNode.value)
	assert.Equal(t, " banana ", secondNode.value)
	assert.Equal(t, " cherry ", thirdNode.value)
}

func TestOrderedList_Delete(t *testing.T) {
	l := &OrderedList[int]{_ascending: true}
	l.Add(5)
	l.Add(2)
	l.Add(4)

	assert.Equal(t, 3, l.Count())

	l.Delete(4)

	assert.Equal(t, 2, l.Count())
}

func TestOrderedList_Add(t *testing.T) {
	l1 := &OrderedList[int]{_ascending: false}
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)
	println()
}

func TestOrderedList(t *testing.T) {
	list := &OrderedList[int]{_ascending: true}

	list.Add(5)

	assert.Equal(t, 5, list.head.value)
	assert.Equal(t, 5, list.tail.value)
	assert.Nil(t, list.head.prev)
	assert.Nil(t, list.head.next)

	list.Add(3)

	assert.Equal(t, 3, list.head.value)
	assert.Equal(t, 5, list.tail.value)
	assert.Nil(t, list.head.prev)
	assert.Equal(t, list.tail, list.head.next)

	list.Add(7)

	assert.Equal(t, 7, list.tail.value)
	assert.Equal(t, 5, list.tail.prev.value)
}

func TestOrderedList_(t *testing.T) {
	l := &OrderedList[int]{
		_ascending: true,
	}

	l.Add(5)
	l.Add(3)
	l.Add(7)
	l.Add(4)

	values := []int{}
	n := l.head
	for n != nil {
		values = append(values, n.value)
		n = n.next
	}

	assert.Equal(t, []int{3, 4, 5, 7}, values)
}
