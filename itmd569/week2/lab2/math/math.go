package math

import ("errors")

func Add(num1 int, num2 int) (result int) {
	result = num1 + num2
	return result
}

func Divide(num1 float64, num2 float64) (result float64, err error) {
	if num1 == 0 {
		return 0, errors.New("Cannot divide by zero!")
	}
	result = num1 / num2
	return result, err
}