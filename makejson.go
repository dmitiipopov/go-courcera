/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	data := make(map[string]string)

	fmt.Printf("\nEnter name: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	data["Name"] = input.Text()

	fmt.Printf("\nEnter address: ")
	input = bufio.NewScanner(os.Stdin)
	input.Scan()
	data["Address"] = input.Text()

	json, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\n")
		fmt.Println(string(json))
		fmt.Println("\n")
	}
}
