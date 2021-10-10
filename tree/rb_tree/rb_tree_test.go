package rb_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.Zero(t, rbt.size)
	assert.Nil(t, rbt.root)
}

func TestRBTree_Put(t *testing.T) {
	var rbt RBTree
	rbt.Put(str("A"), "A")
	rbt.Put(str("X"), "X")
	assert.Equal(t, 2, rbt.size)
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

func TestRBTree(t *testing.T) {
	println()
	var rbt RBTree

	putEachChar(&rbt, "ABCDEFGHIJKLMN")

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
