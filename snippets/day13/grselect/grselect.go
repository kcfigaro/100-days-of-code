package main

import (
	"fmt"
	"sync"
	"time"
)

func consumer(wg *sync.WaitGroup, c chan int, q chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("consumer started")
	// for i := range c {
	// 	fmt.Println("i =", i)
	// 	time.Sleep(1 * time.Second)
	// }
	for {
		select {
		case i := <-c:
			fmt.Println("i =", i)
			time.Sleep(1 * time.Second)
		case <-q:
			fmt.Println("consumer finished. press ctrl+c to exit")
			wg.Done()
			return
		}
	}
}

func main() {

	var wg sync.WaitGroup
	c := make(chan int, 5)
	q := make(chan bool)

	wg.Add(1)
	//consumer
	go consumer(&wg, c, q)

	//prodcuer
	for i := 1; i <= 5; i++ {
		c <- i
		// time.Sleep(1 * time.Second)
	}
	fmt.Println("producer finished", len(c))

	q <- true
	wg.Wait()

	fmt.Println("Waiting for goroutines to finish...")
	fmt.Println("Done!")

}
