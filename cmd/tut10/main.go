package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var host = make(chan bool, 2)

type ChopS struct{ sync.Mutex }

type Philo struct {
	id      int
	leftCS  *ChopS
	rightCS *ChopS
}

func (p Philo) eat() {

	for i := 0; i < 3; i++ {

		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("Starting to eat: ", p.id+1)
		time.Sleep(2 * time.Second)
		fmt.Println("finishing eating: ", p.id+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		<-host

	}
	wg.Done()
}

func main() {

	ChopSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopS)
	}

	Philosophers := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philo{i, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Philosophers[i].eat()
	}

	wg.Wait()

}
