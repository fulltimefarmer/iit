package calc

func DetermineType(num1 int, num2 int) (isEven bool) {

	sum := num1 + num2

	if sum%2 == 0 {
		return true
	} else {
		return false
	}

}
