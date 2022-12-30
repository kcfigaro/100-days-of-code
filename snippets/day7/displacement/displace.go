package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, o_v, o_d float64) func(float64) float64 {
	fn := func(t float64) float64 {
		// https://www.wikihow.com/Calculate-Displacement
		return (a*math.Pow(t, 2))/2 + o_v*t + o_d
	}
	return fn
}

func main() {
	var a, o_v, o_d, t float64
	fmt.Print("Enter acceleration, initial velocity, and initial displacement: ")
	fmt.Scanln(&a, &o_v, &o_d)

	fn := GenDisplaceFn(a, o_v, o_d)

	fmt.Print("Enter a value for time: ")
	fmt.Scan(&t)
	fmt.Printf("---Result---\nafter %v seconds...\nthe final displacement: ", t)
	fmt.Printf("%v\n", fn(t))
}
