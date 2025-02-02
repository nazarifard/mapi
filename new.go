package mapi

import (
	"github.com/nazarifard/bigtype"
	"github.com/nazarifard/gomap"
)

func NewMap[K comparable, V any](engine MapEngine, options ...any) Map[K, V] {
	switch engine {
	case GoMap:
		if len(options) > 0 {
			hintSize, ok := options[0].(int)
			if ok {
				return gomap.New[K, V](hintSize)
			}
		}
		return gomap.New[K, V]()

	case GoSyncMap:
		return newSyncMap[K, V]()
	//case GoWeakMap:
	//	return NewWeakMap[K, P]()
	case BigMap:
		return bigtype.NewMap[K, V](options...)
	}
	panic("invalid map engine")
}
