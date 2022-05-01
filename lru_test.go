package lru_test

import (
	"fmt"
	"lru"
	"testing"
)

func TestInsert(t *testing.T) {

	l := lru.New(2)
	l.Put(1)
	l.Put(2)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
	l.Put(3)              // Will cause 1 to be vactated
	fmt.Println(l.Get(1)) // Return -1 (Not Found -- 1 should be vacated)
	fmt.Println(l.Get(2)) // Should be found
	fmt.Println(l.Get(3)) // Should be found
	l.Put(4)              // Will cause 2 to be vacated
	fmt.Println(l.Get(2)) // Return -1 Should be vacated
}
