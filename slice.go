/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var repeat string

func main() {
	s := make([]int, 0, 3)

	for repeat != "x" {

		fmt.Printf("\ntype only the integer value or X to stop: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if strings.ToLower(input.Text()) != "x" {
			userInput, _ := strconv.ParseInt(input.Text(), 10, 0)
			userInputInt := int(userInput)

			if (userInputInt == 0 && input.Text() == "0") || userInputInt != 0 {
				s = append(s, userInputInt)
				sort.Ints(s)
				fmt.Println(s)
			}

		} else {
			break
		}

	}
}
