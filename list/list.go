package list

type Element struct {
	prev *Element
	next *Element

	belongsTo *List

	Value interface{}
}

func (e *Element) Next() *Element {
	if e.belongsTo == nil || e.next.belongsTo == nil {
		return nil
	}
	return e.next
}

func (e *Element) Prev() *Element {
	if e.belongsTo == nil || e.prev.belongsTo == nil {
		return nil
	}
	return e.prev
}

type IElement interface {
	Next() *Element
	Prev() *Element
}

type List struct {
	root Element
	len  int
}

func New() *List {
	l := List{}
	l.Init()
	return &l
}

type IList interface {
	Back() *Element
	Front() *Element
	Init() *List
	InsertAfter(v interface{}, mark *Element) *Element
	InsertBefore(v interface{}, mark *Element) *Element
	Len() int
	MoveAfter(e, mark *Element)
	MoveBefore(e, mark *Element)
	MoveToBack(e *Element)
	MoveToFront(e *Element)
	PushBack(v interface{}) *Element
	PushBackList(other *List)
	PushFront(v interface{}) *Element
	PushFrontList(other *List)
	Remove(e *Element) interface{}
}

func (l *List) Back() *Element {
	back := l.root.prev
	if back == &l.root {
		return nil
	}
	return back
}

func (l *List) Front() *Element {
	front := l.root.next
	if front == &l.root {
		return nil
	}
	return front
}

func (l *List) Init() *List {
	l.root.prev = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

func (l *List) checkAndInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.belongsTo != l {
		return nil
	}
	e := &Element{
		prev:      mark,
		next:      mark.next,
		belongsTo: l,
		Value:     v,
	}
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.belongsTo != l {
		return nil
	}
	e := &Element{
		prev:      mark.prev,
		next:      mark,
		belongsTo: l,
		Value:     v,
	}
	e.prev.next = e
	e.next.prev = e
	e.belongsTo = l
	l.len++
	return e
}

func (l *List) Len() int {
	return l.len
}

func (l *List) MoveAfter(e, mark *Element) {
	if e.belongsTo != l || mark.belongsTo != l || e == mark {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = mark
	e.next = mark.next
	e.prev.next = e
	e.next.prev = e
}

func (l *List) MoveBefore(e, mark *Element) {
	if e.belongsTo != l || mark.belongsTo != l || e == mark {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = mark
	e.prev = mark.prev
	e.prev.next = e
	e.next.prev = e
}

func (l *List) MoveToBack(e *Element) {
	l.checkAndInit()
	if e.belongsTo != l || e == l.root.prev {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = l.root.prev
	e.next = &l.root
	e.prev.next = e
	e.next.prev = e
}

func (l *List) MoveToFront(e *Element) {
	l.checkAndInit()
	if e.belongsTo != l || e == l.root.next {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = l.root.next
	e.prev = &l.root
	e.prev.next = e
	e.next.prev = e
}

func (l *List) PushBack(v interface{}) *Element {
	l.checkAndInit()
	e := &Element{
		prev:      l.root.prev,
		next:      &l.root,
		belongsTo: l,
		Value:     v,
	}
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

func (l *List) PushBackList(other *List) {
	l.checkAndInit()
	if other.len == 0 {
		return
	}
	tail := l.root.prev

	stop := tail

	for e := other.Front(); e != stop && e != nil; e = e.Next() {
		e2 := &Element{
			prev:      tail,
			next:      tail.next,
			belongsTo: l,
			Value:     e.Value,
		}
		tail.next = e2
		tail = tail.next
		if e == stop {
			break
		}
	}
	l.root.prev = tail
	l.len += other.len
}

func (l *List) PushFront(v interface{}) *Element {
	l.checkAndInit()
	e := &Element{
		prev:      &l.root,
		next:      l.root.next,
		belongsTo: l,
		Value:     v,
	}
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

func (l *List) PushFrontList(other *List) {
	l.checkAndInit()
	if other.len == 0 {
		return
	}
	head := l.root.next
	stop := head

	for e := other.Back(); e != nil; e = e.Prev() {
		e2 := &Element{
			prev:      head.prev,
			next:      head,
			belongsTo: l,
			Value:     e.Value,
		}
		head.prev = e2
		head = head.prev
		if e == stop {
			break
		}
	}
	l.root.next = head
	l.len += other.len
}

func (l *List) Remove(e *Element) interface{} {
	l.checkAndInit()
	if e.belongsTo != l {
		return nil
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	e.belongsTo = nil
	value := e.Value
	//e.Value = nil // should not set to nil, see https://groups.google.com/forum/#!topic/golang-nuts/HGCY7IanlvU
	l.len--
	return value
}
