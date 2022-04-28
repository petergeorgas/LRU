package lru

import (
	"container/list"
	"fmt"
)

type LRUCache interface {
	Insert(any)
	Get(any)
}

type lruCache struct {
	Capacity int
	Mp       map[any]*list.Element
	List     list.List
}

func New(capacity int) *lruCache {
	return &lruCache{
		Capacity: capacity,
		Mp:       make(map[any]*list.Element),
		List:     *list.New(),
	}
}

func (l *lruCache) Insert(v any) {
	newE := l.List.PushFront(v)
	l.Mp[v] = newE
}

func (l *lruCache) Get(v any) {
	if val, ok := l.Mp[v]; ok {
		fmt.Println(val.Value)
		l.List.MoveToFront(val)
		fmt.Printf("l.List: %v\n", l.List)
	} else {
		fmt.Println("VALUE NOT FOUND!")
	}
}
