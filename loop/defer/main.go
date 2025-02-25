package main

import "fmt"

func main() {
	defer fmt.Println("First Deferred Statement")
	defer fmt.Println("Second Deferred Statement")
	defer fmt.Println("Third Deferred Statement")

	fmt.Println("Function Execution Started")
	fmt.Println("Performing some operations...")
}
