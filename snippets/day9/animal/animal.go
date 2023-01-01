package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
}

type cow struct {
	food       string
	locomotion string
	noise      string
}

type bird struct {
	food       string
	locomotion string
	noise      string
}

type snake struct {
	food       string
	locomotion string
	noise      string
}

func Eat(a Animal) {
	switch sh := a.(type) {
	case cow:
		fmt.Println(sh.food)
	case bird:
		fmt.Println(sh.food)
	case snake:
		fmt.Println(sh.food)
	}
}

func Move(a Animal) {
	switch sh := a.(type) {
	case cow:
		fmt.Println(sh.locomotion)
	case bird:
		fmt.Println(sh.locomotion)
	case snake:
		fmt.Println(sh.locomotion)
	}
}

func Speak(a Animal) {
	switch sh := a.(type) {
	case cow:
		fmt.Println(sh.noise)
	case bird:
		fmt.Println(sh.noise)
	case snake:
		fmt.Println(sh.noise)
	}
}

var aMap map[string]Animal

func main() {
	var a Animal
	aMap = make(map[string]Animal)

	fmt.Println("How To Use:\n 1. create a new animal: The first string is \"newanimal\". The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either \"cow\", \"bird\", or \"snake\". e.g newanimal jake snake\n 2. quary an animal: The first string is \"query\". The second string is the name of the animal. The third string is the name of the information requested about the animal, either \"eat\", \"move\", or \"speak\". e.g query jake eat\n ")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		fields := strings.Fields(input)

		if len(fields) > 0 {
			a_name := fields[1]

			if fields[0] == "newanimal" {
				a_type := fields[2]
				fmt.Printf("Created it! %s information for %s\n", a_type, a_name)
				switch a_type {
				case "cow":
					aMap[a_name] = cow{food: "grass", locomotion: "walk", noise: "moo"}
				case "bird":
					aMap[a_name] = bird{food: "worms", locomotion: "fly", noise: "peep"}
				case "snake":
					aMap[a_name] = snake{food: "mice", locomotion: "slither", noise: "hsss"}
				default:
					fmt.Println("Wrong animal type! Please try with either cow, bird, or snake")
				}
			} else if fields[0] == "query" {
				a_action := fields[2]

				a = Animal(aMap[a_name])
				switch a_action {
				case "eat":
					Eat(a)
				case "move":
					Move(a)
				case "speak":
					Speak(a)
				}
			}
		} else {
			fmt.Println("Please enter any input!")
		}
	}
}
