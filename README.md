# mapi
mapi is a map interface. It provide a uniq interface for various type of map engines. currently it supports two map engines: GoMap, GoSycMap, WeakMap and BigMap(bigtype.Map). However its extensible and new map engines can be appended easily. mapfactory just provides simple and basic maps (not parallel or thread-safe concurrent maps like sync.map). Because concurrency is different thing and any simple basic map can be used with concurrent solutions.

# usage
  any application that is using go standard map[K]V; can replace mapfactory to switch between map enginges easily.

# example
```go
m1 := mapf.NewMap[int, string](mapf.GoMap),
m2 := mapf.NewMap[string, float](mapf.BigMap) //, hintSize, kMarshaler, vMarshaler, checkCollision),

m1.Set(10, "ten")
m2.Set("Pi", 3.14)

v, ok := m1.Get(10)  // ten, true
_, ok = m2.Get("pi") // 0, false
```



iwinidea:~/github/mapf$ cat README.md
# mapfactory
mapf is a mapfactory. It provide a uniq interface for various type of map engines. currently it supports two map engines: GoMap, GoSycMap, WeakMap and BigMap(bigtype.Map). However its extensible and new map engines can be appended easily. mapfactory just provides simple and basic maps (not parallel or thread-safe concurrent maps like sync.map). Because concurrency is different thing and any simple basic map can be used with concurrent solutions.

# usage
  any application that is using go standard map[K]V; can replace mapfactory to switch between map enginges easily.

# example
```go
m1 := mapf.NewMap[int, string](mapf.GoMap),
m2 := mapf.NewMap[string, float](mapf.BigMap) //, hintSize, kMarshaler, vMarshaler, checkCollision),

m1.Set(10, "ten")
m2.Set("Pi", 3.14)

v, ok := m1.Get(10)  // ten, true
_, ok = m2.Get("pi") // 0, false
```
