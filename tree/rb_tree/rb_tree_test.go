package rb_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type str string

func (s str) CompareTo(k Key) int {
	if s < k.(str) {
		return -1
	} else if s > k.(str) {
		return 1
	}
	return 0
}

func TestRBTree_Init(t *testing.T) {
	var rbt RBTree
	rbt.Put(str("A"), "A")
	rbt.Init()
	assert.Zero(t, rbt.Size())
	assert.Nil(t, rbt.root)
}

func TestRBTree_Put(t *testing.T) {
	var rbt RBTree
	rbt.Put(str("A"), "A")
	rbt.Put(str("X"), "X")
	assert.Equal(t, 2, rbt.Size())
}

func TestRBTree_Put2(t *testing.T) {
	var rbt RBTree
	putEachChar(&rbt, "AAAAAAA")
	assert.Equal(t, 1, rbt.Size())
	putEachChar(&rbt, "XYZ")
	assert.Equal(t, 4, rbt.Size())
}

func TestRBTree_Get(t *testing.T) {
	var rbt RBTree
	putEachChar(&rbt, "AEIOU")
	assert.Equal(t, "A", rbt.Get(str("A")))
	assert.Equal(t, "E", rbt.Get(str("E")))
	assert.Equal(t, "I", rbt.Get(str("I")))
	assert.Equal(t, "O", rbt.Get(str("O")))
	assert.Equal(t, "U", rbt.Get(str("U")))
	assert.Nil(t, rbt.Get(str("a")))
	assert.Nil(t, rbt.Get(str("z")))
	assert.Nil(t, rbt.Get(str("Z")))
}

func TestRBTree_IsEmpty(t *testing.T) {
	var rbt RBTree
	assert.True(t, rbt.IsEmpty())
	rbt.Put(str("A"), "A")
	assert.Equal(t, 1, rbt.Size())
	assert.False(t, rbt.IsEmpty())
}

func TestRBTree_Min(t *testing.T) {
	var rbt RBTree
	s := "ABCDEFGHI"
	for i := len(s) - 1; i > 0; i-- {
		rbt.Put(str(string(s[i])), "A")
		assert.Equal(t, str(string(s[i])), rbt.Min())
	}
}

func TestRBTree_DeleteMin(t *testing.T) {
	var rbt RBTree
	s := "ABCDEFGHI"
	putEachChar(&rbt, s)
	for _, c := range s {
		assert.False(t, rbt.IsEmpty())
		assert.Equal(t, str(string(c)), rbt.Min())
		rbt.DeleteMin()
	}
	assert.Nil(t, rbt.Min())
	assert.True(t, rbt.IsEmpty())
}

func TestRBTree(t *testing.T) {
	rand.Seed(42)

	println()
	var rbt RBTree

	N := 1000
	L := 5

	dict := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := 0; i < N; i++ {
		l := len(dict)
		s := make([]uint8, L)
		for j := 0; j < L; j++ {
			p := rand.Intn(l)
			s[j] = dict[p]
		}
		rbt.Put(str(string(s)), string(s))
		//fmt.Printf("Insert: %v\n", string(dict[j]))
		if checkViolation(rbt.root) {
			fm := fmt.Sprintf("Red-Black Tree definition violated, i: %d, word: %s\n", i, s)
			assert.Failf(t, fm, "Should not break Red-Black Tree rules at any time")
			break
		}
		//printKeys(rbt.root)
	}
	//bh := calcBH(rbt.root)
	//fmt.Printf("%d\n", bh)
	//printKeys(rbt.root)

	rbt.Init()
	putEachChar(&rbt, "ABC")
	//bh = calcBH(rbt.root)
	//fmt.Printf("%d\n", bh)
	assert.False(t, checkViolation(rbt.root))
	//printKeys(rbt.root)
	println()
}

func getColorInt(x *node) int {
	if x.color == RED {
		return 1
	} else {
		return 0
	}
}

func printKeys(x *node) {
	if x == nil {
		fmt.Printf("[] ")
		return
	}
	if x.left != nil || x.right != nil {
		fmt.Printf("[%v,%v ", x.key, getColorInt(x))
		printKeys(x.left)
		printKeys(x.right)
		fmt.Printf("] ")
	} else {
		fmt.Printf("[%v,%v] ", x.key, getColorInt(x))
	}
}

func putEachChar(tree *RBTree, s string) {
	for _, x := range s {
		tree.Put(str(string(x)), string(x))
	}
}

func calcBH(x *node) int {
	if x == nil {
		return 0
	}

	bh := 0

	if !x.isRed() {
		bh = 1
	}

	bhl := calcBH(x.left)
	bhr := calcBH(x.right)
	if x.left == nil {
		bhl = bhr
	}
	if x.right == nil {
		bhr = bhl
	}
	if bhl != bhr || bhl == -1 {
		bh = -1
	} else {
		bh += bhl
	}
	//fmt.Printf("Key: %v, bhl: %v, bhr: %v, bh: %v\n", x.key, bhl, bhr, bhr)
	return bh
}

func checkViolation(x *node) bool {
	if calcBH(x) == -1 {
		return true
	}

	return false
}
