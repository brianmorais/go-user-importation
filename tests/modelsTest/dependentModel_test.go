package modelsTest

import (
	"database/sql"
	"testing"

	"github.com/brianmorais/go-user-importation/domain/enums/kinshipType"
	dependentModel "github.com/brianmorais/go-user-importation/domain/models/dependent"
	"github.com/brianmorais/go-user-importation/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMustEquals(t *testing.T) {
	dependent := mocks.GetDependent()
	dependentView := mocks.GetDependentView()

	res := dependent.Equals(dependentView)

	assert.Truef(t, res, "Dependente não é igual ao DependenteView")
}

func TestShouldNotEquals(t *testing.T) {
	dependent := mocks.GetDependent()
	dependentView := mocks.GetDependentView()
	dependentView.Cpf = sql.NullString{String: "987456321"}

	res := dependent.Equals(dependentView)

	assert.Falsef(t, res, "Dependente é igual ao DependenteView e deveria ser diferente")
}

func TestMustGetViewDataKindship(t *testing.T) {
	dependent := mocks.GetDependentView()
	expectedKinship := kinshipType.Child

	res := dependent.GetViewDataKindship()

	assert.Equalf(t, res, expectedKinship, "Valor esperado: %v, valor obtido: %v", expectedKinship, res)
}

func TestShouldNotGetViewDataKindship(t *testing.T) {
	dependent := mocks.GetDependentView()
	expectedKinship := kinshipType.Spouse

	res := dependent.GetViewDataKindship()

	assert.NotEqualf(t, res, expectedKinship, "Valor esperado: %v, valor obtido: %v", expectedKinship, res)
}

func TestMustNameIsValid(t *testing.T) {
	dependent := mocks.GetDependentView()
	expectedValue := true

	res := dependent.NameIsValid()

	assert.Equalf(t, res, expectedValue, "Valor esperado: %v, valor obtido: %v", expectedValue, res)
}

func TestShouldNotNameIsValid(t *testing.T) {
	dependent := mocks.GetDependentView()
	dependent.PrimeirioNome.Valid = false
	expectedValue := false

	res := dependent.NameIsValid()

	assert.Equalf(t, res, expectedValue, "Valor esperado: %v, valor obtido: %v", expectedValue, res)
}

func TestConvertDependentRegistrationNumberToString(t *testing.T) {
	dependent := mocks.GetDependentView()
	expectedValue := "1"

	res := dependent.RegistrationNumberToString()

	assert.Equalf(t, res, expectedValue, "Valor esperado: %v, valor obtido: %v", expectedValue, res)
}

func TestMustFindDependentByEmployeeIdAndName(t *testing.T) {
	employeeId := 2
	dependents := mocks.GetDependentsMock()
	fullName := "MARCOS RUBENS"

	res := dependents.FindDependentByEmployeeIdAndName(int64(employeeId), fullName)

	assert.NotEqualf(t, res, dependentModel.Dependent{}, "Deveria ter encontrado dependentes")
}

func TestShouldNotFindDependentByEmployeeIdAndName(t *testing.T) {
	employeeId := 30
	dependents := mocks.GetDependentsMock()
	fullName := "JOAO CLAUDIO"

	res := dependents.FindDependentByEmployeeIdAndName(int64(employeeId), fullName)

	assert.Equalf(t, res, dependentModel.Dependent{}, "Não deveria ter encontrado dependentes")
}

func TestMustFindDependentByRegistrationNumberAndName(t *testing.T) {
	registrationNumber := 3
	fullName := "MARIA CLARA"
	dependents := mocks.GetDependentsViewMock()

	res := dependents.FindDependentByRegistrationNumberAndName(int64(registrationNumber), fullName)

	assert.NotEqualf(t, res, dependentModel.DependentView{}, "Deveria ter encontrado dependentes")
}

func TestShouldNotFindDependentByRegistrationNumberAndName(t *testing.T) {
	registrationNumber := 4
	fullName := "JOAO CLAUDIO"
	dependents := mocks.GetDependentsViewMock()

	res := dependents.FindDependentByRegistrationNumberAndName(int64(registrationNumber), fullName)

	assert.Equalf(t, res, dependentModel.DependentView{}, "Não deveria ter encontrado dependentes")
}
