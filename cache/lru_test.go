package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	cache := New(5)

	for i := 1; i <= 10; i++ {
		cache.Put(i, i)
	}

	li := cache.list

	assert.Equal(t, li.Len(), 5)

	for i := 0; i < 5; i++ {
		x := li.Front()
		assert.Equal(t, entry{key: 10 - i, value: 10 - i}, x.Value)
		li.Remove(x)
	}
}

func TestLRUCache_Get(t *testing.T) {
	cache := New(10)

	for i := 1; i <= 10; i++ {
		cache.Put(i, i)
	}

	x, hit := cache.Get(3)
	assert.True(t, hit)
	assert.Equal(t, 3, x)
	assert.Equal(t, 3, cache.list.Front().Value.(entry).value)
	x, hit = cache.Get(8)
	assert.True(t, hit)
	assert.Equal(t, 8, x)
	assert.Equal(t, 8, cache.list.Front().Value.(entry).value)
	x, hit = cache.Get(9)
	assert.True(t, hit)
	assert.Equal(t, 9, x)
	assert.Equal(t, 9, cache.list.Front().Value.(entry).value)
	x, hit = cache.Get(1)
	assert.True(t, hit)
	assert.Equal(t, 1, x)
	assert.Equal(t, 1, cache.list.Front().Value.(entry).value)
	x, hit = cache.Get(2)
	assert.True(t, hit)
	assert.Equal(t, 2, x)
	assert.Equal(t, 2, cache.list.Front().Value.(entry).value)
	x, hit = cache.Get(4)
	assert.True(t, hit)
	assert.Equal(t, 4, x)
	assert.Equal(t, 4, cache.list.Front().Value.(entry).value)
}
