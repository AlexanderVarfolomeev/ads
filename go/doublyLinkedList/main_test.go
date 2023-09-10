package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList2_AddInTail_EmptyList(t *testing.T) {
	l := LinkedList2{}

	l.AddInTail(Node{value: 12})

	assert.NotNil(t, l.tail)
	assert.NotNil(t, l.head)
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 12, l.head.value)
	assert.Equal(t, 12, l.tail.value)
}

func TestLinkedList2_AddInTail_NotEmptyList(t *testing.T) {
	l := LinkedList2{}
	l.AddInTail(Node{value: 1})
	l.AddInTail(Node{value: 2})
	l.AddInTail(Node{value: 12})

	assert.NotNil(t, l.tail)
	assert.NotNil(t, l.head)
	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 12, l.tail.value)
}

func TestLinkedList2_Count_EmptyList(t *testing.T) {
	l := LinkedList2{}
	c := l.Count()

	assert.Equal(t, 0, c)
}

func TestLinkedList2_Count_NotEmptyList(t *testing.T) {
	l := LinkedList2{}
	l.AddInTail(Node{value: 1})
	l.AddInTail(Node{value: 2})
	l.AddInTail(Node{value: 3})

	c := l.Count()

	assert.Equal(t, 3, c)
}

func TestLinkedList2_Count_NotEmptyListAfterDeleteElement(t *testing.T) {
	l := LinkedList2{}
	l.AddInTail(Node{value: 1})
	l.AddInTail(Node{value: 2})
	l.AddInTail(Node{value: 3})

	l.Delete(2, false)
	c := l.Count()

	assert.Equal(t, 2, c)
}

func TestLinkedList2_Count_EmptyListAfterDeleteAllElements(t *testing.T) {
	l := LinkedList2{}
	l.AddInTail(Node{value: 2})
	l.AddInTail(Node{value: 2})
	l.AddInTail(Node{value: 2})

	l.Delete(2, true)
	c := l.Count()

	assert.Equal(t, 0, c)
}

func TestLinkedList2_Find_EmptyList(t *testing.T) {
	l := LinkedList2{}

	find, err := l.Find(2)

	assert.Error(t, err)
	assert.Equal(t, Node{}, find)
}

func TestLinkedList2_Find_NotFound(t *testing.T) {
	l := LinkedList2{}
	l.AddInTail(Node{value: 1})
	l.AddInTail(Node{value: 2})
	l.AddInTail(Node{value: 3})

	find, err := l.Find(4)

	assert.Error(t, err)
	assert.Equal(t, Node{}, find)
}

func TestLinkedList2_Find_SuccessFound(t *testing.T) {
	l := LinkedList2{}
	l.AddInTail(Node{value: 1})
	l.AddInTail(Node{value: 2})
	l.AddInTail(Node{value: 3})

	find, err := l.Find(3)

	assert.NoError(t, err)
	assert.Equal(t, 3, find.value)
	assert.Equal(t, *l.tail, find)
}

func TestLinkedList2_FindAll_EmptyList(t *testing.T) {
	l := LinkedList2{}

	find := l.FindAll(3)

	assert.Nil(t, find)
}

func TestLinkedList2_FindAll_NotEmptyList(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 2}
	n2 := Node{value: 1}
	n3 := Node{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n3)

	find := l.FindAll(2)

	assert.NotNil(t, find)
	assert.Len(t, find, 2)
	assert.Equal(t, *l.head, find[0])
	assert.Equal(t, *l.tail, find[1])
}

func TestLinkedList2_DeleteFirst_EmptyList(t *testing.T) {
	l := LinkedList2{}

	l.Delete(2, false)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}

func TestLinkedList2_DeleteFirst_OneElement(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 2}
	l.AddInTail(n1)

	l.Delete(2, false)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}

func TestLinkedList2_DeleteFirst_NotEmptyList_HeadElement(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 2}
	n2 := Node{value: 1}
	n3 := Node{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n3)

	l.Delete(2, false)

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 2, l.head.next.value)
	assert.Equal(t, 2, l.tail.value)
	assert.Equal(t, 1, l.tail.prev.value)
}

func TestLinkedList2_DeleteFirst_NotEmptyList_NotHeadElement(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 2}
	n2 := Node{value: 1}
	n3 := Node{value: 2}
	n4 := Node{value: 1}
	n5 := Node{value: 3}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n3)
	l.AddInTail(n4)
	l.AddInTail(n5)

	del := l.tail

	l.Delete(3, false)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, n4.value, l.tail.value)
	assert.Equal(t, 2, l.head.value)
	assert.True(t, l.tail != del)
}

func TestLinkedList2_DeleteAll_NotEmptyList_SameRow(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 1}
	n2 := Node{value: 2}
	n3 := Node{value: 3}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n2)
	l.AddInTail(n2)
	l.AddInTail(n2)
	l.AddInTail(n3)

	l.Delete(2, true)

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, l.head.value)
}

func TestLinkedList2_DeleteAll_EmptyList(t *testing.T) {
	l := LinkedList2{}

	l.Delete(2, true)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.tail)
	assert.Nil(t, l.head)
}

func TestLinkedList2_DeleteAll_NotEmptyList_OneElement(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 1}

	l.AddInTail(n1)

	l.Delete(1, true)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.tail)
	assert.Nil(t, l.head)
}

func TestLinkedList2_DeleteAll_NotEmptyList_AllElements(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 1}

	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n1)

	l.Delete(1, true)

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.tail)
	assert.Nil(t, l.head)
}

func TestLinkedList2_DeleteAll_NotEmptyList_HeadElement(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 1}
	n2 := Node{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n1)

	l.Delete(1, true)

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.tail.value)
	assert.Equal(t, 2, l.head.value)
}

func TestLinkedList2_InsertFirst_EmptyList(t *testing.T) {
	l := LinkedList2{}

	n := Node{value: 1}
	l.InsertFirst(n)

	assert.Equal(t, n, *l.head)
	assert.Equal(t, n, *l.tail)
}

func TestLinkedList2_InsertFirst_NotEmptyList(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 1}
	n2 := Node{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n1)

	n3 := Node{value: 5}
	l.InsertFirst(n3)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 1, l.tail.value)
	assert.Equal(t, 5, l.head.value)
}

func TestLinkedList2_Insert_InMiddle(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 1}
	n2 := Node{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n1)
	l.AddInTail(n2)
	l.AddInTail(n2)

	n3 := Node{value: 5}
	l.Insert(l.head.next, n3)

	assert.Equal(t, 5, l.Count())
	assert.Equal(t, 2, l.tail.value)
	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 5, l.head.next.next.value)
	assert.Equal(t, 5, l.tail.prev.prev.value)
}

func TestLinkedList2_Insert_AfterHead(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 1}
	n2 := Node{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)

	n3 := Node{value: 5}
	l.Insert(l.head, n3)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 2, l.tail.value)
	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 5, l.head.next.value)
}

func TestLinkedList2_Insert_AfterTail(t *testing.T) {
	l := LinkedList2{}
	n1 := Node{value: 1}
	n2 := Node{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)

	n3 := Node{value: 5}
	l.Insert(l.tail, n3)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 5, l.tail.value)
	assert.Equal(t, 1, l.head.value)
}

func TestLinkedList2_Clean_EmptyList(t *testing.T) {
	l := LinkedList2{}

	l.Clean()

	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.tail)
	assert.Nil(t, l.head)
}

func TestLinkedList2_Clean_NotEmptyList(t *testing.T) {
	l := LinkedList2{}

	n1 := Node{value: 1}
	n2 := Node{value: 2}

	l.AddInTail(n1)
	l.AddInTail(n2)

	l.Clean()
	c := l.Count()

	assert.Equal(t, 0, c)
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.tail)
	assert.Nil(t, l.head)
}
