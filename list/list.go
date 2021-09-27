package list

type Element struct {
	prev *Element
	next *Element

	belongsTo *List

	Value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

type IElement interface {
	Next() *Element
	Prev() *Element
}

type List struct {
	root *Element
	len  int
}

func New() *List {
	root := new(Element)
	root.prev = root
	root.next = root

	return &List{
		root: root,
		len:  0,
	}
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
	if back == l.root {
		return nil
	}
	return back
}

func (l *List) Front() *Element {
	front := l.root.next
	if front == l.root {
		return nil
	}
	return front
}

func (l *List) Init() *List {
	l.root.prev = l.root
	l.root.next = l.root
	l.len = 0
	return l
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
	panic("implement me")
}

func (l *List) MoveBefore(e, mark *Element) {
	panic("implement me")
}

func (l *List) MoveToBack(e *Element) {
	panic("implement me")
}

func (l *List) MoveToFront(e *Element) {
	panic("implement me")
}

func (l List) PushBack(v interface{}) *Element {
	panic("implement me")
}

func (l *List) PushBackList(other *List) {
	panic("implement me")
}

func (l *List) PushFront(v interface{}) *Element {
	panic("implement me")
}

func (l *List) PushFrontList(other *List) {
	panic("implement me")
}

func (l *List) Remove(e *Element) interface{} {
	panic("implement me")
}
