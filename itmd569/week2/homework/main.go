package main

import (
	"fmt"
	"strconv"
	"homework/math"
)

func main(){
	fmt.Println("Calculator program")

	fmt.Println("Please select which operation do you want to perform?")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	var operation int
	fmt.Scanln(&operation)
	if operation < 1 || operation > 4 {
		fmt.Println("Please enter valid operation!")
		return
	}

	fmt.Println("Please enter two numbers:")
	var num1Str, num2Str string
	fmt.Print("Number 1: ")
	fmt.Scanln(&num1Str)
	fmt.Print("Number 2: ")
	fmt.Scanln(&num2Str)
	num1Float, err1 := strconv.ParseFloat(num1Str, 64)
	num2Float, err2 := strconv.ParseFloat(num2Str, 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Please enter two valid numbers.")
		return
	}

	switch operation {
	case 1:
		result := math.Add(int(num1Float), int(num2Float))
		fmt.Printf("The result of addition is: %d\n", result)
	case 2:
		result := math.Subtract(int(num1Float), int(num2Float))
		fmt.Printf("The result of subtraction is: %d\n", result)
	case 3:
		result := math.Multiply(int(num1Float), int(num2Float))
		fmt.Printf("The result of multiplication is: %d\n", result)
	case 4:
		result, err := math.Divide(num1Float, num2Float)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Printf("The result of division is: %.2f\n", result)
	default:
		return
	}
}