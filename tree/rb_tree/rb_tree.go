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
	KeysBetween() []Key
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
	if x.color == BLACK {
		x.color = RED
		x.left.color = BLACK
		x.right.color = BLACK
	} else {
		x.color = BLACK
		x.left.color = RED
		x.right.color = RED
	}
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
	if nx.right.isRed() {
		nx = nx.rotateLeft()
	}
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

func (t *RBTree) Init() {
	t.root = nil
}

func (t *RBTree) Get(k Key) Value {
	x := t.root.find(k)
	if x == nil {
		return nil
	}
	return x.value
}

func (t *RBTree) Put(k Key, v Value) {
	if t.root == nil {
		t.root = newNode(k, v, BLACK)
	} else {
		t.root = t.root.insert(k, v)
		t.root.color = BLACK
	}
}

func (t *RBTree) IsEmpty() bool {
	return t.Size() == 0
}

func (t *RBTree) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.size
}

func (t *RBTree) Min() Key {
	min := t.root.findMin()
	if min == nil {
		return nil
	}
	return min.key
}

func (t *RBTree) Max() Key {
	max := t.root.findMax()
	if max == nil {
		return nil
	}
	return max.key
}

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
