package lru

import (
	"container/list"
)

type LRUCache interface {
	Put(any)
	Get(any)
}

type lruCache struct {
	Capacity int
	Mp       map[any]*list.Element
	List     *list.List
}

func New(capacity int) *lruCache {
	return &lruCache{
		Capacity: capacity,
		Mp:       make(map[any]*list.Element),
		List:     list.New(),
	}
}

func (l *lruCache) Put(v any) {
	val, ok := l.Mp[v]
	if ok { // If the inserted value is already present
		l.List.MoveToFront(val)
	} else {
		newE := l.List.PushFront(v)
		l.Mp[v] = newE

		if l.List.Len() > l.Capacity { // We went over capacity -- vacate LRU item
			back := l.List.Back()
			delete(l.Mp, back.Value)
			l.List.Remove(back)
		}
	}
}

func (l *lruCache) Get(v any) any {
	if val, ok := l.Mp[v]; ok {
		l.List.MoveToFront(val)
		return val.Value
	}

	return -1
}
