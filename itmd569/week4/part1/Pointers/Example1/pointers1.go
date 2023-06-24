package main

import "fmt"

// Helpful tips with pointers:

// Tip 1: Understanding the address-of operator (&)
// The ampersand symbol (&) is used to get the memory address of a variable.
// You can remember this by associating & with "Address" in "Memory Address."
// When you see &x, read it as "address of variable x."

// Tip 2: Understanding the dereference operator (*)
// The asterisk symbol (*) is used to declare a pointer type or to access the value stored at a memory address.
// Visualize the * symbol as a point or arrow, indicating that it's pointing to something.
// For example, *int is a pointer to an integer variable.
// When you see *x, read it as "the value pointed to by x."

func main() {
	var x int = 10 // Declare a variable x of type int and assign it the value 10
	var p *int     // Declare a pointer to an int

	p = &x // Assign the memory address of x to p using the address-of operator &

	fmt.Println("Value of x: ", x)                // Print the value of x
	fmt.Println("Address of x: ", &x)             // Print the memory address of x using the address-of operator &
	fmt.Println("Value of p (address of x): ", p) // Print the value of p, which is the memory address of x
	fmt.Println("Value pointed to by p: ", *p)    // Use the dereference operator * to access the value stored at the memory address pointed to by p
}
