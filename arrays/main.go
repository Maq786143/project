package main

import "fmt"

func main() {
	// Initialize an array of integers
	numbers := [5]int{10, 20, 30, 40, 50}

	// Print the first element
	fmt.Println("First Element:", numbers[0])

	// Loop through the array using range to print each element and its index
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Optionally, handle cases where the array is empty or dynamically adjust based on length
	if len(numbers) > 0 {
		fmt.Println("Array is not empty")
	} else {
		fmt.Println("Array is empty")
	}
}
