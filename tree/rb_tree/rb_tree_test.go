package rb_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"strings"
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
	rbt.Put(nil, nil)
	assert.Equal(t, 0, rbt.Size())
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

func TestRBTree_Size(t *testing.T) {
	var rbt RBTree
	assert.Zero(t, rbt.Size())
	rbt.Put(str("A"), "A")
	assert.Equal(t, 1, rbt.Size())
	putEachChar(&rbt, "ABCDE")
	assert.Equal(t, 5, rbt.Size())
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
	for i := len(s) - 1; i >= 0; i-- {
		rbt.Put(str(string(s[i])), "A")
		assert.Equal(t, str(string(s[i])), rbt.Min())
	}
}

func TestRBTree_Max(t *testing.T) {
	var rbt RBTree
	s := "ABCDEFGHI"
	for i := 0; i < len(s); i++ {
		rbt.Put(str(string(s[i])), "A")
		assert.Equal(t, str(string(s[i])), rbt.Max())
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

func TestRBTree_DeleteMax(t *testing.T) {
	var rbt RBTree
	s := "ABCDEFGHI"
	putEachChar(&rbt, s)
	for i := len(s) - 1; i >= 0; i-- {
		assert.False(t, rbt.IsEmpty())
		assert.Equal(t, str(string(s[i])), rbt.Max())
		rbt.DeleteMax()
	}
	assert.Nil(t, rbt.Max())
	assert.True(t, rbt.IsEmpty())
}

func TestRBTree_Delete(t *testing.T) {
	var rbt RBTree
	l := 100
	dict := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	mp := make(map[rune]bool)
	var b strings.Builder
	for i := 0; i < l; i++ {
		b.WriteByte(dict[rand.Intn(len(dict))])
	}
	s := b.String()

	for _, c := range s {
		rbt.Put(str(string(c)), true)
		mp[c] = true
	}

	for i, c := range s {
		k := str(string(c))
		val := rbt.Get(k)
		if (val == nil && mp[c]) || (val == true && !mp[c]) {
			t.Fatalf("Returned result differs from built-in map\nString: %s\nKey: %v\nIndex: %d\n", s, k, i)
		}
		mp[c] = false
		rbt.Delete(str(string(c)))
		val = rbt.Get(k)
		if (val == nil && mp[c]) || (val == true && !mp[c]) {
			t.Fatalf("Returned result differs from built-in map\nString: %s\nKey: %v\nIndex: %d\n", s, k, i)
		}
	}
}

func TestRBTree_Contains(t *testing.T) {
	var rbt RBTree
	putEachChar(&rbt, "ABCDE")
	assert.True(t, rbt.Contains(str("A")))
	assert.True(t, rbt.Contains(str("B")))
	assert.True(t, rbt.Contains(str("C")))
	assert.True(t, rbt.Contains(str("D")))
	assert.True(t, rbt.Contains(str("E")))
	assert.False(t, rbt.Contains(str("F")))
	assert.False(t, rbt.Contains(str("G")))
	assert.False(t, rbt.Contains(nil))
}

func TestRBTree_Keys(t *testing.T) {
	var rbt RBTree
	s := "9876543210DCBA"
	putEachChar(&rbt, s)
	ks := rbt.Keys()
	assert.Equal(t, len(s), len(ks))

	assert.Equal(t, []Key{
		str("0"),
		str("1"),
		str("2"),
		str("3"),
		str("4"),
		str("5"),
		str("6"),
		str("7"),
		str("8"),
		str("9"),
		str("A"),
		str("B"),
		str("C"),
		str("D"),
	}, ks)
}

func TestRBTree_Keys2(t *testing.T) {
	var rbt RBTree
	rbt.Put(str("A"), 1)
	rbt.Put(str("ABA"), 1)
	rbt.Put(str("BA"), 1)
	rbt.Put(str("AC"), 1)
	rbt.Put(str("AA"), 1)
	rbt.Put(str("ZZZZZ"), 1)
	rbt.Put(str("ZZZ"), 1)
	rbt.Put(str(""), 1)

	ks := rbt.Keys()
	assert.Equal(t, rbt.Size(), len(ks))

	assert.Equal(t, []Key{
		str(""),
		str("A"),
		str("AA"),
		str("ABA"),
		str("AC"),
		str("BA"),
		str("ZZZ"),
		str("ZZZZZ"),
	}, ks)
}

func TestRBTree_KeysBetween(t *testing.T) {
	var rbt RBTree
	s := "9876543210DCBA"
	putEachChar(&rbt, s)
	ks := rbt.KeysBetween(str("5"), str("B"))
	assert.Equal(t, 7, len(ks))

	assert.Equal(t, []Key{
		str("5"),
		str("6"),
		str("7"),
		str("8"),
		str("9"),
		str("A"),
		str("B"),
	}, ks)

	assert.Equal(t, 0, rbt.SizeBetween(nil, nil))
	assert.Equal(t, 0, rbt.SizeBetween(str("0"), nil))
	assert.Equal(t, 0, rbt.SizeBetween(nil, str("Z")))
}

func TestRBTree_KeysBetween2(t *testing.T) {
	rand.Seed(42)
	var rbt RBTree
	mp := make(map[string]string)
	dict := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	N := 1000
	L := 5

	for i := 0; i < N; i++ {
		var b strings.Builder
		for j := 0; j < L; j++ {
			b.WriteByte(dict[rand.Intn(len(dict))])
		}
		s := b.String()
		mp[s] = s
		rbt.Put(str(s), s)
	}

	keys := make([]string, 0)

	for k := range mp {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	assert.Greater(t, len(keys), 10)

	lb := keys[4]
	ub := keys[len(keys)-5]

	ks2 := rbt.KeysBetween(str(lb), str(ub))
	assert.Equal(t, len(ks2), len(keys)-8)

	for i, x := range ks2 {
		assert.Equal(t, str(keys[i+4]), x)
	}
}

func TestRBTree_Floor(t *testing.T) {
	var rbt RBTree
	s := "ACEGIKMOQSUWY"
	putEachChar(&rbt, s)

	for _, x := range s {
		assert.Equal(t, str(x), rbt.Floor(str(x)))
	}
	for _, x := range s {
		assert.Equal(t, str(x), rbt.Floor(str(x+1)))
	}
	assert.Nil(t, rbt.Floor(str(s[0]-1)))
	assert.Nil(t, rbt.Floor(nil))
}

func TestRBTree_Ceiling(t *testing.T) {
	var rbt RBTree
	s := "ACEGIKMOQSUWY"
	putEachChar(&rbt, s)

	for _, x := range s {
		assert.Equal(t, str(x), rbt.Ceiling(str(x)))
	}
	for _, x := range s {
		assert.Equal(t, str(x), rbt.Ceiling(str(x-1)))
	}
	assert.Nil(t, rbt.Ceiling(str(s[len(s)-1]+1)))
	assert.Nil(t, rbt.Floor(nil))
}

func TestRBTree_Select(t *testing.T) {
	var rbt RBTree
	s := "ABCDEFGHIJKLMN"
	putEachChar(&rbt, s)

	for i, x := range s {
		assert.Equal(t, str(string(x)), rbt.Select(i))
		assert.Equal(t, i, rbt.Rank(rbt.Select(i)))
	}
}

func TestRBTree_Rank(t *testing.T) {
	var rbt RBTree
	s := "ABCDEFGHIJKLMN"
	putEachChar(&rbt, s)

	for i, x := range s {
		assert.Equal(t, i, rbt.Rank(str(string(x))))
		assert.Equal(t, str(string(x)), rbt.Select(rbt.Rank(str(string(x)))))
	}
	assert.Equal(t, -1, rbt.Rank(nil))
}

func TestRBTree_SizeBetween(t *testing.T) {
	var rbt RBTree
	s := "ABCDEFGHIJKLMN"
	putEachChar(&rbt, s)

	assert.Equal(t, 0, rbt.SizeBetween(str("Z"), str("A")))
	assert.Equal(t, 3, rbt.SizeBetween(str("C"), str("E")))

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if j < i {
				assert.Equal(t, 0, rbt.SizeBetween(str(string(s[i])), str(string(s[j]))))
			} else {
				assert.Equal(t, j+1-i, rbt.SizeBetween(str(string(s[i])), str(string(s[j]))))
			}
		}
	}
	assert.Zero(t, rbt.SizeBetween(str("A"), nil))
	assert.Zero(t, rbt.SizeBetween(nil, nil))
	assert.Zero(t, rbt.SizeBetween(nil, str("Z")))
}

func TestRBTree(t *testing.T) {
	rand.Seed(42)

	var rbt RBTree

	N := 1000
	L := 5

	dict := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	mp := make(map[string]bool)

	for i := 0; i < N; i++ {
		l := len(dict)
		s := make([]uint8, L)
		for j := 0; j < L; j++ {
			p := rand.Intn(l)
			s[j] = dict[p]
		}
		rbt.Put(str(string(s)), string(s))
		mp[string(s)] = true
		if checkViolation(rbt.root) {
			fm := fmt.Sprintf("Red-Black Tree definition violated, i: %d, word: %s\n", i, s)
			assert.Failf(t, fm, "Should not break Red-Black Tree rules at any time")
			break
		}
	}

	var i = 0
	for k := range mp {
		rbt.Delete(str(k))
		if checkViolation(rbt.root) {
			fm := fmt.Sprintf("Red-Black Tree definition violated, i: %d, word: %s\n", i, k)
			assert.Failf(t, fm, "Should not break Red-Black Tree rules at any time")
			break
		}
		i++
	}

	rbt.Init()
	putEachChar(&rbt, "ABC")
	assert.False(t, checkViolation(rbt.root))
}

//func getColorInt(x *node) int {
//	if x.color == RED {
//		return 1
//	} else {
//		return 0
//	}
//}

//func printKeys(x *node) {
//	if x == nil {
//		fmt.Printf("[] ")
//		return
//	}
//	if x.left != nil || x.right != nil {
//		fmt.Printf("[%v,%v ", x.key, getColorInt(x))
//		printKeys(x.left)
//		printKeys(x.right)
//		fmt.Printf("] ")
//	} else {
//		fmt.Printf("[%v,%v] ", x.key, getColorInt(x))
//	}
//}

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
