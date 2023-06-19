package main

import "fmt"

// structs are a collection of fields
// they are useful for grouping information together
// can be thought of as a "class"
type Person struct {
	name       string
	age        int
	occupation string
}

func main() {
	// using structs
	// create a new empty person struct
	teacherAssistant := Person{}

	// set fields
	teacherAssistant.name = "Travis"
	teacherAssistant.age = 22
	teacherAssistant.occupation = "Platform Engineer"

	// now teacherAssistant variable holds references to those subfields (name, age, occupation)
	// can also declare like this
	newPerson := Person{name: "Travis", age: 22, occupation: "Platform Engineer"}
	fmt.Println(newPerson.name)
	fmt.Println(newPerson.age)
	fmt.Println(newPerson.occupation)
}
