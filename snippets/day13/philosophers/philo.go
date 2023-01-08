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
	num, count      int
	leftCs, rightCs *ChopS
}

func host(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for i := 0; i < 15; i++ {
		c <- i % 5
		time.Sleep(1 * time.Second)
	}
}

func (p Philo) eat(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		if p.count > 0 {
			p.leftCs.Lock()
			p.rightCs.Lock()
			fmt.Printf(" starting to eat %v\n", p.num+1)
			p.count -= 1
			time.Sleep(1 * time.Second)
			fmt.Printf("finishing to eat %v\n", p.num+1)
			p.rightCs.Unlock()
			p.leftCs.Unlock()
		}
	}
}

func main() {
	CStick := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CStick[i] = new(ChopS)
	}

	philo := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philo[i] = &Philo{i, 3, CStick[i], CStick[(i+1)%5]}
	}

	wg := new(sync.WaitGroup)
	wg.Add(5)

	// no more than 2 philosophers to eat concurrently
	ch := make(chan int, 2)
	defer close(ch)

	// execute host to put philosophers into the channel
	go host(wg, ch)

	// eat
	for i := 0; i < 5; i++ {
		go philo[i].eat(wg, ch)
	}
	wg.Wait()
}
