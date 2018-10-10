package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type BeeMap struct {
	lock *sync.RWMutex
	bm   map[string]int
}

func NewBeeMap() *BeeMap {
	return &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[string]int),
	}
}

//Get from maps return the k's value
func (m *BeeMap) Get(k string) int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.bm[k]; ok {
		return val

	}
	return 0
}

// Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *BeeMap) Set(k string, v int) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.bm[k]; !ok {
		m.bm[k] = v
	} else if val != v {
		m.bm[k] = v
	} else {
		return false
	}
	return true
}

// Returns true if k is exist in the map.
func (m *BeeMap) Check(k string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.bm[k]; !ok {
		return false
	}
	return true
}

func (m *BeeMap) Delete(k string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, k)
}

var quit chan int = make(chan int, 10)

func call(i int, beemap *BeeMap) {
	beemap.Set("aa"+strconv.Itoa(i), i)
	quit <- i
	fmt.Println("len(quit) = ", len(quit))
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 10
}

func main() {
	// false is illegal
	beemap := new(BeeMap)
	beemap.bm = map[string]int{}
	beemap.lock = new(sync.RWMutex)

	aeemap := map[string]int{}

	t1 := currentTimeMillis()

	for i := 0; i < 2; i++ {
		aeemap["aa"+strconv.Itoa(i)] = i
	}
	t2 := currentTimeMillis()
	for i := 0; i < 2; i++ {
		a := aeemap["aa"+strconv.Itoa(i)]
		if a < 0 {
			fmt.Println(a)
		}
	}
	t3 := currentTimeMillis()
	for i := 0; i < 2; i++ {
		go call(i, beemap)
	}

	for i := 0; i < 2; i++ {
		<-quit
	}
	t4 := currentTimeMillis()
	for i := 0; i < 1; i++ {
		a := beemap.Get("aa" + strconv.Itoa(i))
		if a < 0 {
			fmt.Println(a)
		}
	}
	t5 := currentTimeMillis()

	// fmt.Println(t1, t2, t3)
	fmt.Println(len(aeemap), len(beemap.bm))
	fmt.Println("Program exit. time->", (t2 - t1))
	fmt.Println("Program exit. time->", (t3 - t2))
	fmt.Println("Program exit. time->", (t4 - t3))
	fmt.Println("Program exit. time->", (t5 - t4))

}
