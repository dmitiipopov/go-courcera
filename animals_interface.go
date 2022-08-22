/*
Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

Animal - Food eaten - Locomtion method - Spoken sound

cow - grass - walk - moo
bird - worm - fly - peep
snake - mice - slither hsss


Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal interface {
	eat()
	move()
	speak()
	getName() string
}

type cow struct {
	name string
}
type snake struct {
	name string
}
type bird struct {
	name string
}

func (c cow) getName() string {
	return c.name
}

func (s snake) getName() string {
	return s.name
}

func (b bird) getName() string {
	return b.name
}

func (c cow) eat() {
	fmt.Printf("%s eats grass\n", c.name)
}

func (c cow) move() {
	fmt.Printf("%s walks\n", c.name)
}

func (c cow) speak() {
	fmt.Printf("%s moos\n", c.name)
}

func (s snake) eat() {
	fmt.Printf("%s eats mice\n", s.name)
}

func (s snake) move() {
	fmt.Printf("%s slithers\n", s.name)
}

func (s snake) speak() {
	fmt.Printf("%s hsss\n", s.name)
}

func (b bird) eat() {
	fmt.Printf("%s eats worms\n", b.name)
}

func (b bird) move() {
	fmt.Printf("%s flys\n", b.name)
}

func (b bird) speak() {
	fmt.Printf("%s peeps\n", b.name)
}

func main() {
	listOfAnimals := []animal{}
	repeat := true

	for repeat == true {
		fmt.Printf("\nEnter the request type ('query'), animal name and what data you want ('eat', 'move', 'speak')")
		fmt.Println("\nor 'newanimal', name and type ('cow','bird','snake') to add a new one\n")
		fmt.Print("> ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		userInput := input.Text()
		userInputSlice := strings.Split(userInput, " ")

		if len(userInputSlice) != 3 {
			fmt.Println("Invalid request")
		}

		switch userInputSlice[0] {
		case "newanimal":
			switch userInputSlice[2] {
			case "cow":
				listOfAnimals = append(listOfAnimals, cow{name: userInputSlice[1]})
				fmt.Println("Created it!")
			case "snake":
				listOfAnimals = append(listOfAnimals, snake{name: userInputSlice[1]})
				fmt.Println("Created it!")
			case "bird":
				listOfAnimals = append(listOfAnimals, bird{name: userInputSlice[1]})
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid request type")
			}

		case "query":
			if len(listOfAnimals) != 0 {
				for _, animal := range listOfAnimals {
					if animal.getName() == userInputSlice[1] {
						if userInputSlice[2] == "move" {
							animal.move()
						} else if userInputSlice[2] == "eat" {
							animal.eat()
						} else if userInputSlice[2] == "speak" {
							animal.speak()
						}
					}
				}
			} else {
				fmt.Println("Animal list is empty! Add at least 1 animal before query")
			}
		default:
			fmt.Println("Invalid request type")
		}

		fmt.Printf("\nWant again? enter 'y' to repeat and anything else to stop\n> ")
		repeatInput := bufio.NewScanner(os.Stdin)
		repeatInput.Scan()
		repeatInputAns := strings.ToLower(repeatInput.Text())
		if repeatInputAns != "y" {
			repeat = false
		}
	}
}
