package main

import (
	"fmt"
)

func main() {

	// store the user input into variables that we can use later
	var firstName string
	var lastName string

	// we will go over what & means in front of the variables,
	// just know for now that its used to put values into the variables we declared up top using Scanln

	fmt.Println("Please enter your first name")
	fmt.Scanln(&firstName)

	fmt.Println("Please enter your last name")
	fmt.Scanln(&lastName)

	fmt.Println("Your name is " + firstName + " " + lastName)
}
