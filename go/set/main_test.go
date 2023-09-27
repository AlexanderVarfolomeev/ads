package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPowerSet_Put_Empty(t *testing.T) {
	set := Init[string]()

	set.Put("123")
	set.Put("1234")

	res := set.Get("123")
	assert.True(t, res)
	res1 := set.Get("1234")
	assert.True(t, res1)
}

func TestPowerSet_Remove_NotExist(t *testing.T) {
	set := Init[string]()

	set.Put("123")

	remove := set.Remove("1234")
	assert.False(t, remove)
}

func TestPowerSet_Remove_Exist(t *testing.T) {
	set := Init[string]()

	set.Put("123")

	remove := set.Remove("123")
	assert.True(t, remove)
}

func TestPowerSet_Intersection_Empty(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	set2.Put("4")
	set2.Put("5")
	set2.Put("6")

	intersection := set.Intersection(set2)
	assert.Equal(t, 0, intersection.Size())
}

func TestPowerSet_Intersection_NotEmpty(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	set2.Put("1")
	set2.Put("5")
	set2.Put("3")

	intersection := set.Intersection(set2)
	assert.Equal(t, 2, intersection.Size())

	assert.True(t, intersection.Get("1"))
	assert.True(t, intersection.Get("3"))
}

func TestPowerSet_Union_NotEmpty(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	set2.Put("1")
	set2.Put("5")
	set2.Put("3")

	intersection := set.Union(set2)
	assert.Equal(t, 4, intersection.Size())
	assert.True(t, intersection.Get("1"))
	assert.True(t, intersection.Get("2"))
	assert.True(t, intersection.Get("3"))
	assert.True(t, intersection.Get("5"))
}

func TestPowerSet_Union_Empty(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	intersection := set.Union(set2)
	assert.Equal(t, 3, intersection.Size())
	assert.True(t, intersection.Get("1"))
	assert.True(t, intersection.Get("2"))
	assert.True(t, intersection.Get("3"))
}

func TestPowerSet_Difference_NotEmpty(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	difference := set.Difference(set2)
	assert.Equal(t, 3, difference.Size())
	assert.True(t, difference.Get("1"))
	assert.True(t, difference.Get("2"))
	assert.True(t, difference.Get("3"))
}

func TestPowerSet_Difference_Empty(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	set2.Put("1")
	set2.Put("2")
	set2.Put("3")

	difference := set.Difference(set2)
	assert.Equal(t, 0, difference.Size())
}

func TestPowerSet_IsSubset_AllInSet1In(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	set2.Put("1")
	set2.Put("2")
	set2.Put("3")
	set2.Put("4")

	isSubset := set.IsSubset(set2)
	assert.False(t, isSubset)
}

func TestPowerSet_IsSubset_NotAllInSet2In(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	set2.Put("1")
	set2.Put("34")

	isSubset := set.IsSubset(set2)
	assert.False(t, isSubset)
}

func TestPowerSet_IsSubset_AllInSet2In(t *testing.T) {
	set := Init[string]()

	set.Put("1")
	set.Put("2")
	set.Put("3")

	set2 := Init[string]()

	set2.Put("1")
	set2.Put("2")
	set2.Put("3")

	isSubset := set.IsSubset(set2)
	assert.True(t, isSubset)
}

func TestPowerSet_IsFast(t *testing.T) {
	set1 := Init[int]()
	set2 := Init[int]()

	for i := 0; i < 20000; i++ {
		set1.Put(i)
		set2.Put(i + 1)
	}

	now := time.Now()
	set1.IsSubset(set2)
	res := time.Since(now)
	assert.Less(t, res, time.Second*2)

	now = time.Now()
	set1.Union(set2)
	res = time.Since(now)
	assert.Less(t, res, time.Second*2)

	now = time.Now()
	set1.Difference(set2)
	res = time.Since(now)
	assert.Less(t, res, time.Second*2)

	now = time.Now()
	set1.Intersection(set2)
	res = time.Since(now)
	assert.Less(t, res, time.Second*2)
}
