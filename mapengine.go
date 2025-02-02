package mapi

type MapEngine int

const (
	NotDefinedMap MapEngine = iota
	GoMap
	GoSyncMap
	BigMap
	GoWeakMap
)

func (e MapEngine) String() string {
	switch e {
	case NotDefinedMap:
		return "NotDefinedMap"
	case GoMap:
		return "StdMap"
	case GoSyncMap:
		return "SyncMap"
	case BigMap:
		return "BigMap"
	case GoWeakMap:
		return "WeakMap"
	}
	return "Unknown"
}
