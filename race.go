/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.
*/

package main

import (
	"fmt"
	"time"
)

var curr int = 0

func incBy1() {
	for curr < 10 {
		prev := curr
		time.Sleep(time.Millisecond)
		curr = curr + 1
		fmt.Printf("Increase curr value %d by 1 is %d\n", prev, curr)
	}

}

func incBy2() {
	for curr < 10 {
		prev := curr
		time.Sleep(time.Millisecond)
		curr = curr + 2
		fmt.Printf("Increase curr value %d by 2 is %d\n", prev, curr)
	}

}

/*


Race condition is defined as a program where the outcome of the program depends on the interleaving,
while interleaving is non-deterministic.
In this code, the interleaving of incBy1 and incBy2 can change in every run
For example, in funtion incBy1(), before it starts to execute the line "curr = curr + 1",
curr might have already been updated by funtion incBy2()
*/

func main() {
	go incBy1()
	go incBy2()
	time.Sleep(2 * time.Second)
}
