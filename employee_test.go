package main

import (
	"github.com/mcuadros/go-defaults"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type EmployeeComputePayrollTestSuite struct {
	suite.Suite
	Employee *Employee
}

func (suite *EmployeeComputePayrollTestSuite) SetupTest() {
	suite.Employee = new(Employee)
	defaults.SetDefaults(suite.Employee)
}

func TestEmployeePayout(t *testing.T) {
	tests := new(EmployeeComputePayrollTestSuite)
	suite.Run(t, tests)
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_Returns_Float() {
	assert.Equal(suite.T(), reflect.TypeOf(suite.Employee.ComputePayout()).String(), "float32")
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_NoCommission_NoHours() {
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), float32(100))
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_NoCommission() {
	suite.Employee.HoursWorked = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), float32(200))
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_WithCommission() {
	suite.Employee.HoursWorked = 10
	suite.Employee.HasCommision = true
	suite.Employee.Commision = 10
	suite.Employee.NoOfProjects = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), float32(300))
}

func (suite *EmployeeComputePayrollTestSuite) Test_EmployeePayout_Commission_Disabled() {
	suite.Employee.HoursWorked = 10
	suite.Employee.HasCommision = false
	suite.Employee.Commision = 10
	assert.Equal(suite.T(), suite.Employee.ComputePayout(), float32(200))
}
