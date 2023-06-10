package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var input string
	fmt.Println("Enter anything except for 569")
	fmt.Scanln(&input)

	// error checking conventions in go
	// if the error returned from the function is not nil, it has a value, and something went wrong
	err := checkInput(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Thanks for following directions!")

}

func checkInput(input string) (err error) {
	if input == "569" {
		err = errors.New("You had one job") // change err from nil to having a value
		return err
	}
	return err // still return err, even though its nil. nil means nothing went wrong!
}
