package mapi

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	bigopts "github.com/nazarifard/bigtype/options"
)

const UNIT = 10_000
const MAX_SIZE = 2 * UNIT

// Person = goserbench.SmallStruct
type Person struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

var aPerson = Person{
	Name:     "",
	Phone:    "0987654321",
	BirthDay: time.Now(),
	Siblings: 7,
	Spouse:   true,
	Money:    3000.1415,
}

var inputString = func() [256]string {
	alphabet := "0123456789qwertyuioplkjhgfdsazxcvbnmMNBVCXZASDFGHJKLPOIUYTREWQ"
	var out [256]string
	out[0] = "0"
	for i := range 255 {
		out[i+1] = out[i] + string(alphabet[i%len(alphabet)])
	}
	return out
}()

func Upsert(m Map[string, Person]) {
	for i := range MAX_SIZE / UNIT {
		start := time.Now()
		for range UNIT {
			aPerson.Name = inputString[rand.Int31n(256)]
			m.Set(fmt.Sprintf("%012d", rand.Int31n(MAX_SIZE)), aPerson)
		}
		_, _ = i, start
		//fmt.Printf("i:%d time:%v\n", i, time.Since(start))
	}
}

func Search(m Map[string, Person]) {
	for i := range MAX_SIZE / UNIT {
		start := time.Now()
		for range UNIT {
			person, ok := m.Get(fmt.Sprintf("%012d", rand.Int31n(MAX_SIZE)))
			_, _ = person, ok
		}
		_, _ = i, start
		//fmt.Printf("i:%d time:%v\n", i, time.Since(start))
	}
}

func Checkup(m Map[string, Person]) bool {
	for i := range 1000 {
		aPerson.Name = fmt.Sprintf("%012d", i)
		m.Set(aPerson.Name, aPerson)
	}
	for i := range 1000 {
		aPerson.Name = fmt.Sprintf("%012d", i)
		m.Set(aPerson.Name+aPerson.Name, aPerson)
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

//go:generate fastape Person
func TestMap(t *testing.T) {
	var bigOpts bigopts.MapOptions[string, Person]
	_ = bigOpts
	maps := [...]Map[string, Person]{
		NewMap[string, Person](GoMap),
		NewMap[string, Person](GoSyncMap),
		NewMap[string, Person](BigMap, bigOpts.WithVTape(PersonTape{})),
	}
	fmt.Printf("MAX_SIZE:%v, UNIT: %v\n", MAX_SIZE, UNIT)

	var insert, update, search []time.Duration
	//fmt.Println("Insert: ----------")
	for i := range maps {
		//fmt.Printf("Map: %v ...\n", MapEngine(i+1).String())
		now := time.Now()
		Upsert(maps[i])
		insert = append(insert, time.Since(now))
	}

	//fmt.Println("Update: ----------")
	for i := range maps {
		//fmt.Printf("Map: %v ...\n", MapEngine(i+1).String())
		now := time.Now()
		Upsert(maps[i])
		update = append(update, time.Since(now))
	}

	//fmt.Println("Search: ----------")
	for i := range maps {
		//fmt.Printf("Map: %v ...\n", MapEngine(i+1).String())
		now := time.Now()
		Search(maps[i])
		search = append(search, time.Since(now))
	}

	fmt.Println("\nEngine\tinsert\t\tupdate\t\tsearch")
	for i := range len(maps) {
		fmt.Printf("%v\t%v\t%v\t%v\n", MapEngine(i+1).String(), insert[i], update[i], search[i])
	}
}
