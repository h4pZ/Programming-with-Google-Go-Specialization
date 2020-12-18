package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func (animal Animal) PrintAttribute(field string) {
	switch field {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("No information available for:", field)
	}
}

func main() {
	// Initializing the animals.
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peek"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	var animal, attribute string
	animals := make(map[string]Animal)
	animals["cow"] = cow
	animals["bird"] = bird
	animals["snake"] = snake

	// Infinite prompt.
	fmt.Println("Please input the query on the prompt as: animal attribute. (eg. cow eat)")

	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&animal, &attribute)

		if err != nil {
			fmt.Println("[ERROR]", err)
			continue
		}

		animal, attribute = strings.ToLower(animal), strings.ToLower(attribute)

		if val, ok := animals[animal]; ok {
			val.PrintAttribute(attribute)
		} else {
			fmt.Printf("[ERROR] Animal %s not found\n", animal)
		}

	}

}
