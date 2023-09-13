package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var numItems, max int
	fmt.Printf("# of items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 10)
	fmt.Println("----")
	quicksort(slice)
	printSlice(slice, 10)
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

// Quicksort sorts the slice using the quicksort algorithm.
func quicksort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	p := partition(slice)
	quicksort(slice[0:p])
	quicksort(slice[p+1:])
}

// Partition the slice into two partitions and return the pivot position.
func partition(slice []int) int {
	low := 0
	high := len(slice) - 1

	pivot := slice[high]
	i := low

	for j := low; j < high; j++ {
		if slice[j] < pivot {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}
	slice[i], slice[high] = slice[high], slice[i]
	return i
}
