// golang "inheritance" using embedded types
package main

import (
	"fmt"
)

type Class struct {
	Professor string
	TA        string
	Size      int
	Location  string
}

type ITMD469 struct {
	Class               // embedd Class struct, inheriting class fields
	ProgrammingLanguage string
}

func main() {

	// create new ITM469 struct and set Class fields
	itm469 := ITMD469{
		Class: Class{
			Professor: "Professor Billy",
			TA:        "Travis",
			Size:      10,
			Location:  "Online",
		},
		ProgrammingLanguage: "Golang",
	}

	fmt.Println(itm469.TA)
	fmt.Println(itm469.ProgrammingLanguage)

}
