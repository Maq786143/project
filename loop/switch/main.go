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

	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("Excellent")
	case 'B':
		fmt.Println("Good")
	case 'C':
		fmt.Println("Average")
	case 'D':
		fmt.Println("Below Average")
	case 'F':
		fmt.Println("Fail")
	default:
		fmt.Println("Invalid grade")
	}
}
