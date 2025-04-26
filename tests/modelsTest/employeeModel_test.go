package modelsTest

import (
	"database/sql"
	"testing"

	employeeModel "github.com/brianmorais/go-user-importation/domain/models/employee"
	"github.com/brianmorais/go-user-importation/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestConvertEmployeeRegistrationNumberToInt64(t *testing.T) {
	employee := mocks.GetEmployeeMock()
	expectedValue := 1

	res := employee.RegistrationNumberToInt64()

	assert.Equalf(t, res, int64(expectedValue), "Valor esperado: %v, valor obtido: %v", expectedValue, res)
}

func TestMustIsTemporaryEmployee(t *testing.T) {
	employee := mocks.GetEmployeeMock()
	employee.RoleId.Int32 = 7
	expectedValue := true

	res := employee.IsTemporaryEmployee()

	assert.Equalf(t, res, expectedValue, "Valor esperado: %v, valor obtido: %v", expectedValue, res)
}

func TestShouldNotIsTemporaryEmployee(t *testing.T) {
	employee := mocks.GetEmployeeMock()
	expectedValue := false

	res := employee.IsTemporaryEmployee()

	assert.Equalf(t, res, expectedValue, "Valor esperado: %v, valor obtido: %v", expectedValue, res)
}

func TestMustConvertEmployeeRegistrationNumberToString(t *testing.T) {
	employee := mocks.GetEmployeeViewMock()
	expectedValue := "1"

	res := employee.RegistrationNumberToString()

	assert.Equalf(t, res, expectedValue, "Valor esperado: %v, valor obtido: %v", expectedValue, res)
}

func TestMustFindRoleByName(t *testing.T) {
	roles := mocks.GetEmployeeRoles()
	expectedRole := "Gerentes"

	res := roles.FindRoleByName(expectedRole)

	assert.NotEqualf(t, res, employeeModel.EmployeeRole{}, "Erro ao localizar role por nome")
}

func TestShouldNotFindRoleByName(t *testing.T) {
	roles := mocks.GetEmployeeRoles()
	expectedRole := "Usuario Temporário"

	res := roles.FindRoleByName(expectedRole)

	assert.Equalf(t, res, employeeModel.EmployeeRole{}, "Role não deveria ser encontrada")
}

func TestEmployeeMustEquals(t *testing.T) {
	employee := mocks.GetEmployeeMock()
	employeeView := mocks.GetEmployeeViewMock()

	res := employee.Equals(employeeView)

	assert.Truef(t, res, "Employee não é igual ao EmployeeView")
}

func TestEmployeeShouldNotEquals(t *testing.T) {
	employee := mocks.GetEmployeeMock()
	employeeView := mocks.GetEmployeeViewMock()
	employeeView.Cpf = sql.NullString{String: "987456321"}

	res := employee.Equals(employeeView)

	assert.Falsef(t, res, "Employee é igual ao EmployeeView e deveria ser diferente")
}

func TestMustFindHierarchByRoleId(t *testing.T) {
	roles := mocks.GetEmployeeRoles()
	roleId := int32(1)

	res := roles.FindHierarchByRoleId(roleId)

	assert.NotEqualf(t, res, 0, "Erro ao localizar role por nome")
}

func TestShouldNotFindHierarchByRoleId(t *testing.T) {
	roles := mocks.GetEmployeeRoles()
	roleId := int32(0)

	res := roles.FindHierarchByRoleId(roleId)

	assert.Equalf(t, int(res), 0, "Role não deveria ser encontrada")
}

func TestMustFindEmployeeById(t *testing.T) {
	employeeId := 3
	employees := mocks.GetEmployeesMock()

	res := employees.FindEmployeeById(int64(employeeId))

	assert.NotEqualf(t, res, employeeModel.Employee{}, "Deveria ter encontrado employees")
}

func TestShouldNotFindEmployeeById(t *testing.T) {
	employeeId := 40
	employees := mocks.GetEmployeesMock()

	res := employees.FindEmployeeById(int64(employeeId))

	assert.Equalf(t, res, employeeModel.Employee{}, "Não deveria ter encontrado employees")
}

func TestMustFindEmployeeByRegistrationNumber(t *testing.T) {
	registrationNumber := "3"
	employees := mocks.GetEmployeesMock()

	res := employees.FindEmployeeByRegistrationNumber(registrationNumber)

	assert.NotEqualf(t, res, employeeModel.Employee{}, "Não deveria ter encontrado employees")
}

func TestShouldNotFindEmployeeByRegistrationNumber(t *testing.T) {
	registrationNumber := "30"
	employees := mocks.GetEmployeesMock()

	res := employees.FindEmployeeByRegistrationNumber(registrationNumber)

	assert.Equalf(t, res, employeeModel.Employee{}, "Não deveria ter encontrado employees")
}

func TestGetNotExceptionAndActiveEmployeeRegistrationNumbers(t *testing.T) {
	employees := mocks.GetEmployeesMock()
	expectedAmount := 2

	res := employees.GetNotExceptionAndActiveEmployeeRegistrationNumbers()

	assert.Equalf(t, len(res), expectedAmount, "Tamanho experado do array: %v, tamanho obtido: %v", expectedAmount, len(res))
}

func TestGetEmployeeViewRegistrationNumbers(t *testing.T) {
	employees := mocks.GetEmployeesViewMock()

	res := employees.GetEmployeeViewRegistrationNumbers()

	assert.NotEqual(t, len(res), 0)
}

func TestGetEmployeesToDeactivate(t *testing.T) {
	employees := mocks.GetEmployeesMock()
	registrationNumbers := []string{"1", "2"}
	expectedAmount := len(registrationNumbers)

	res := employees.GetEmployeesToDeactivate(registrationNumbers)

	assert.Equalf(t, len(*res), expectedAmount, "Tamanho experado do array: %v, tamanho obtido: %v", expectedAmount, len(*res))
}

func TestGetExceptRegistrationNumber(t *testing.T) {
	employeeIds := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	employeeViewIds := []int64{1, 2, 3, 4, 5, 6, 7}
	expectedAmount := len(employeeIds) - len(employeeViewIds)

	res := employeeModel.GetExceptRegistrationNumber(employeeIds, employeeViewIds)

	assert.Equalf(t, len(res), expectedAmount, "Tamanho experado do array: %v, tamanho obtido: %v", expectedAmount, len(res))
}
