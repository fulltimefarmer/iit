package math

import "errors"

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Subtract(num1 int, num2 int) int {
	return num1 - num2
}

func Multiply(num1 int, num2 int) int {
	return num1 * num2
}

func Divide(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Cannot divide by zero!")
	}
	return num1 / num2, nil
}