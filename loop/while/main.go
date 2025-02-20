package main

import "fmt"

func main() {
	fmt.Println("Countdown starts:")

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	fmt.Println("Countdown finished!")
}
