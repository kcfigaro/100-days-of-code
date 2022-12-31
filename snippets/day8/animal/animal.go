package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) PrintResult(action string) {
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Wrong action! Please try with either eat, move, or speak")
	}
}

func main() {
	for {
		var name, action string

		fmt.Print("Enter name of an animal and action: \n> ")
		_, err := fmt.Scanln(&name, &action)

		if err == nil {
			switch name {
			case "cow":
				a := Animal{food: "grass", locomotion: "walk", noise: "moo"}
				a.PrintResult(action)
			case "bird":
				a := Animal{food: "worms", locomotion: "fly", noise: "peep"}
				a.PrintResult(action)
			case "snake":
				a := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
				a.PrintResult(action)
			default:
				fmt.Println("Wrong animal name! Please try with either cow, bird, or snake")
			}
		} else {
			fmt.Print("Please enter both name and action properly! e.g cow eat\n")
		}
	}
}

// func main()  {
// 	animals := map[string]Animal{
// 			"cow":   Animal{food: "grass", locomotion: "walk", noise: "moo"},
// 			"bird":  Animal{food: "worms", locomotion: "fly", noise: "peep"},
// 			"snake": Animal{food: "mice", locomotion: "slither", noise: "hsss"}}

// 	var animal, request string
// 	for {
// 			fmt.Print(">")
// 			fmt.Scan(&animal, &request)
// 			a := animals[animal]
// 			switch request {
// 			case "eat":
// 					a.Eat()
// 			case "move":
// 					a.Move()
// 			case "speak":
// 					a.Speak()
// 			}
// 	}

// }
