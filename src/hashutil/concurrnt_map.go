package hashutil

import (
	"sync"
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

func (m *BeeMap) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	for key, _ := range m.bm {
		delete(m.bm, key)
	}
}

func (m *BeeMap) Size() int {
	return len(m.bm)
}
