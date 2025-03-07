package main

import "fmt"

func main() {
	count := 0
	var limit int
	fmt.Print("Enter the number of times to run: ")
	fmt.Scan(&limit)

	for count < limit {
		fmt.Println("Running...", count+1)
		count++
	}

	fmt.Println("Loop ended.")
}
