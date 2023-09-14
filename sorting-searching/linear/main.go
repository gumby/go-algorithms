package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var numItems, max int
	fmt.Printf("# of items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)

	var entry string
	fmt.Printf("Target: ")
	fmt.Scanln(&entry)

	if entry == "" {
		return
	}
	target, err := strconv.Atoi(entry)
	if err != nil {
		log.Fatalf("Invalid entry for search")
	}

	index, tries := linearSearch(slice, target)
	if index >= 0 {
		fmt.Printf("values[%d] = %d, %d tests\n", index, target, tries)
	} else {
		fmt.Printf("Target %d not found, %d tests", target, tries)
	}
}

// Perform linear search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func linearSearch(slice []int, target int) (index, numTests int) {
	for i, candidate := range slice {
		numTests++
		if candidate == target {
			return i, numTests
		}
	}
	return -1, numTests
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
	fmt.Printf("[ ")
	if len(items) < numItems {
		for _, item := range items {
			fmt.Printf("%d ", item)
		}
	} else {
		for i := 0; i < numItems; i++ {
			fmt.Printf("%d ", items[i])
		}
	}
	fmt.Println("]")
}
