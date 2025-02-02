package mapi

import (
	"sync"
	"sync/atomic"
)

type syncMap[K comparable, V any] struct {
	m sync.Map

	//syncMap doesnt support Len() natively
	//then this counter was not always accurate
	//but almost for all application it can be used
	//to have a Len of Maps approximately(not accurately)
	atomicCounter int64
}

func newSyncMap[K comparable, V any]() *syncMap[K, V] {
	return &syncMap[K, V]{}
}

func (m *syncMap[K, V]) Set(key K, value V) {
	m.m.Store(key, value)
	atomic.AddInt64(&m.atomicCounter, 1)
}

func (m *syncMap[K, V]) Get(key K) (value V, ok bool) {
	val, ok := m.m.Load(key)
	if ok {
		value, ok = val.(V)
	}
	return value, ok
}

func (m *syncMap[K, V]) Len() int {
	return int(atomic.LoadInt64(&m.atomicCounter))
}

func (m *syncMap[K, V]) Seq(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *syncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
	atomic.AddInt64(&m.atomicCounter, -1)
}
