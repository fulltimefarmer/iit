package main

import "fmt"

type Person struct {
	name       string
	age        string
	occupation string
}

func (p *Person) summary() {
	fmt.Println("This person has the following properties")
	fmt.Println("Name: " + p.name)
	fmt.Println("Age: " + p.age)
	fmt.Println("Occupation: " + p.occupation)
}

func main() {

	// create a new person struct
	newPerson := Person{name: "Travis", age: "22", occupation: "Platform Engineer"}

	// call summary method using newPerson struct type
	newPerson.summary()

	// show how we cant just call summary by itself
	// summary()

}
