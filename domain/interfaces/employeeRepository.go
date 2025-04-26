package interfaces

import employeeModel "github.com/brianmorais/go-user-importation/domain/models/employee"

type IEmployeeRepository interface {
	GetEmployees() (*employeeModel.Employees, error)
	GetEmployeesView() (*employeeModel.EmployeesView, error)
	UpdateEmployee(employee *employeeModel.Employee) (int64, error)
	CreateEmployee(employee *employeeModel.Employee) (int64, error)
	GetEmployeeRoles() (*employeeModel.EmployeeRoles, error)
	GetEmployeeByRegistrationNumber(registrationNumber string) (*employeeModel.Employee, error)
}
