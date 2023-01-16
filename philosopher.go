package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

var host = make(chan bool, 2)
var wg sync.WaitGroup

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		host <- true // 1
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("%d Philosopher is eating\n", p.id+1)
		time.Sleep(2 * time.Second)
		fmt.Printf("%d Philospher finishes eating\n", p.id+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host
	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat()
	}

	wg.Wait()
	var input string
	fmt.Scanln(&input)
}
