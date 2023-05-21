package main

import "fmt"

type Employee struct {
	Name          string
	Id            int
	HourlyPayRate float32 `default:"10"`
	HoursWorked   int
	EmployerCost  float32 `default:"100"`
	HasCommision  bool
	Commision     float32 `default:"100"`
	NoOfProjects  int
}

func (Employee) ComputePayout() (payout float32) {
	payout = Employee
}

func main() {
	fmt.Println("Program to Compute Employee payout :- ")
	employee := Employee{
		Name:        "Employee Name",
		Id:          1,
		HoursWorked: 10,
	}
	fmt.Println("The Payout of the employee is", employee.ComputePayout())
}
