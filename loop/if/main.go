package main

import "fmt"

func main() {
	defer fmt.Println("First Deferred Statement")
	defer fmt.Println("Second Deferred Statement")
	defer fmt.Println("Third Deferred Statement")

	fmt.Println("Function Execution Started")
	fmt.Println("Performing some operations...")

	num := 42
	ptr := &num

	fmt.Println("Original Value:", num)
	fmt.Println("Pointer Address:", ptr)
	fmt.Println("Value from Pointer:", *ptr)

	*ptr = 100
	fmt.Println("Modified Value:", num)

	if num > 1 {
		fmt.Println("The number is greater than 1.")
	} else if num == 1 {
		fmt.Println("The number is 1.")
	} else {
		fmt.Println("The number is zero or negative.")
	}
}
