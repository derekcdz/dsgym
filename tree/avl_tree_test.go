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
	return int(a - b.(intKey))
}

func TestAvlTree(t *testing.T) {
	assert.True(t, true, "True is true!")
	//fmt.Printf("%d", intKey(1).CompareTo(intKey(3)))
	tree := New()
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
	tree = New()

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

}
