package main

import "fmt"

func main() {
	var x int = 10 // Declare a variable x of type int and assign it the value 10
	var p1 *int    // Declare a pointer p1 to an int
	var p2 *int    // Declare a pointer p2 to an int

	p1 = &x // Assign the memory address of x to p1
	p2 = p1 // Assign the value of p1 (memory address of x) to p2. Now p2 also points to x.

	fmt.Println("Value of x: ", x)                  // Print the value of x
	fmt.Println("Address of x: ", &x)               // Print the memory address of x using the address-of operator &
	fmt.Println("Value of p1 (address of x): ", p1) // Print the value of p1, which is the memory address of x
	fmt.Println("Value pointed to by p1: ", *p1)    // Use the dereference operator * to access the value stored at the memory address pointed to by p1
	fmt.Println("Value of p2 (p1's value): ", p2)   // Print the value of p2, which is the value of p1 (memory address of x)
	fmt.Println("Value pointed to by p2: ", *p2)    // Use the dereference operator * to access the value stored at the memory address pointed to by p2
}
