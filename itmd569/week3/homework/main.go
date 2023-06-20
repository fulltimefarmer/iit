package main

import "fmt"

type Employee struct {
    Name string
    Age int
    Position string
    MonthlySalary float64
    TotalCost     float64
}

func (e Employee) calcYearlyCost() float64 {
    e.TotalCost = 12.0 * e.MonthlySalary
    return e.TotalCost
}

type Facilities struct {
    Type string
    RentCost float64
    MaintenanceCost float64
    TotalCost       float64
}

func (f Facilities) calcYearlyCost() float64 {
    f.TotalCost = 12.0 * f.RentCost + f.MaintenanceCost
    return f.TotalCost
}

type Government struct {
    CorporateTax float64
    TotalEarnings float64
    TotalCosts float64
}

func (g Government) calcYearlyCosts() float64 {
    taxableAmount := g.TotalEarnings - g.TotalCosts
    taxableAmount -= taxableAmount * g.CorporateTax
    if taxableAmount < 0 {
        return 0
    }
    return taxableAmount
}

type TotalCosts struct {
    Employee
    Facilities
    Government
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

    fac := Facilities {
        Type:            "Downtown Office",
        RentCost:        100000,
        MaintenanceCost: 1000,
    }
    facCost := calcYearlyCost(fac)
    fmt.Printf("%s costs the company $%.0f per year.\n", fac.Type, facCost)

    totalCost := empCost + facCost
    fmt.Printf("Total company costs: $%.0f per year.\n", totalCost)

    gov := Government {
        CorporateTax: 0.21,
        TotalEarnings: 1000000.0,
        TotalCosts: totalCost,
    }

    govCost := gov.calcYearlyCosts()
    fmt.Printf("The government costs: $%.2f per year.\n", govCost)

    total := TotalCosts {
        Employee: emp,
        Facilities: fac,
        Government: gov,
    }

    totalProfit := total.TotalEarnings - total.TotalCosts - govCost
    fmt.Printf("The total profit for the company: $%.2f per year.\n", totalProfit)
}
