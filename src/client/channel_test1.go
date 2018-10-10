package main

import (
	"fmt"
	"strconv"
	"time"
)

var map_test = map[string]int{}
var quit chan int = make(chan int, 10000)
var quit2 chan int = make(chan int, 2)

func write(i int) {
	quit <- i
}
func read() {
	for {
		if len(quit) > 0 {
			i := <-quit
			map_test["aa"+strconv.Itoa(i)] = i
		}
		if len(map_test) == 500000 {
			fmt.Println("Done")
			break
		}
	}
	quit2 <- 1
}

func main() {

	start := time.Now().UnixNano()

	go read()
	for i := 0; i < 500000; i++ {
		go write(i)
	}

	end := time.Now().UnixNano()
	<-quit2
	fmt.Println(len(map_test))
	fmt.Println("time =", (end - start))
}
