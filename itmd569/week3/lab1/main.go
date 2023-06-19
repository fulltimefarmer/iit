package main

import "fmt"

type Employee struct {
    Name string
    Age int
    Position string
    MonthlySalary float64
}

func (e Employee) calcYearlyCost() float64 {
    return 12.0 * e.MonthlySalary
}

func (e Employee) printInfo() {
    fmt.Printf("Employee Info\n\nName: %s\nAge: %d\nPosition: %s\nMonthly Salary: $%.0f\n", e.Name, e.Age, e.Position, e.MonthlySalary)
}

func main () {
	employee := Employee {
		Name: "Max",
		Age: 36,
		Position: "Software Engineer",
		MonthlySalary: 10000.0,
	}
	
	employee.printInfo()
    fmt.Printf("Yearly Salary: $%.0f\n", employee.calcYearlyCost())
}
