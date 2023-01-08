package main

import (
	"fmt"
	"sync"
	"time"
)

// func squares(c, quit chan int) {
// 	for {
// 		select {
// 		case num := <-c:
// 			fmt.Println(num * num)
// 			// case <-quit:
// 			// 	fmt.Println("quit")
// 			// 	return
// 		}
// 	}
// }

func writeChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	defer wg.Done()
	for i := 1; i <= stop; i++ {
		fmt.Printf("Writing %v\n", i)
		limitchannel <- i
	}
	// time.Sleep(1 * time.Second)

}

func readChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	// defer wg.Done()
	// for i := 1; i <= stop; i++ {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Print("Reading ")
	// 	fmt.Println(<-limitchannel)
	// }
	// for {
	// 	select {
	// 	case x := <-limitchannel:
	// 		time.Sleep(time.Second)
	// 		fmt.Print("Reading ")
	// 		fmt.Println(x)
	// 		fmt.Println(len(limitchannel))
	// 	case <-quit:
	// 		fmt.Println("quit")
	// 		fmt.Println(len(quit))
	// 		wg.Done()
	// 		return
	// 	default:
	// 	}
	// }
	for x := range limitchannel {
		fmt.Print("Reading ")
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	limitchannel := make(chan int, 3)
	// quit := make(chan int)
	// defer close(limitchannel)
	// defer close(quit)
	go writeChannel(wg, limitchannel, 3)
	go readChannel(wg, limitchannel, 3)
	time.Sleep(1 * time.Second)
	close(limitchannel)
	wg.Wait()
	time.Sleep(1 * time.Second)
}

// https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/#:~:text=Communicating%20between%20Goroutines&text=Go%20provides%20a%20way%20for,or%20receive%20data%20between%20goroutines.
// note that the number of sending operations must be equal to the number of receiving operations because if you send data to a channel and don't receive it elsewhere, you get a fatal error: all goroutines are asleep - deadlock!.
