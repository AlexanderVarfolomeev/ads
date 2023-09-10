package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummyLinkedList2_AddInTail_EmptyList(t *testing.T) {
	l := NewDummyList()

	l.AddInTail(DummyNode{value: 12})

	assert.NotNil(t, l.Tail())
	assert.NotNil(t, l.Head())
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 12, l.Head().value)
	assert.Equal(t, 12, l.Tail().value)
}

func TestDummyLinkedList2_AddInTail_NotEmptyList(t *testing.T) {
	l := NewDummyList()
	l.AddInTail(DummyNode{value: 1})
	l.AddInTail(DummyNode{value: 2})
	l.AddInTail(DummyNode{value: 12})

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 1, l.Head().value)
	assert.Equal(t, 12, l.Tail().value)
}

func TestDummyLinkedList2_Count_EmptyList(t *testing.T) {
	l := NewDummyList()
	c := l.Count()

	assert.Equal(t, 0, c)
}

func TestDummyLinkedList2_Count_NotEmptyList(t *testing.T) {
	l := NewDummyList()
	l.AddInTail(DummyNode{value: 1})
	l.AddInTail(DummyNode{value: 2})
	l.AddInTail(DummyNode{value: 3})

	c := l.Count()

	assert.Equal(t, 3, c)
}

func TestDummyLinkedList2_Count_NotEmptyListAfterDeleteElement(t *testing.T) {
	l := NewDummyList()
	l.AddInTail(DummyNode{value: 1})
	l.AddInTail(DummyNode{value: 2})
	l.AddInTail(DummyNode{value: 3})

	l.Delete(2, false)
	c := l.Count()

	assert.Equal(t, 2, c)
}

func TestDummyLinkedList2_Count_EmptyListAfterDeleteAllElements(t *testing.T) {
	l := NewDummyList()
	l.AddInTail(DummyNode{value: 2})
	l.AddInTail(DummyNode{value: 2})
	l.AddInTail(DummyNode{value: 2})

	l.Delete(2, true)
	c := l.Count()

	assert.Equal(t, 0, c)
}

func TestDummyLinkedList2_Find_EmptyList(t *testing.T) {
	l := NewDummyList()

	find, err := l.Find(2)

	assert.Error(t, err)
	assert.Equal(t, DummyNode{}, find)
}

func TestDummyLinkedList2_Find_NotFound(t *testing.T) {
	l := NewDummyList()
	l.AddInTail(DummyNode{value: 1})
	l.AddInTail(DummyNode{value: 2})
	l.AddInTail(DummyNode{value: 3})

	find, err := l.Find(4)

	assert.Error(t, err)
	assert.Equal(t, DummyNode{}, find)
}

func TestDummyLinkedList2_Find_SuccessFound(t *testing.T) {
	l := NewDummyList()
	l.AddInTail(DummyNode{value: 1})
	l.AddInTail(DummyNode{value: 2})
	l.AddInTail(DummyNode{value: 3})

	find, err := l.Find(3)

	assert.NoError(t, err)
	assert.Equal(t, 3, find.value)
	assert.Equal(t, *l.Tail(), find)
}

func TestDummyLinkedList2_FindAll_EmptyList(t *testing.T) {
	l := NewDummyList()

	find := l.FindAll(3)

	assert.Nil(t, find)
}

func TestDummyLinkedList2_FindAll_NotEmptyList(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 2}
	n2 := DummyNode{value: 1}
	n3 := DummyNode{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n3)

	find := l.FindAll(2)

	assert.NotNil(t, find)
	assert.Len(t, find, 2)
	assert.Equal(t, *l.Head(), find[0])
	assert.Equal(t, *l.Tail(), find[1])
}

func TestDummyLinkedList2_DeleteFirst_EmptyList(t *testing.T) {
	l := NewDummyList()

	l.Delete(2, false)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Head())
	assert.Nil(t, l.Tail())
}

func TestDummyLinkedList2_DeleteFirst_OneElement(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 2}
	l.AddInTail(n1)

	l.Delete(2, false)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Head())
	assert.Nil(t, l.Tail())
}

func TestDummyLinkedList2_DeleteFirst_NotEmptyList_HeadElement(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 2}
	n2 := DummyNode{value: 1}
	n3 := DummyNode{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n3)

	l.Delete(2, false)

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, l.Head().value)
	assert.Equal(t, 2, l.Head().next.value)
	assert.Equal(t, 2, l.Tail().value)
	assert.Equal(t, 1, l.Tail().prev.value)
}

func TestDummyLinkedList2_DeleteFirst_NotEmptyList_NotHeadElement(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 2}
	n2 := DummyNode{value: 1}
	n3 := DummyNode{value: 2}
	n4 := DummyNode{value: 1}
	n5 := DummyNode{value: 3}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n3)
	l.AddInTail(n4)
	l.AddInTail(n5)

	del := l.Tail()

	l.Delete(3, false)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, n4.value, l.Tail().value)
	assert.Equal(t, 2, l.Head().value)
	assert.True(t, l.Tail() != del)
}

func TestDummyLinkedList2_DeleteAll_NotEmptyList_SameRow(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 1}
	n2 := DummyNode{value: 2}
	n3 := DummyNode{value: 3}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n2)
	l.AddInTail(n2)
	l.AddInTail(n2)
	l.AddInTail(n3)

	l.Delete(2, true)

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, l.Head().value)
}

func TestDummyLinkedList2_DeleteAll_EmptyList(t *testing.T) {
	l := NewDummyList()

	l.Delete(2, true)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Tail())
	assert.Nil(t, l.Head())
}

func TestDummyLinkedList2_DeleteAll_NotEmptyList_OneElement(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 1}

	l.AddInTail(n1)

	l.Delete(1, true)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Tail())
	assert.Nil(t, l.Head())
}

func TestDummyLinkedList2_DeleteAll_NotEmptyList_AllElements(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 1}

	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n1)

	l.Delete(1, true)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Tail())
	assert.Nil(t, l.Head())
}

func TestDummyLinkedList2_DeleteAll_NotEmptyList_HeadElement(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 1}
	n2 := DummyNode{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n1)

	l.Delete(1, true)

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Tail().value)
	assert.Equal(t, 2, l.Head().value)
}

func TestDummyLinkedList2_InsertFirst_EmptyList(t *testing.T) {
	l := NewDummyList()

	n := DummyNode{value: 1}
	l.InsertFirst(n)
	n1 := l.Head()
	n2 := l.Tail()

	assert.Equal(t, n1, n2)
	assert.Equal(t, l.Head(), l.Tail())
}

func TestDummyLinkedList2_InsertFirst_NotEmptyList(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 1}
	n2 := DummyNode{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n1)

	n3 := DummyNode{value: 5}
	l.InsertFirst(n3)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 1, l.Tail().value)
	assert.Equal(t, 5, l.Head().value)
}

func TestDummyLinkedList2_Insert_InMiddle(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 1}
	n2 := DummyNode{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n2)

	n3 := DummyNode{value: 5}
	l.Insert(l.Head().next, n3)

	assert.Equal(t, 5, l.Count())
	assert.Equal(t, 2, l.Tail().value)
	assert.Equal(t, 1, l.Head().value)
	assert.Equal(t, 5, l.Head().next.next.value)
	assert.Equal(t, 5, l.Tail().prev.prev.value)
}

func TestDummyLinkedList2_Insert_AfterHead(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 1}
	n2 := DummyNode{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)

	n3 := DummyNode{value: 5}
	l.Insert(l.Head(), n3)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 2, l.Tail().value)
	assert.Equal(t, 1, l.Head().value)
	assert.Equal(t, 5, l.Head().next.value)
}

func TestDummyLinkedList2_Insert_AfterTail(t *testing.T) {
	l := NewDummyList()
	n1 := DummyNode{value: 1}
	n2 := DummyNode{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)

	n3 := DummyNode{value: 5}
	l.Insert(l.Tail(), n3)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 5, l.Tail().value)
	assert.Equal(t, 1, l.Head().value)
}

func TestDummyLinkedList2_Clean_EmptyList(t *testing.T) {
	l := NewDummyList()

	l.Clean()

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Tail())
	assert.Nil(t, l.Head())
}

func TestDummyLinkedList2_Clean_NotEmptyList(t *testing.T) {
	l := NewDummyList()

	n1 := DummyNode{value: 1}
	n2 := DummyNode{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)

	l.Clean()
	c := l.Count()

	assert.Equal(t, 0, c)
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Tail())
	assert.Nil(t, l.Head())
}
