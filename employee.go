package main

import "fmt"

type Employee struct {
	name            string
	id              int
	hourly_pay_rate float32 `default:"100"`
	hours_worked    int     `default:"10"`
	employer_cost   float32 `default:"100"`
	has_commision   bool
	commision       float32 `default:"100"`
	no_of_projects  int
}

func (Employee) ComputePayout() (payout float32) {
	panic("compute payout func not implemented")
}

func main() {
	fmt.Println("Program to Compute Employee payout :- ")
}
