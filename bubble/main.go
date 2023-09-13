package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(5, 1000)
	printSlice(slice, 5)
	fmt.Println("----")
	bubbleSort(slice)
	printSlice(slice, 5)
	checkSorted(slice)
}

// Make a slice containing random numbers [0, max).
func makeRandomSlice(numItems, max int) []int {
	// Initialize rand number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

// Prints the full slice or numItems, whichever is smaller.
func printSlice(items []int, numItems int) {
	if len(items) < numItems {
		for _, item := range items {
			fmt.Println(item)
		}
	} else {
		for i := 0; i < numItems; i++ {
			fmt.Println(items[i])
		}
	}
}

// Checks the slice to see if it is sorted.
func checkSorted(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted.")
}

// Sorts the slice using the bubble sort aglorithm.
func bubbleSort(slice []int) {
	for {
		swapped := false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
