package main

import "fmt"

type Employee struct {
	name            string
	id              int
	hourly_pay_rate float32 `default:"10"`
	hours_worked    int
	employer_cost   float32 `default:"100"`
	has_commision   bool
	commision       float32 `default:"100"`
	no_of_projects  int
}

func (Employee) ComputePayout() (payout float32) {
	panic("Compute payout func not implemented")
}

func main() {
	fmt.Println("Program to Compute Employee payout :- ")
	employee := Employee{
		name:         "Employee Name",
		id:           1,
		hours_worked: 10,
	}
	fmt.Println("The Payout of the employee is", employee.ComputePayout())

}
