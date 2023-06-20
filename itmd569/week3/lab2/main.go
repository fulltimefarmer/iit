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

type Facilities struct {
    Type string
    RentCost float64
    MaintenanceCost float64
}

func (f Facilities) calcYearlyCost() float64 {
    return 12.0 * f.RentCost + f.MaintenanceCost
}

type Cost interface {
    calcYearlyCost() float64
}

func calcYearlyCost(c Cost) float64 {
    return c.calcYearlyCost()
}

func main () {
	emp := Employee {
		Name: "Max",
		Age: 36,
		Position: "Software Engineer",
		MonthlySalary: 10000.0,
	}
	empCost := calcYearlyCost(emp)
    fmt.Printf("%s costs the company $%.0f per year.\n", emp.Name, empCost)

    fac := Facilities{
        Type:            "Downtown Office",
        RentCost:        100000,
        MaintenanceCost: 1000,
    }
    facCost := calcYearlyCost(fac)
    fmt.Printf("%s costs the company $%.0f per year.\n", fac.Type, facCost)

    totalCost := empCost + facCost
    fmt.Printf("Total company costs: $%.0f per year.\n", totalCost)
}
