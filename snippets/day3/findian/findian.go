package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	text := strings.ToLower(scanner.Text())
	if strings.HasPrefix(text, "i") && strings.Contains(text, "a") && strings.HasSuffix(text, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
