package main

import "fmt"

func doubleValue(x *int) {
	*x = *x * 2 // Multiply the value at the memory address pointed to by x by 2 and update the value
}

func main() {
	var y int = 10                         // Declare a variable y of type int and assign it the value 10
	fmt.Println("Initial value of y: ", y) // Print the initial value of y

	doubleValue(&y) // Pass the memory address of y to the doubleValue function

	fmt.Println("Value of y after function call: ", y) // Print the value of y after the function call, which will be 20
}
