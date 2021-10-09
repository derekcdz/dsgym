package tree

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type intKey int

func (a intKey) CompareTo(b Key) int {
	if a < b.(intKey) {
		return -1
	} else if a > b.(intKey) {
		return 1
	} else {
		return 0
	}
}

func TestNew(t *testing.T) {
	tree := New()
	assert.NotNil(t, tree)
	assert.Nil(t, tree.root)
	assert.Equal(t, 0, tree.size)
}

func TestAvlTree_Add(t *testing.T) {
	tree := New()
	tree.Add(intKey(1))
	assert.Equal(t, intKey(1), tree.root.key)
	assert.Equal(t, 1, tree.root.height)
	tree.Add(intKey(3))
	tree.Add(intKey(2))
	tree.Add(intKey(1))
	assert.Equal(t, 3, tree.size)
	assert.Equal(t, 2, tree.root.height)
}

func TestAvlTree_Remove(t *testing.T) {
	var tree AvlTree
	sl1 := []int{0, 1, 2, 3, 4, 5, 6, -1, -2, -3}
	sl2 := make([]Key, len(sl1))
	for _, x := range sl1 {
		tree.Add(intKey(x))
	}
	sort.Ints(sl1)
	for i, x := range sl1 {
		sl2[i] = intKey(x)
	}
	assert.Equal(t, 10, tree.Size())
	assert.Equal(t, true, tree.Remove(intKey(0)))
	assert.Equal(t, true, tree.Remove(intKey(-1)))
	assert.Equal(t, true, tree.Remove(intKey(-2)))
	assert.Equal(t, true, tree.Remove(intKey(-3)))
	assert.Equal(t, false, tree.Remove(intKey(0)))
	assert.Equal(t, false, tree.Remove(intKey(-1)))
	assert.Equal(t, false, tree.Remove(intKey(-2)))
	assert.Equal(t, false, tree.Remove(intKey(-3)))
	assert.Equal(t, 6, tree.Size())
	assert.Equal(t, sl2[4:], tree.ToSlice())
	assert.Equal(t, true, tree.Remove(intKey(1)))
	assert.Equal(t, true, tree.Remove(intKey(2)))
	assert.Equal(t, true, tree.Remove(intKey(3)))
	assert.Equal(t, true, tree.Remove(intKey(4)))
	assert.Equal(t, true, tree.Remove(intKey(5)))
	assert.Equal(t, true, tree.Remove(intKey(6)))
	assert.Equal(t, 0, tree.Size())
	assert.Nil(t, tree.root)
}

func TestAvlTree_Min(t *testing.T) {
	var tree AvlTree
	sl1 := []int{0, 1, 2, 3, 4, 5, 6, -1, -2, -3}
	for _, x := range sl1 {
		tree.Add(intKey(x))
	}
	assert.Equal(t, intKey(-3), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(-2), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(-1), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(0), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(1), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(2), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(3), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(4), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(5), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Equal(t, intKey(6), tree.Min())
	assert.True(t, tree.Remove(tree.Min()))
	assert.Nil(t, tree.Min())
}

func TestAvlTree_Max(t *testing.T) {
	var tree AvlTree
	sl1 := []int{0, 1, 2, 3, 4, 5, 6, -1, -2, -3}
	for _, x := range sl1 {
		tree.Add(intKey(x))
	}

	assert.Equal(t, intKey(6), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(5), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(4), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(3), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(2), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(1), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(0), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(-1), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(-2), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Equal(t, intKey(-3), tree.Max())
	assert.True(t, tree.Remove(tree.Max()))
	assert.Nil(t, tree.Max())
}

func TestAvlTree_Init(t *testing.T) {
	tree := New()
	tree.Add(intKey(1))
	tree.Add(intKey(2))
	tree.Add(intKey(3))
	tree.Add(intKey(4))
	tree.Init()
	assert.NotNil(t, tree)
	assert.Nil(t, tree.root)
	assert.Equal(t, 0, tree.size)
}

func TestAvlTree_Init2(t *testing.T) {
	var tree AvlTree
	assert.Nil(t, tree.root)
	assert.Zero(t, tree.size)
	tree.Add(intKey(2))
	assert.NotNil(t, tree.root)
	tree.Add(intKey(1))
	tree.Add(intKey(3))
	assert.Equal(t, 3, tree.Size())
}

func TestAvlTree_Size(t *testing.T) {
	tree := New()
	assert.Equal(t, 0, tree.Size())
	tree.Add(intKey(1))
	assert.Equal(t, 1, tree.Size())
	assert.Equal(t, tree.size, tree.Size())
	tree.Add(intKey(2))
	tree.Add(intKey(3))
	tree.Add(intKey(4))
	assert.Equal(t, 4, tree.Size())
	assert.Equal(t, tree.size, tree.Size())
	tree.Add(intKey(0))
	tree.Add(intKey(2))
	tree.Add(intKey(3))
	tree.Add(intKey(4))
	tree.Add(intKey(0))
	tree.Add(intKey(4))
	assert.Equal(t, 5, tree.Size())
	assert.Equal(t, tree.size, tree.Size())
}

func TestAvlTree_Contains(t *testing.T) {
	var tree AvlTree
	tree.Add(intKey(0))
	tree.Add(intKey(1))
	tree.Add(intKey(2))
	tree.Add(intKey(3))
	tree.Add(intKey(4))
	tree.Add(intKey(5))
	assert.True(t, tree.Contains(intKey(0)))
	assert.True(t, tree.Contains(intKey(1)))
	assert.True(t, tree.Contains(intKey(2)))
	assert.True(t, tree.Contains(intKey(3)))
	assert.True(t, tree.Contains(intKey(4)))
	assert.True(t, tree.Contains(intKey(5)))
	assert.False(t, tree.Contains(intKey(10)))
	assert.False(t, tree.Contains(intKey(100)))
	assert.True(t, tree.Remove(intKey(0)))
	assert.False(t, tree.Contains(intKey(0)))
}

func TestAvlTree_ToSlice(t *testing.T) {
	tree := New()
	sl1 := []int{0, 1, 2, 3, 4, 5, 6}
	sl2 := make([]Key, 7)
	for i, x := range sl1 {
		tree.Add(intKey(x))
		sl2[i] = intKey(x)
	}
	sl := tree.ToSlice()
	assert.Len(t, sl, 7)
	assert.Equal(t, sl2, sl)
}

func TestAvlTree(t *testing.T) {
	var tree AvlTree
	tree.Add(intKey(1))
	tree.Add(intKey(2))
	tree.Add(intKey(3))
	sl := tree.ToSlice()

	assert.Equal(t, []Key{intKey(1), intKey(2), intKey(3)}, sl)

	tree.Add(intKey(-1))
	tree.Add(intKey(-2))
	tree.Add(intKey(-3))
	sl = tree.ToSlice()
	assert.Equal(t, []Key{intKey(-3), intKey(-2), intKey(-1), intKey(1), intKey(2), intKey(3)}, sl)

	rand.Seed(time.Now().UnixNano())

	randNums := make([]int, 30)
	tree.Init()

	for i := 0; i < 30; i++ {
		randNums[i] = rand.Int()
		tree.Add(intKey(randNums[i]))
	}

	sort.Ints(randNums)

	sl = tree.ToSlice()
	keys := make([]Key, 30)
	for i := 0; i < 30; i++ {
		keys[i] = intKey(randNums[i])
	}
	assert.Equal(t, keys, sl)

	tree.Init()

}
