package mapi

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	bigopts "github.com/nazarifard/bigtype/options"
)

func WUpsert(m Map[string, *Person]) {
	for i := range MAX_SIZE / UNIT {
		start := time.Now()
		for range UNIT {
			aPerson.Name = inputString[rand.Int31n(256)]
			m.Set(fmt.Sprintf("%012d", rand.Int31n(MAX_SIZE)), &aPerson)
		}
		_, _ = i, start
		//fmt.Printf("i:%d time:%v\n", i, time.Since(start))
	}
}

func WSearch(m Map[string, *Person]) {
	for i := range MAX_SIZE / UNIT {
		start := time.Now()
		for range UNIT {
			person, ok := m.Get(fmt.Sprintf("%012d", rand.Int31n(MAX_SIZE)))
			if ok {
				_, _ = *person, ok
			}
		}
		_, _ = i, start
		//fmt.Printf("i:%d time:%v\n", i, time.Since(start))
	}
}

func WCheckup(m Map[string, *Person]) bool {
	for i := range 1000 {
		aPerson.Name = fmt.Sprintf("%012d", i)
		m.Set(aPerson.Name, &aPerson)
	}
	for i := range 1000 {
		aPerson.Name = fmt.Sprintf("%012d", i)
		m.Set(aPerson.Name+aPerson.Name, &aPerson)
	}
	for i := range 1000 {
		key := fmt.Sprintf("%012d", i)
		person, ok := m.Get(key + key)
		if !ok || person.Name+person.Name != key+key {
			return false
		}
	}
	return true
}

func TestWeakMap(t *testing.T) {
	var bigOpts bigopts.MapOptions[string, Person]
	_ = bigOpts
	maps := [...]Map[string, *Person]{
		NewWeakMap[string, *Person](),
	}
	fmt.Printf("MAX_SIZE:%v, UNIT: %v\n", MAX_SIZE, UNIT)

	var insert, update, search []time.Duration
	//fmt.Println("Insert: ----------")
	for i := range maps {
		//fmt.Printf("Map: %v ...\n", MapEngine(i+1).String())
		now := time.Now()
		WUpsert(maps[i])
		insert = append(insert, time.Since(now))
	}

	//fmt.Println("Update: ----------")
	for i := range maps {
		//fmt.Printf("Map: %v ...\n", MapEngine(i+1).String())
		now := time.Now()
		WUpsert(maps[i])
		update = append(update, time.Since(now))
	}

	//fmt.Println("Search: ----------")
	for i := range maps {
		//fmt.Printf("Map: %v ...\n", MapEngine(i+1).String())
		now := time.Now()
		WSearch(maps[i])
		search = append(search, time.Since(now))
	}

	fmt.Println("\nEngine\tinsert\t\tupdate\t\tsearch")
	for i := range len(maps) {
		fmt.Printf("%v\t%v\t%v\t%v\n", MapEngine(i+1).String(), insert[i], update[i], search[i])
	}
}
