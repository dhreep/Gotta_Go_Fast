package main

import (
	"fmt"
	"strings"
)

// Animal interface defines the methods for animals
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type implements Animal interface
type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird type implements Animal interface
type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake type implements Animal interface
type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// Map to store animals
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")

		var command, name, animalType string
		fmt.Scanln(&command, &name, &animalType)

		// Convert input to lowercase for case-insensitive comparison
		command = strings.ToLower(command)
		animalType = strings.ToLower(animalType)

		switch command {
		case "newanimal":
			var newAnimal Animal
			switch animalType {
			case "cow":
				newAnimal = Cow{name: name}
			case "bird":
				newAnimal = Bird{name: name}
			case "snake":
				newAnimal = Snake{name: name}
			default:
				fmt.Println("Invalid animal type. Choose cow, bird, or snake.")
				continue
			}

			// Check if animal name already exists
			if _, exists := animals[name]; exists {
				fmt.Println("An animal with this name already exists.")
				continue
			}

			animals[name] = newAnimal
			fmt.Println("Created it!")

		case "query":
			var queryType string
			fmt.Scanln(&queryType)
			queryType = strings.ToLower(queryType)

			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch queryType {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid query type. Choose eat, move, or speak.")
			}

		default:
			fmt.Println("Invalid command. Use 'newanimal' or 'query'.")
		}
	}
}
