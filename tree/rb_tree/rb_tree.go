// A map, as known as key-value tables, implemented with red-black tree
// complexity of inserting, searching and deleting are all O(logN)
package rb_tree

type Key interface {
	CompareTo(Key) int
}

type Value interface{}

type Entry struct {
	Key   Key
	Value Value
}

// RBTree should implements SortedMap
type SortedMap interface {
	Init()
	Get(Key) Value
	Put(Key, Value)
	Delete(Key)
	Contains(Key) bool
	IsEmpty() bool
	Size() int
	Min() Key
	Max() Key
	Floor(Key) Key
	Ceiling(Key) Key
	Select(int) Key
	Rank(Key) int
	DeleteMin()
	DeleteMax()
	SizeBetween(Key, Key) int
	Keys() []Key
	KeysBetween(Key, Key) []Key
}

const (
	RED   = true
	BLACK = false
)

type color bool

type node struct {
	color color
	size  int
	left  *node
	right *node
	key   Key
	value Value
}

type RBTree struct {
	root *node
}

func (x *node) isRed() bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

func (x *node) flipColors() {
	x.color = !x.color
	x.left.color = !x.left.color
	x.right.color = !x.right.color
}

func newNode(k Key, v Value, c color) *node {
	return &node{
		color: c,
		size:  1,
		left:  nil,
		right: nil,
		key:   k,
		value: v,
	}
}

func (x *node) getLeft() *node {
	if x == nil {
		return nil
	}
	return x.left
}

func (x *node) getRight() *node {
	if x == nil {
		return nil
	}
	return x.right
}

func (x *node) getSize() int {
	if x == nil {
		return 0
	}
	return x.size
}

func (x *node) calcSize() {
	if x == nil {
		return
	}
	x.size = x.left.getSize() + x.right.getSize() + 1
}

func (x *node) rotateLeft() *node {
	if x == nil {
		return nil
	}
	res := x.right
	x.right = res.left
	res.left = x
	res.color, x.color = x.color, res.color

	res.size = x.size
	x.calcSize()
	return res
}

func (x *node) rotateRight() *node {
	if x == nil {
		return nil
	}
	res := x.left
	x.left = res.right
	res.right = x
	res.color, x.color = x.color, res.color

	res.size = x.size
	x.calcSize()
	return res
}

func (x *node) insert(k Key, v Value) *node {
	if x == nil {
		return newNode(k, v, RED)
	} else {
		cmp := k.CompareTo(x.key)
		if cmp < 0 {
			x.left = x.left.insert(k, v)
		} else if cmp > 0 {
			x.right = x.right.insert(k, v)
		} else {
			x.value = v
		}

		res := x
		if res.right.isRed() && !res.left.isRed() {
			res = res.rotateLeft()
		}
		if res.left.isRed() && res.left.getLeft().isRed() {
			res = res.rotateRight()
		}
		if res.left.isRed() && res.right.isRed() {
			res.flipColors()
		}
		res.calcSize()
		return res
	}
}

func (x *node) find(k Key) *node {
	if x == nil {
		return nil
	}
	cmp := k.CompareTo(x.key)
	if cmp < 0 {
		return x.left.find(k)
	} else if cmp > 0 {
		return x.right.find(k)
	} else {
		return x
	}
}

func (x *node) findMin() *node {
	if x.getLeft() == nil {
		return x
	}
	return x.left.findMin()
}

func (x *node) findMax() *node {
	if x.getRight() == nil {
		return x
	}
	return x.right.findMax()
}

func (x *node) balance() *node {
	nx := x
	if !nx.left.isRed() && nx.right.isRed() {
		nx = nx.rotateLeft()
	}
	if nx.left.isRed() && nx.left.getLeft().isRed() {
		nx = nx.rotateRight()
	}
	if nx.left.isRed() && nx.right.isRed() {
		nx.flipColors()
	}
	nx.calcSize()
	return nx
}

// x is red and both x.left and x.left.left are black
func (x *node) moveRedLeft() *node {
	nx := x
	nx.flipColors()

	if nx.getRight().getLeft().isRed() {
		nx.right = nx.getRight().rotateRight()
		nx = nx.rotateLeft()
		nx.flipColors()
	}
	return nx
}

func (x *node) deleteMin() *node {
	if x.left == nil {
		return nil
	}
	nx := x
	if !nx.left.isRed() && !nx.left.getLeft().isRed() {
		nx = nx.moveRedLeft()
	}
	nx.left = nx.left.deleteMin()
	return nx.balance()
}

// x is red and both x.right and x.right.left are black
func (x *node) moveRedRight() *node {
	nx := x
	nx.flipColors()
	if nx.left.getLeft().isRed() {
		nx = nx.rotateRight()
		nx.flipColors()
	}
	return nx
}

func (x *node) deleteMax() *node {
	nx := x
	if nx.left.isRed() {
		nx = nx.rotateRight()
	}
	if nx.right == nil {
		return nil
	}
	if !nx.right.isRed() && !nx.right.getLeft().isRed() {
		nx = nx.moveRedRight()
	}
	nx.right = nx.right.deleteMax()
	return nx.balance()
}

func (x *node) delete(k Key) *node {
	if x == nil {
		return nil
	}
	nx := x
	if k.CompareTo(nx.key) < 0 {
		if nx.left != nil && !nx.left.isRed() && !nx.left.getLeft().isRed() {
			nx = nx.moveRedLeft()
		}
		nx.left = nx.left.delete(k)
	} else {
		if nx.left.isRed() {
			// keep current node is the right key of a 3-node
			// because of the rotation, the new root's key is less than k
			nx = nx.rotateRight()
			nx.right = nx.right.delete(k)
		} else {
			// here nx is the right key of a 3-node, which also means nx.left.color == BLACK
			// when it has no right child, it must not have left child neither, we can safely delete it
			if nx.right == nil && k.CompareTo(nx.key) == 0 {
				return nil
			}
			if nx.right != nil && !nx.right.isRed() && !nx.right.getLeft().isRed() {
				nx = nx.moveRedRight()
			}
			if k.CompareTo(nx.key) == 0 {
				rmin := nx.right.findMin()
				nx.key = rmin.key
				nx.value = rmin.value
				nx.right = nx.right.deleteMin()
			} else {
				nx.right = nx.right.delete(k)
			}
		}
	}
	return nx.balance()
}

func (x *node) floor(k Key) *node {
	if x == nil {
		return nil
	}
	cmp := x.key.CompareTo(k)
	if cmp <= 0 {
		r := x.right.floor(k)
		if r != nil {
			return r
		}
		return x
	} else {
		return x.left.floor(k)
	}
}

func (x *node) ceiling(k Key) *node {
	if x == nil {
		return nil
	}
	cmp := x.key.CompareTo(k)
	if cmp >= 0 {
		l := x.left.ceiling(k)
		if l != nil {
			return l
		}
		return x
	} else {
		return x.right.ceiling(k)
	}
}

func (x *node) getNth(n int) *node {
	if x == nil {
		return nil
	}
	rank := 0
	if x.left != nil {
		rank += x.left.size
	}
	if rank == n {
		return x
	}
	if rank > n {
		return x.left.getNth(n)
	} else {
		return x.right.getNth(n - rank - 1)
	}
}

func (x *node) rank(k Key, acc int) int {
	if x == nil {
		return -1
	}
	cmp := k.CompareTo(x.key)
	if cmp < 0 {
		return x.left.rank(k, acc)
	} else {
		if x.left != nil {
			acc += x.left.size
		}
		if cmp == 0 {
			return acc
		}
		return x.right.rank(k, acc+1)
	}
}

// sizeLess returns number of Key which are strictly less than k
func (x *node) sizeLess(k Key, acc int) int {
	if x == nil {
		return 0
	}
	cmp := k.CompareTo(x.key)
	if cmp < 0 {
		return x.left.sizeLess(k, acc)
	} else {
		if x.left != nil {
			acc += x.left.size
		}
		if cmp > 0 {
			return x.right.sizeLess(k, acc+1)
		}
		return acc
	}
}

// sizeGreater returns number of Key which are strictly greater than k
func (x *node) sizeGreater(k Key, acc int) int {
	if x == nil {
		return 0
	}
	cmp := k.CompareTo(x.key)
	if cmp > 0 {
		return x.right.sizeGreater(k, acc)
	} else {
		if x.right != nil {
			acc += x.right.size
		}
		if cmp < 0 {
			return x.left.sizeGreater(k, acc+1)
		}
		return acc
	}
}

// Init initializes the tree, it deletes all keys from the tree
func (t *RBTree) Init() {
	t.root = nil
}

// Given a Key k, Get will return the Value associated with k
// When k is nil, a nli Value will be returned
func (t *RBTree) Get(k Key) Value {
	if k == nil {
		return nil
	}
	x := t.root.find(k)
	if x == nil {
		return nil
	}
	return x.value
}

// Put will store v in the tree associated with Key k
// When k is nil, the method will directly return and no values will be stored
func (t *RBTree) Put(k Key, v Value) {
	if k == nil { // refuse nil key
		return
	}
	if t.root == nil {
		t.root = newNode(k, v, BLACK)
	} else {
		t.root = t.root.insert(k, v)
		t.root.color = BLACK
	}
}

// Delete deletes the Value associated with Key k from the tree if the key exists
func (t *RBTree) Delete(k Key) {
	if k == nil {
		return
	}
	if t.root == nil {
		return
	}
	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = RED
	}
	t.root = t.root.delete(k)
	if t.root != nil {
		t.root.color = BLACK
	}
}

// Contains returns whether a Value associated with Key k is stored in the tree
func (t *RBTree) Contains(k Key) bool {
	if k == nil {
		return false
	}
	return t.root.find(k) != nil
}

// IsEmpty returns whether the tree is empty
func (t *RBTree) IsEmpty() bool {
	return t.Size() == 0
}

// Size returns the number of Key in the tree
func (t *RBTree) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.size
}

// Min returns the minimum key of the tree
func (t *RBTree) Min() Key {
	min := t.root.findMin()
	if min == nil {
		return nil
	}
	return min.key
}

// Min returns the minimum key of the tree
func (t *RBTree) Max() Key {
	max := t.root.findMax()
	if max == nil {
		return nil
	}
	return max.key
}

// Floor returns the maximum key which is less than or equals k
func (t *RBTree) Floor(k Key) Key {
	if k == nil {
		return nil
	}
	x := t.root.floor(k)
	if x == nil {
		return nil
	}
	return x.key
}

// Floor returns the minimum key which is greater than or equals k
func (t *RBTree) Ceiling(k Key) Key {
	if k == nil {
		return nil
	}
	x := t.root.ceiling(k)
	if x == nil {
		return nil
	}
	return x.key
}

// Select returns the key which ranks r-th (staring from 0, in order of Key's comparator) in the tree
// nil is returned when r < 0 or r >= t.Size()
// e.g. Select(0) returns the smallest Key
func (t *RBTree) Select(r int) Key {
	if r < 0 || r >= t.Size() {
		return nil
	}
	return t.root.getNth(r).key
}

// If Key k is in the tree, the rank of k is returned (staring from 0, in order of Key's comparator)
// otherwise -1 is returned
func (t *RBTree) Rank(k Key) int {
	if k == nil {
		return -1
	}
	return t.root.rank(k, 0)
}

// DeleteMin deletes the minimum key from the tree
func (t *RBTree) DeleteMin() {
	if t.root == nil {
		return
	}
	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = RED
	}
	t.root = t.root.deleteMin()
	if t.root != nil {
		t.root.color = BLACK
	}
}

// DeleteMin deletes the maximum key from the tree
func (t *RBTree) DeleteMax() {
	if t.root == nil {
		return
	}
	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = RED
	}
	t.root = t.root.deleteMax()
	if t.root != nil {
		t.root.color = BLACK
	}
}

// SizeBetween returns the number of Key that is in interval [lb, ub] (both sides included)
func (t *RBTree) SizeBetween(lb, ub Key) int {
	if lb == nil || ub == nil {
		return 0
	}
	if lb.CompareTo(ub) > 0 {
		return 0
	}
	return t.Size() - t.root.sizeLess(lb, 0) - t.root.sizeGreater(ub, 0)
}

// Keys returns a slice of Key which including all keys in the tree
// the returned slice is sorted
func (t *RBTree) Keys() []Key {
	keys := make([]Key, 0, t.Size())

	var getKeys func(x *node)
	getKeys = func(x *node) {
		if x == nil {
			return
		}
		getKeys(x.left)
		keys = append(keys, x.key)
		getKeys(x.right)
	}
	getKeys(t.root)

	return keys
}

// KeysBetween returns a sorted slice of Key, for each element k, k >= lb and k <= ub hold
// if lb == nil or ub == nil, empty slice is returned
func (t *RBTree) KeysBetween(lb, ub Key) []Key {
	if lb == nil || ub == nil {
		return []Key{}
	}
	keys := make([]Key, 0)
	if lb.CompareTo(ub) > 0 {
		return keys
	}

	var getKeys func(x *node)
	getKeys = func(x *node) {
		if x == nil {
			return
		}
		if x.key.CompareTo(lb) >= 0 {
			getKeys(x.left)
			if x.key.CompareTo(ub) <= 0 {
				// x.key is in [lb, ub]
				keys = append(keys, x.key)
			} else {
				return
			}
		}
		getKeys(x.right)
	}

	getKeys(t.root)
	return keys
}
