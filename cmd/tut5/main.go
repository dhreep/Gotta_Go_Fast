package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal struct to hold animal information
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat method prints the animal's food
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move method prints the animal's locomotion method
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak method prints the animal's noise
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// Predefined animals
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Infinite loop to process user requests
	for {
		fmt.Print("> ")
		
		// Read user input
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		
		// Split input into parts
		parts := strings.Fields(input)
		
		// Validate input
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter an animal and an action.")
			continue
		}

		animalName := parts[0]
		action := parts[1]

		// Check if animal exists
		animal, exists := animals[animalName]
		if !exists {
			fmt.Println("Unknown animal. Choose cow, bird, or snake.")
			continue
		}

		// Call the appropriate method based on action
		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid action. Choose eat, move, or speak.")
		}
	}
}