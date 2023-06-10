package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	addResult := Add(2, 2)
	fmt.Println(addResult)

	divideResult, err := Divide(1, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(divideResult)

}

// one return
func Add(num1 int, num2 int) (result int) {
	result = num1 + num2
	return result
}

// two returns
func Divide(num1 float64, num2 float64) (result float64, err error) {

	if num1 == 0 {
		return 0, errors.New("Cannot divide by zero!")
	}

	result = num1 / num2
	return result, err

}
