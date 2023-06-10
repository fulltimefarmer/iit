package main

import (
	"fmt"

	"github.com/tsmith-41/go-examples/Week2/Day2/Lab/calc"
)

// An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.
func main() {
	fmt.Println("Enter the first integer")
	var num1String int
	fmt.Scanln(&num1String)

	fmt.Println("Enter the second integer")
	var num2String int
	fmt.Scanln(&num2String)

	result := calc.DetermineType(num1String, num2String)

	if result {
		fmt.Println("Sum is even")
	} else {
		fmt.Println("Sum is odd")
	}

}
