package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(ch chan int) {
	for {
		ch <- rand.Intn(101)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int, 3)
	go generator(ch)

	for i := 0; i < 15; i++ {
		fmt.Println(<-ch)
	}

	fmt.Printf("len : %d\n", len(ch))
	fmt.Printf("cap : %d", cap(ch))
}
