/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal:
1) the food that it eats, 2) its method of locomotion, and
3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal - Food eaten - Locomotion method - Spoken sound
cow - grass - walk - moo
bird - worms - fly - peep
snake - mice - slither - hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

var typeRequest string
var inputString string

func (a *animal) Eat() string {
	//	fmt.Printf("\nthis animal eats %s ", a.food)
	fmt.Printf(a.food)
	return a.food
}

func (a *animal) Move() string {
	fmt.Printf("\nthis animal locomotions is %s ", a.locomotion)
	return a.locomotion
}

func (a *animal) Speak() string {
	fmt.Printf("\nthis animal speaks like '%s' ", a.noise)
	return a.noise
}

func inputCheck(checkInput []string) bool {
	var inputCheckResult bool
	if len(checkInput) == 2 {
		if (checkInput[0] == strings.ToLower("bird") || checkInput[0] == strings.ToLower("cow") || checkInput[0] == strings.ToLower("snake")) && (checkInput[1] == strings.ToLower("eats") || checkInput[1] == strings.ToLower("moves") || checkInput[1] == strings.ToLower("speaks")) {
			inputCheckResult = true
		} else if (checkInput[0] != strings.ToLower("bird") || checkInput[0] != strings.ToLower("cow") || checkInput[0] != strings.ToLower("snake")) && (checkInput[1] == strings.ToLower("eats") || checkInput[1] == strings.ToLower("moves") || checkInput[1] == strings.ToLower("speaks")) {
			inputCheckResult = false
			fmt.Println("Incorrect animal!")
		} else if (checkInput[0] == strings.ToLower("bird") || checkInput[0] == strings.ToLower("cow") || checkInput[0] == strings.ToLower("snake")) && (checkInput[1] != strings.ToLower("eats") || checkInput[1] != strings.ToLower("moves") || checkInput[1] != strings.ToLower("speaks")) {
			inputCheckResult = false
			fmt.Println("Incorrect data request!")
		}
	} else {
		inputCheckResult = false
		fmt.Println("Incorrect request")
	}
	return inputCheckResult
}

func spiltInput(strInput string) []string {
	inputWords := strings.Split(strings.TrimSuffix(strInput, "\n"), " ")
	return inputWords
}

func input() string {
	repeat := true
	for repeat == true {
		fmt.Printf("\nEnter the animal name (cow/bird/snake) and what data you want (eats, moves, speaks): > \n")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		inputString = input.Text()
		if inputCheck(spiltInput(inputString)) == true {
			repeat = false
		}
	}
	return inputString
}

func main() {
	cow := animal{
		name:       "cow",
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	bird := animal{
		name:       "bird",
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	snake := animal{
		name:       "snake",
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
	repeat := true
	for repeat == true {
		inputSlice := spiltInput(input())
		animalName, dataRequested := inputSlice[0], inputSlice[1]
		switch dataRequested {
		case "eats":
			switch animalName {
			case "cow":
				cow.Eat()
			case "bird":
				bird.Eat()
			case "snake":
				snake.Eat()
			}
		case "moves":
			switch animalName {
			case "cow":
				cow.Move()
			case "bird":
				bird.Move()
			case "snake":
				snake.Move()
			}
		case "speaks":
			switch animalName {
			case "cow":
				cow.Speak()
			case "bird":
				bird.Speak()
			case "snake":
				snake.Speak()
			}
		}
		fmt.Printf("\n\nWant again? enter 'y' to repeat and anything else to stop) > ")
		repeatInput := bufio.NewScanner(os.Stdin)
		repeatInput.Scan()
		repeatInputAns := strings.ToLower(repeatInput.Text())
		if repeatInputAns != "y" {
			repeat = false
		}
	}
}
