package main

import (
	"fmt"
	"math/rand"
	"time"
)

type customer struct {
	id           string
	numPurchases int
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}

func countingSort(customers []customer, max int) []customer {
	counts := make([]int, max)
	for _, c := range customers {
		counts[c.numPurchases]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	sorted := make([]customer, len(customers))
	for i := len(customers) - 1; i >= 0; i-- {
		sorted[counts[customers[i].numPurchases]-1] = customers[i]
		counts[customers[i].numPurchases]--
	}

	return sorted
}

func makeRandomSlice(numCustomers, max int) []customer {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	customers := make([]customer, numCustomers)
	for i := 0; i < numCustomers; i++ {
		customers[i] = customer{
			id:           fmt.Sprintf("C-%d", i),
			numPurchases: random.Intn(max),
		}
	}

	return customers
}

// Prints the customers using min(customers, numCustomers).
func printSlice(customers []customer, numCustomers int) {
	if len(customers) < numCustomers {
		for _, c := range customers {
			fmt.Printf("%s - %d\n", c.id, c.numPurchases)
		}
	} else {
		for i := 0; i < numCustomers; i++ {
			c := customers[i]
			fmt.Printf("%s - %d", c.id, c.numPurchases)
		}
	}
}

// Checks if the customers are sorted by numPurchases.
func checkSorted(customers []customer) {
	for i := 0; i < len(customers)-1; i++ {
		if customers[i].numPurchases > customers[i+1].numPurchases {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted.")
}
