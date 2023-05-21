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

func (suite *EmployeeComputePayrollTestSuite) SetupTest() {
	suite.Employee = Employee{Name: "Employee Name", Id: 1}
}

func TestEmployeePayout(t *testing.T) {
	tests := new(EmployeeComputePayrollTestSuite)
	suite.Run(t, tests)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_Returns_Float() {
	assert.Equal(suite.T(), reflect.TypeOf(suite.Employee.ComputePayout()).String(), "float32")
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_NoCommission_NoHours() {
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), 100)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_NoCommission() {
	suite.Employee.HoursWorked = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), 200)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_WithCommission() {
	suite.Employee.HoursWorked = 10
	suite.Employee.HasCommision = true
	suite.Employee.Commision = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), 300)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_Commission_Disabled() {
	suite.Employee.HoursWorked = 10
	suite.Employee.HasCommision = false
	suite.Employee.Commision = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), 200)
}
