package lru

import "container/list"

type LRUCache interface {
	Insert(any)
	Get(any)
}

type lruCache struct {
	Capacity int
	Mp       map[any]list.Element
	List     list.List
}

func New(capacity int) *lruCache {
	return &lruCache{Capacity: capacity}
}

func (l *lruCache) Insert(any) {

}
