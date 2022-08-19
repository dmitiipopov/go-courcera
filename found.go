/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//var aCheck bool

func main() {
	//loop for 1+ interaction
	repeat := "y"
	for strings.ToLower(repeat) == "y" {
		aCheck := false
		//fmt.Scan(&userString)
		fmt.Printf("\ntype the string: ")
		text := bufio.NewScanner(os.Stdin)
		text.Scan()
		userString := text.Text()
		userStringLen := len(userString)
		// checking 1st and last symbols and if Yes => checking A between them. Break if 1st A found
		if strings.ToLower(string([]rune(userString)[0])) == "i" && strings.ToLower(string([]rune(userString)[userStringLen-1])) == "n" {
			for i := 0; i < userStringLen-1; i++ {
				if strings.ToLower(string([]rune(userString)[i])) == "a" {
					fmt.Printf("\nFound!\n")
					aCheck = true
					break
				}
			}
		}

		if aCheck == false {
			fmt.Printf("\nNot found!\n")
		}
		// ask for additional loop

		fmt.Printf("\nWant again? Print 'y' if yes and any other to stop: ")
		answer := bufio.NewScanner(os.Stdin)
		answer.Scan()
		repeat = answer.Text()

	}
}
