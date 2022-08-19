/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. 
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line. 

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. 
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, 
your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	
)

type Person struct {
	fName string
	lName string
	}

var persons []Person
var prsn Person
var file_line []string


func main() {

//open txt file
	fmt.Printf("\nEnter the file name (do not include .txt). File should be in the same folder as .go file: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fileName := input.Text() + ".txt"

	file, err := os.Open(fileName)
 	if err != nil {
        log.Fatalf("failed to open your file")
  
    }

//read the file

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
        file_line = append(file_line, scanner.Text())
    	}
   	file.Close()

//split each string for names and create new instance of Person
	
	for _, each_ln := range file_line {
 		splitLine := strings.Split(each_ln, " ")
		prsn.fName = splitLine[0]
		prsn.lName = splitLine[1]
		persons = append(persons, prsn)
	}

//Print person slice
	for _, each_prsn := range persons {
        fmt.Println(each_prsn)
	}
}

