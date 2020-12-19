package main

import (
	"fmt"
	"strings"
)

func main() {
	var command string
	var animal_name string
	var function string
	animals := make([]Animal, 0)

	// Infinite prompt.
	fmt.Println("Please input the query on the prompt as: command name/type function. (eg. query cow eat)")

	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&command, &animal_name, &function)

		if err != nil {
			fmt.Println("[ERROR]", err)
			continue
		}

		// Making sure every input is lowercase
		command = strings.ToLower(command)
		animal_name = strings.ToLower(animal_name)
		function = strings.ToLower(function)

		switch command {
		case "query":
			QueryAnimal(animals, animal_name, function)

		case "newanimal":
			new_animal := CreateNewAnimal(animal_name, function)

			if new_animal == nil {
				fmt.Printf("[ERROR] No animal type %s found \n", animal_name)
			} else {
				animals = append(animals, new_animal)
				fmt.Println("Created it!")
			}

		default:
			fmt.Printf("Command %s not available", command)

		}

	}

}

type Animal interface {
	Eat()
	Move()
	Speak()
	Name() string
}

type Cow struct{ name string }
type Bird struct{ name string }
type Snake struct{ name string }

// Cow methods for the animal interface.
func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (cow Cow) Name() string {
	return cow.name
}

// Bird methods for the animal interface.
func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

func (bird Bird) Name() string {
	return bird.name
}

// Snake methods for the animal interface.
func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func (snake Snake) Name() string {
	return snake.name
}

func CreateNewAnimal(name, animal_type string) Animal {
	var animal Animal

	switch animal_type {
	case "cow":
		animal = Cow{name: name}
	case "bird":
		animal = Bird{name: name}
	case "snake":
		animal = Snake{name: name}
	}

	return animal
}

func QueryAnimal(animals []Animal, name, function string) {
	for _, animal := range animals {
		if animal.Name() == name {
			switch function {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Printf("The animal %s function %s has not been found\n", name, function)
			}

			return
		}
	}

	fmt.Printf("The animal %s couldn't be found!\n", name)
	return
}
