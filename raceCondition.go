package main

import (
	"fmt"
	"sync"
)

var counter int

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go increment("goroutine 1: ")
	go increment("goroutine 2: ")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func increment(s string) {
	for i := 0; i < 100; i++ {
		counter++
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
