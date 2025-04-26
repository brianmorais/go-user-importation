package dependentModel

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/brianmorais/go-user-importation/domain/utils"
)

type Dependent struct {
	DependentId  sql.NullInt64
	FirstName    sql.NullString
	LastName     sql.NullString
	DateOfBirth  sql.NullTime
	Gender       sql.NullString
	Active       sql.NullBool
	CreatedDate  sql.NullTime
	ModifiedDate sql.NullTime
	ModifiedUser sql.NullString
	EmployeeId   sql.NullInt64
	Kinship      sql.NullInt32
	Cpf          sql.NullString
}

type Dependents []Dependent

func (dependent Dependent) Equals(dependentView DependentView) bool {
	firstNameAreEquals := utils.CleanString(dependent.GetFirstName()) == utils.CleanString(dependentView.GetPrimeiroNome())
	lastNameAreEquals := utils.CleanString(dependent.GetLastName()) == utils.CleanString(dependentView.GetSobrenome())
	dateOfBirthAreEquals := dependent.GetDateOfBirth() == dependentView.GetDataNascimento()
	genderAreEquals := dependent.GetGender() == dependentView.GetSexo()
	cpfAreEquals := dependent.GetCpf() == dependentView.GetCpf()
	activeAreEquals := dependent.GetActive() == (dependentView.GetCdiSituacaoDependente() == 1)

	if firstNameAreEquals && lastNameAreEquals && dateOfBirthAreEquals && genderAreEquals && cpfAreEquals && activeAreEquals {
		return true
	}

	return false
}

func (dependents *Dependents) FindDependentByEmployeeIdAndName(employeeId int64, fullName string) Dependent {
	for i := range *dependents {
		name := fmt.Sprintf("%v %v", utils.CleanString((*dependents)[i].GetFirstName()), utils.CleanString((*dependents)[i].GetLastName()))
		if (*dependents)[i].EmployeeId.Int64 == employeeId && name == fullName {
			return (*dependents)[i]
		}
	}

	return Dependent{}
}

func (dependent *Dependent) SetFirstName(firstName string) {
	dependent.FirstName = sql.NullString{String: firstName, Valid: true}
}

func (dependent *Dependent) SetLastName(lastName string) {
	dependent.LastName = sql.NullString{String: lastName, Valid: true}
}

func (dependent *Dependent) SetDateOfBirth(dateOfBirth time.Time) {
	dependent.DateOfBirth = sql.NullTime{Time: dateOfBirth, Valid: true}
}

func (dependent *Dependent) SetGender(gender string) {
	dependent.Gender = sql.NullString{String: gender, Valid: true}
}

func (dependent *Dependent) SetActive(active bool) {
	dependent.Active = sql.NullBool{Bool: active, Valid: true}
}

func (dependent *Dependent) SetCreatedDate(createdDate time.Time) {
	dependent.CreatedDate = sql.NullTime{Time: createdDate, Valid: true}
}

func (dependent *Dependent) SetModifiedDate(modifiedDate time.Time) {
	dependent.ModifiedDate = sql.NullTime{Time: modifiedDate, Valid: true}
}

func (dependent *Dependent) SetModifiedUser(modifiedUser string) {
	dependent.ModifiedUser = sql.NullString{String: modifiedUser, Valid: true}
}

func (dependent *Dependent) SetEmployeeId(employeeId int64) {
	dependent.EmployeeId = sql.NullInt64{Int64: employeeId, Valid: true}
}

func (dependent *Dependent) SetKinship(kinshipType int32) {
	dependent.Kinship = sql.NullInt32{Int32: kinshipType, Valid: true}
}

func (dependent *Dependent) SetCpf(cpf string) {
	dependent.Cpf = sql.NullString{String: cpf, Valid: true}
}

func (dependent *Dependent) GetEmployeeId() int64 {
	return dependent.EmployeeId.Int64
}

func (dependent *Dependent) GetFirstName() string {
	return dependent.FirstName.String
}

func (dependent *Dependent) GetLastName() string {
	return dependent.LastName.String
}

func (dependent *Dependent) GetKinship() int32 {
	return dependent.Kinship.Int32
}

func (dependent *Dependent) GetDateOfBirth() time.Time {
	return dependent.DateOfBirth.Time
}

func (dependent *Dependent) GetGender() string {
	return dependent.Gender.String
}

func (dependent *Dependent) GetCpf() string {
	return dependent.Cpf.String
}

func (dependent *Dependent) GetActive() bool {
	return dependent.Active.Bool
}
