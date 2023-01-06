package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mySort(s []int, c chan []int) {
	fmt.Printf("subarray %v is sorting ", s)
	sort.Ints(s)
	fmt.Println(s)
	c <- s
}

func main() {
	fmt.Println("Enter integers, please! e.g 20 19 43 36 17 1 9 2 49 11 28 6")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fields := strings.Fields(input)

	fmt.Println("---goroutine starts---")
	sli := []int{}
	for i := 0; i < len(fields); i++ {
		val, e := strconv.Atoi(fields[i])
		if e == nil {
			sli = append(sli, val)
		}
	}

	// channels to transfer data between goroutines
	c := make(chan []int)

	// entire sorted slice
	large := make([]int, 0)

	// partition the slice into 4 parts
	part := len(fields) / 4
	end := 0
	for i := 0; i < len(fields); i += part {
		end += part
		if end > len(fields) {
			end = len(fields)
		}
		// sorted by a different goroutine
		go mySort(sli[i:end], c)
		x := <-c
		large = append(large, x...)
	}

	fmt.Println("---goroutine ends---")
	sort.Ints(large)
	fmt.Printf("entire sorted list: %v\n", large)
}
