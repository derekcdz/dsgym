package tree

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	height int
	key    Key
}

type Key interface {
	CompareTo(Key) int
}

type AvlTree struct {
	root *Node
	size int
}

type Tree AvlTree

func New() *Tree {
	t := Tree{
		root: nil,
		size: 0,
	}
	return &t
}

func (t *Tree) Init() *Tree {
	t.root = nil
	t.size = 0
	return t
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Add(k Key) {
	if t.root == nil {
		t.root = &Node{
			parent: nil,
			left:   nil,
			right:  nil,
			height: 1,
			key:    k,
		}
		t.size = 1
	} else {

		t.root = insert(t.root, k)
		t.size++
	}
}

func newNode(key Key) *Node {
	return &Node{
		parent: nil,
		left:   nil,
		right:  nil,
		height: 1,
		key:    key,
	}
}

func (t *Node) getHeight() int {
	if t == nil {
		return 0
	}
	return t.height
}

func (t *Node) calcHeight() int {
	if t == nil {
		return 0
	}
	if t.left.getHeight() > t.right.getHeight() {
		t.height = t.left.getHeight() + 1
	} else {
		t.height = t.right.getHeight() + 1
	}
	return t.height
}

func adjust(t *Node) *Node {
	if t == nil {
		return nil
	}

	if t.left.getHeight()-t.right.getHeight() > 1 {
		if t.left.getLeft().getHeight() < t.left.getRight().getHeight() {
			t.left = rotateLeft(t.left)
		}
		return rotateRight(t)
	}

	if t.right.getHeight()-t.left.getHeight() > 1 {
		if t.right.getLeft().getHeight() > t.right.getRight().getHeight() {
			t.right = rotateRight(t.right)
		}
		return rotateLeft(t)
	}

	t.calcHeight()
	return t
}

func (t *Node) getLeft() *Node {
	if t == nil {
		return nil
	}
	return t.left
}

func (t *Node) getRight() *Node {
	if t == nil {
		return nil
	}
	return t.right
}

func rotateLeft(t *Node) *Node {
	if t == nil {
		return nil
	}
	if t.left.getHeight() == t.right.getHeight() {
		return t
	}
	res := t.right
	t.right = res.getLeft()
	res.left = t

	res.getLeft().calcHeight()
	res.getRight().calcHeight()
	res.calcHeight()
	return res
}

func rotateRight(t *Node) *Node {
	if t == nil {
		return nil
	}
	if t.left.getHeight() == t.right.getHeight() {
		return t
	}
	res := t.left
	t.left = res.getRight()
	res.right = t

	res.getLeft().calcHeight()
	res.getRight().calcHeight()
	res.calcHeight()
	return res
}

func insert(t *Node, k Key) *Node {
	if t == nil {
		return newNode(k)
	}
	res := k.CompareTo(t.key)
	if res < 0 {
		t.left = insert(t.left, k)
	} else if res > 0 {
		t.right = insert(t.right, k)
	} else {
		return t // existing key not accepted
	}
	return adjust(t)
}

func (t *Tree) Apply(f func(Key)) {
	if t.root == nil {
		return
	}
	t.root.traverse(f)
}

// ToSlice returns a slice of stored values, which are sorted
func (t *Tree) ToSlice() []Key {
	res := make([]Key, 0, t.size)
	appendToSlice := func(k Key) {
		res = append(res, k)
	}
	t.Apply(appendToSlice)
	return res
}

func (t *Node) traverse(f func(Key)) {
	if t == nil {
		return
	}
	t.left.traverse(f)
	f(t.key)
	t.right.traverse(f)
}
