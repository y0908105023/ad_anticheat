package main

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

var sum AtomicFloat64 = 0.0

type AtomicFloat64 uint64

func (f *AtomicFloat64) Value() float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(f)))
}

func (f *AtomicFloat64) Add(n float64) {
	for {
		a := atomic.LoadUint64((*uint64)(f))
		b := math.Float64bits(math.Float64frombits(a) + n)
		if atomic.CompareAndSwapUint64((*uint64)(f), a, b) {
			return
		}
	}
}

func call(b float64) {
	sum.Add(b)
}

func main() {
	start := time.Now().UnixNano() / 1000000

	for i := 0; i < 100; i++ {
		go call(float64(i))
	}
	time.Sleep(time.Second * 1)
	fmt.Println("sum = %f", sum.Value())

	end := time.Now().UnixNano() / 1000000
	fmt.Println("Program exit. time->", (end - start))
}
