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
    return 12.0 * e.MonthlySalary
}

type Facilities struct {
    Type string
    RentCost float64
    MaintenanceCost float64
    TotalCost       float64
}

func (f Facilities) calcYearlyCost() float64 {
    return 12.0 * (f.RentCost + f.MaintenanceCost)
}

type Government struct {
    CorporateTax float64
    TotalEarnings float64
    TotalCosts float64
}

func (g Government) calcYearlyCosts() float64 {
    return g.TotalEarnings * 0.21
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
    emp.TotalCost = empCost
    fmt.Printf("%s costs the company $%.0f per year.\n", emp.Name, empCost)

    fac := Facilities {
        Type:            "Downtown Office",
        RentCost:        100000,
        MaintenanceCost: 1000,
    }
    facCost := calcYearlyCost(fac)
    fac.TotalCost = facCost
    fmt.Printf("%s costs the company $%.0f per year.\n", fac.Type, facCost)

    totalCompanyCost := empCost + facCost
    fmt.Printf("Total company costs: $%.0f per year.\n", totalCompanyCost)

    gov := Government {
        CorporateTax: 0.21,
        TotalEarnings: 10000000.0,
    }

    govCost := gov.calcYearlyCosts()
    gov.TotalCosts = govCost
    fmt.Printf("The government costs: $%.2f per year.\n", govCost)

    totalCosts := TotalCosts {
        Employee: emp,
        Facilities: fac,
        Government: gov,
    }

    profit := totalCosts.Government.TotalEarnings - 
        (totalCosts.Employee.TotalCost) - 
        (totalCosts.Facilities.TotalCost) - 
        (totalCosts.Government.TotalCosts)

    fmt.Printf("Total Profit: $%.2f \n", profit)
}
