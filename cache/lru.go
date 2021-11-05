package lru

import "container/list"

type LRUCache struct {
	cap  uint
	dict map[int]*list.Element
	list *list.List
}

type entry struct {
	key   int
	value interface{}
}

func New(cap uint) LRUCache {
	return LRUCache{
		cap:  cap,
		dict: make(map[int]*list.Element, 0),
		list: list.New(),
	}
}

func (c *LRUCache) Clear() {
	*c = New(c.cap)
}

func (c *LRUCache) Get(key int) (result interface{}, hit bool) {
	elem, found := c.dict[key]
	hit = found
	result = nil
	if found {
		result = (elem.Value).(entry).value
		c.list.MoveToFront(elem)
	}

	return result, hit
}

func (c *LRUCache) Put(key int, value interface{}) {
	elem, found := c.dict[key]
	if found {
		elem.Value = entry{key, value}
		c.list.MoveToFront(elem)
	} else {
		elem = c.list.PushFront(entry{key, value})
		c.dict[key] = elem
		if c.list.Len() > int(c.cap) {
			backElem := c.list.Back()
			delete(c.dict, backElem.Value.(entry).key)
			c.list.Remove(backElem)
		}
	}
}
