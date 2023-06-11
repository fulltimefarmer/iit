package main

import (
	"fmt"
	"strconv"
	"lab2/math"
)

func main() {
	var num1Str, num2Str string

	fmt.Println("Please enter two numbers:")
	fmt.Print("Number 1: ")
	fmt.Scanln(&num1Str)
	fmt.Print("Number 2: ")
	fmt.Scanln(&num2Str)

	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)

	if err1 != nil || err2 != nil {
		fmt.Println("Please enter valid numbers.")
		return
	}

	result := math.Add(num1, num2)

	if result % 2 == 0 {
		fmt.Println("The sum of the two numbers is even.")
	} else {
		fmt.Println("The sum of the two numbers is odd.")
	}
}