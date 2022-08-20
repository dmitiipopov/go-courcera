/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var floatValue float64

func inputFloat() float64 {
	//input function, converts to float64 and asks to p[rovide correct value
	inputErr := true
	for inputErr == true {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		inputValue := input.Text()
		if s, err := strconv.ParseFloat(inputValue, 64); err == nil {
			floatValue = s
			inputErr = false
		} else {
			fmt.Println("Try again! only float64")
			inputErr = true
		}
	}
	return floatValue
}

func main() {
	//arguments for f(t) anf its preparation for next calc
	fmt.Println("Enter A (float64)")
	a := inputFloat()
	fmt.Println("Enter v0 (float64)")
	v0 := inputFloat()
	fmt.Println("Enter s0 (float64)")
	s0 := inputFloat()
	fmt.Printf("Arguments values for function: \nA = %g, v0 = %g, s0 = %g\n", a, v0, s0)
	preparedFn := GenDisplaceFn(a, v0, s0)
	//request to re-calculate with new time value
	repeat := "y"
	for strings.ToLower(repeat) == "y" {
		fmt.Println("Enter time")
		time := inputFloat()
		fmt.Printf("Function value = %g", preparedFn(time))
		fmt.Println("\nRepeat? Type 'y' if yes or anything else to stop")
		inputRepeat := bufio.NewScanner(os.Stdin)
		inputRepeat.Scan()
		repeat = inputRepeat.Text()
	}
}

//function to prepare f(t)
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fnT := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fnT
}
