/*
 Write a program to sort an array of integers. 
 The program should partition the array into 4 parts, each of which is sorted by a different goroutine. 
 Each partition should be of approximately equal size. 
 Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func splitSortList(intList []int, partSize float64, partLength int, i int, wg *sync.WaitGroup, c chan []int) {
    list := intList[i*partLength : (i*partLength)+partLength]
    sort.Ints(list)
    c <- list
    fmt.Println("Sorted part: n = ", i, " List", list)
    wg.Done()
}

func stringToInt(strList string, partSize float64) []int {
    var intList []int
    for _, s := range strings.Fields(strList) {
        i, err := strconv.Atoi(s)
        if err == nil {
            intList = append(intList, i)
        }
    }
    return intList
}

func main() {
    fmt.Println("Enter a list of integers: ")
    reader := bufio.NewReader(os.Stdin)
    numList, _ := reader.ReadString('\n')

    partSize := 4.0
    intList := stringToInt(numList, partSize)
    floatListLen := float64(len(intList))
    partLength := int(math.Ceil(float64(floatListLen / partSize)))

    c := make(chan []int, int(partSize))
    semiSortedList := []int{}

    var wg sync.WaitGroup

    for i := 0; i < int(partSize); i++ {
        wg.Add(1)
        go splitSortList(intList, partSize, partLength, i, &wg, c)
    }
    wg.Wait()
    close(c)
    for l := range c {
        semiSortedList = append(semiSortedList, l...)
    }

    fmt.Println("after sort", semiSortedList)
}