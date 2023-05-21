package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type EmployeeComputePayrollTestSuite struct {
	suite.Suite
	Employee Employee
}

// Setup is called before every test
func (suite *EmployeeComputePayrollTestSuite) SetupTest() {
	suite.Employee = Employee{
		name: "Employee Name",
		id:   1,
	}
}

func TestEmployeePayout(t *testing.T) {
	tests := new(EmployeeComputePayrollTestSuite)
	suite.Run(t, tests)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_Returns_Float() {
	//Check Payout returns a Float
	assert.Equal(suite.T(), reflect.TypeOf(suite.Employee.ComputePayout()).String(), "float32")
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_NoCommission_NoHours() {
	//Check if payout is correctly computed in case of no commission and no hours worked
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), 100)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_NoCommission() {
	//Check if payout is correctly computed in case of no commission and 10 hours worked
	suite.Employee.hours_worked = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), 2000)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_WithCommission() {
	//Check if payout is correctly computed in case of commission and 10 hours worked
	suite.Employee.hours_worked = 10
	suite.Employee.has_commision = true
	suite.Employee.commision = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), 3000)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_Commission_Disabled() {
	//Check if payout is correctly computed in case of commission and 10 hours worked
	suite.Employee.hours_worked = 10
	suite.Employee.has_commision = false
	suite.Employee.commision = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), 2000)
}
