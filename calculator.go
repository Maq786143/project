package calculator

import "fmt"

// Function to perform addition
func Add(a, b float64) float64 {
	return a + b
}

// Function to perform subtraction
func Subtract(a, b float64) float64 {
	return a - b
}

// Function to perform multiplication
func Multiply(a, b float64) float64 {
	return a * b
}

// Function to perform division
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}
