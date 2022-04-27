package lru

type LRUCache interface {
	Insert(any)
	Get(any)
}

type lruCache struct {
	Capacity int
	mp       map[any]any
}

func New(capacity int) *lruCache {
	return &lruCache{Capacity: capacity}
}

func (l *lruCache) Insert(any) {

}
