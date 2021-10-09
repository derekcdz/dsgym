package avl_tree

type Node struct {
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

func New() *AvlTree {
	t := AvlTree{
		root: nil,
		size: 0,
	}
	return &t
}

func (t *AvlTree) Init() *AvlTree {
	t.root = nil
	t.size = 0
	return t
}

func (t *AvlTree) Size() int {
	return t.size
}

func (t *AvlTree) Add(k Key) bool {
	if t.root == nil {
		t.root = &Node{
			left:   nil,
			right:  nil,
			height: 0,
			key:    k,
		}
		t.size = 1
	} else {
		if t.root.find(k) != nil {
			return false
		}
		t.root = t.root.insert(k)
		t.size++
	}
	return true
}

func (t *AvlTree) Remove(k Key) bool {
	node := t.root.find(k)
	if node == nil {
		return false
	} else {
		t.root = t.root.remove(k)
		t.size--
		return true
	}
}

func (t *AvlTree) Min() Key {
	if t.root == nil {
		return nil
	}
	return t.root.findMin().key
}

func (t *AvlTree) Max() Key {
	if t.root == nil {
		return nil
	}
	return t.root.findMax().key
}

func (t *AvlTree) Contains(k Key) bool {
	if t.root == nil {
		return false
	}

	node := t.root.find(k)
	return node != nil
}

func newNode(key Key) *Node {
	return &Node{
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

func (t *Node) adjust() *Node {
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

func (t *Node) insert(k Key) *Node {
	if t == nil {
		return newNode(k)
	}
	res := k.CompareTo(t.key)
	if res < 0 {
		t.left = t.left.insert(k)
	} else if res > 0 {
		t.right = t.right.insert(k)
	} else {
		return t // existing key not accepted
	}
	return t.adjust()
}

func (t *Node) findMin() *Node {
	if t == nil {
		return nil
	}
	if t.left != nil {
		return t.left.findMin()
	}
	return t
}

func (t *Node) findMax() *Node {
	if t == nil {
		return nil
	}
	if t.right != nil {
		return t.right.findMax()
	}
	return t
}

func (t *Node) remove(k Key) *Node {
	if t == nil {
		return nil
	}

	newRoot := t
	res := k.CompareTo(t.key)
	if res < 0 {
		t.left = t.left.remove(k)
	} else if res > 0 {
		t.right = t.right.remove(k)
	} else {
		if t.left == nil {
			newRoot = t.right
			t.right = nil
		} else if t.right == nil {
			newRoot = t.left
			t.left = nil
		} else {
			pred := t.left.findMax()
			t.left = t.left.remove(pred.key)
			t.key = pred.key
		}
	}

	return newRoot.adjust()
}

func (t *AvlTree) apply(f func(Key)) {
	if t.root == nil {
		return
	}
	t.root.traverse(f)
}

// ToSlice returns a slice of stored values, which are sorted
func (t *AvlTree) ToSlice() []Key {
	res := make([]Key, 0, t.size)
	appendToSlice := func(k Key) {
		res = append(res, k)
	}
	t.apply(appendToSlice)
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

func (t *Node) find(k Key) *Node {
	if t == nil {
		return nil
	}
	res := k.CompareTo(t.key)
	if res < 0 {
		return t.left.find(k)
	} else if res > 0 {
		return t.right.find(k)
	} else {
		return t
	}
}
