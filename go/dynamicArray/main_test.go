package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynArray_Append(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	for i := 0; i < 16; i++ {
		val, err := a.GetItem(i)
		assert.NoError(t, err)
		assert.Equal(t, i, val)
	}

	assert.Equal(t, 16, a.count)
	assert.Equal(t, 16, a.capacity)
}

func TestDynArray_Append_WithIncrease(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 32; i++ {
		a.Append(i)
	}

	for i := 0; i < 32; i++ {
		val, err := a.GetItem(i)
		assert.NoError(t, err)
		assert.Equal(t, i, val)
	}
	assert.Equal(t, 32, a.count)
	assert.Equal(t, 32, a.capacity)
}

func TestDynArray_Append_WithIncrease_NotFull(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 25; i++ {
		a.Append(i)
	}

	for i := 0; i < a.count; i++ {
		val, err := a.GetItem(i)
		assert.NoError(t, err)
		assert.Equal(t, i, val)
	}

	_, err := a.GetItem(27)
	assert.Error(t, err)
	assert.Equal(t, 25, a.count)
	assert.Equal(t, 32, a.capacity)
}

func TestDynArray_GetItem_OutOfRange(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	_, err := a.GetItem(27)
	assert.Error(t, err)
}

func TestDynArray_GetItem_OutOfRange_LessThenZero(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	_, err := a.GetItem(-15)
	assert.Error(t, err)
}

func TestDynArray_Insert_FirstElement(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 8; i++ {
		a.Append(i)
	}

	err := a.Insert(9, 0)
	assert.NoError(t, err)
	val, err := a.GetItem(0)
	assert.NoError(t, err)
	assert.Equal(t, 9, val)
	for i := 1; i < 9; i++ {
		val, err := a.GetItem(i)

		assert.NoError(t, err)
		assert.Equal(t, i-1, val)
	}
}

func TestDynArray_Insert_MiddleElement(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 8; i++ {
		a.Append(i)
	}

	err := a.Insert(9, 3)
	assert.NoError(t, err)

	for i := 0; i < 9; i++ {
		val, err := a.GetItem(i)
		assert.NoError(t, err)

		if i == 3 {
			assert.Equal(t, 9, val)
		} else if i < 3 {
			assert.Equal(t, i, val)
		} else {
			assert.Equal(t, i-1, val)
		}
	}
}

func TestDynArray_Insert_LastElement_WithIncrease(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	err := a.Insert(16, a.count)
	assert.NoError(t, err)
	assert.Equal(t, 32, a.capacity)
	assert.Equal(t, 17, a.count)
	for i := 0; i < a.count; i++ {
		val, err := a.GetItem(i)

		assert.NoError(t, err)
		assert.Equal(t, i, val)
	}
}

func TestDynArray_Insert_FirstElement_WithIncrease(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	err := a.Insert(-5, 0)
	assert.NoError(t, err)
	assert.Equal(t, 32, a.capacity)
	assert.Equal(t, 17, a.count)

	val, err := a.GetItem(0)
	assert.NoError(t, err)
	assert.Equal(t, -5, val)

	for i := 1; i < a.count; i++ {
		val, err := a.GetItem(i)

		assert.NoError(t, err)
		assert.Equal(t, i-1, val)
	}
}

func TestDynArray_Insert_LastElement(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 9; i++ {
		a.Append(i)
	}

	err := a.Insert(9, a.count)
	assert.NoError(t, err)

	for i := 1; i < 9; i++ {
		val, err := a.GetItem(i)

		assert.NoError(t, err)
		assert.Equal(t, i, val)
	}
}

func TestDynArray_Insert_OutOfRange(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 9; i++ {
		a.Append(i)
	}

	err := a.Insert(9, a.count+1)
	assert.Error(t, err)
}

func TestDynArray_Insert_LessThenZero(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 9; i++ {
		a.Append(i)
	}

	err := a.Insert(9, -1)
	assert.Error(t, err)
}

func TestDynArray_Remove_First(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	err := a.Remove(0)
	assert.NoError(t, err)
	assert.Equal(t, 15, a.count)
	assert.Equal(t, 16, a.capacity)

	for i := 0; i < a.count; i++ {
		val, err := a.GetItem(i)

		assert.NoError(t, err)
		assert.Equal(t, i+1, val)
	}
}

func TestDynArray_Remove_Last(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	err := a.Remove(a.count - 1)
	assert.NoError(t, err)
	assert.Equal(t, 15, a.count)
	assert.Equal(t, 16, a.capacity)

	for i := 0; i < a.count; i++ {
		val, err := a.GetItem(i)

		assert.NoError(t, err)
		assert.Equal(t, i, val)
	}
}

func TestDynArray_Remove_NoDecrease(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 32; i++ {
		a.Append(i)
	}

	assert.Equal(t, 32, a.count)
	assert.Equal(t, 32, a.capacity)

	for i := 0; i < 16; i++ {
		err := a.Remove(i)
		assert.NoError(t, err)
	}

	assert.Equal(t, 16, a.count)
	assert.Equal(t, 32, a.capacity)
}

func TestDynArray_Remove_Decrease(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 32; i++ {
		a.Append(i)
	}

	assert.Equal(t, 32, a.count)
	assert.Equal(t, 32, a.capacity)

	for i := 0; i < 17; i++ {
		err := a.Remove(0)
		assert.NoError(t, err)
	}

	assert.Equal(t, 15, a.count)
	assert.Equal(t, 21, a.capacity)

	for i := 0; i < 7; i++ {
		err := a.Remove(0)
		assert.NoError(t, err)
	}

	assert.Equal(t, 8, a.count)
	assert.Equal(t, 16, a.capacity)
}

func TestDynArray_IncreaseAfterDecrease(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 32; i++ {
		a.Append(i)
	}

	assert.Equal(t, 32, a.count)
	assert.Equal(t, 32, a.capacity)

	for i := 0; i < 17; i++ {
		err := a.Remove(0)
		assert.NoError(t, err)
	}

	assert.Equal(t, 15, a.count)
	assert.Equal(t, 21, a.capacity)

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	assert.Equal(t, 31, a.count)
	assert.Equal(t, 42, a.capacity)
}

func TestDynArray_Remove_Middle(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	err := a.Remove(5)
	assert.NoError(t, err)
	assert.Equal(t, 15, a.count)
	assert.Equal(t, 16, a.capacity)

	for i := 0; i < a.count; i++ {
		val, err := a.GetItem(i)
		assert.NoError(t, err)

		if i < 5 {
			assert.Equal(t, i, val)
		} else {
			assert.Equal(t, i+1, val)
		}
	}
}

func TestDynArray_Remove_OutOfRange(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	err := a.Remove(21)
	assert.Error(t, err)
}

func TestDynArray_Remove_LessThenZero(t *testing.T) {
	a := DynArray[int]{}
	a.Init()

	for i := 0; i < 16; i++ {
		a.Append(i)
	}

	err := a.Remove(-1)
	assert.Error(t, err)
}
