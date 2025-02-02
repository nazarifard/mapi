package mapi

import (
	"fmt"
	"runtime"
	"sync"
	"weak" // Hypothetical package for weak pointers
)

type WeakMap[K comparable, PV interface{ *V }, V any] struct {
	store map[K]*weak.Pointer[V]
	mu    sync.RWMutex
}

func NewWeakMap[K comparable, PV interface{ *V }, V any]() Map[K, PV] {
	return &WeakMap[K, PV, V]{
		store: make(map[K]*weak.Pointer[V]),
	}
}

func (wm *WeakMap[K, PV, V]) Len() int {
	wm.mu.RLock()
	defer wm.mu.RUnlock()
	return len(wm.store)
}

func (wm *WeakMap[K, PV, V]) Get(key K) (PV, bool) {
	wm.mu.RLock()
	defer wm.mu.RUnlock()
	if wp, ok := wm.store[key]; ok {
		return wp.Value(), true
	}
	return nil, false
}

func (wm *WeakMap[K, PV, V]) Set(key K, value PV) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	wp := weak.Make(value)
	runtime.SetFinalizer(value, func(v *V) {
		delete(wm.store, key)
	})
	wm.store[key] = &wp
}

func (wm *WeakMap[K, PV, V]) Delete(key K) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	delete(wm.store, key)
}

func (wm *WeakMap[K, PV, V]) Seq(f func(key K, value PV) bool) {
	wm.mu.RLock()
	defer wm.mu.RUnlock()
	for k, wp := range wm.store {
		v := wp.Value()
		if v != nil {
			if !f(k, v) {
				break
			}
		}
	}
}

func main() {
	// Example usage
	wm := NewWeakMap[int, *string]()
	one := "one"
	two := "two"
	wm.Set(1, &one)
	wm.Set(2, &two)

	fmt.Println("Length:", wm.Len())

	for k, v := range wm.Seq {
		fmt.Printf("Key: %d, Value: %s\n", k, *v)
	}
	one = ""
	runtime.GC()

	for k, v := range wm.Seq {
		fmt.Printf("Key: %d, Value: %s\n", k, *v)
	}
}
