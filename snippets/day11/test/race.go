package main

import (
	"fmt"
	"time"
)

var x int = 1

func say() {
	fmt.Printf("say : %v\n", x)
}

func plus() {
	x = x + 1
	fmt.Printf("plus : %v\n", x)
}

func main() {
	go plus()
	go say()
	time.Sleep(1 * time.Second)
}

// Explaination: Race condition
// Outcome depends on non-deterministic ordering
// means that sometimes say function comes out before plus function
// for instance,
// » ./race
// plus : 2
// say : 2
// » ./race
// say : 1
// plus : 2
