package lru_test

import (
	"lru"
	"testing"
)

func TestInsert(t *testing.T) {

	l := lru.New(12)
	l.Insert(1)
	l.Get(1)
}
