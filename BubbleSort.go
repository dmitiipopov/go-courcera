/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/


package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

var userslice []int
var sliceToSort []int
var inputRaw string
var intInputValue int 

const (
	SliceLen int = 10
)

func main() {
	bubbleSort(intInput())
	fmt.Println("Sorted result:")
	fmt.Println(userslice)
}

func inputCheck(checkInput *string, checkintInputValue *int) bool {
	var inputCheckResult bool
	 if (*checkInput == "0"  &&  *checkintInputValue == 0) || *checkintInputValue != 0 {
		inputCheckResult = true
	} else {
		inputCheckResult = false
	}	
	return inputCheckResult
}

func intInput() []int {
	for i:=0; i < SliceLen; i++ {
		inputCheckR := false
		for inputCheckR == false {
			fmt.Printf("Please enter %d int element: ", i+1)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			inputRaw = input.Text()
			intInputValue,_ = strconv.Atoi(input.Text())
			inputCheckR = inputCheck(&inputRaw, &intInputValue)
		}
		userslice = append(userslice, intInputValue)
	}
	return userslice
}

func bubbleSort(usersliceSort []int)  {
	for i:=0; i< len(usersliceSort)-1; i++ {
		swap(usersliceSort, i)
    }
}

	
func swap(usersliceSwap []int, i int)  {
      for j:=0; j < len(usersliceSwap)-i-1; j++ {
         if (usersliceSwap[j] > usersliceSwap[j+1]) {
            usersliceSwap[j], usersliceSwap[j+1] = usersliceSwap[j+1], usersliceSwap[j]
        }
    }
}