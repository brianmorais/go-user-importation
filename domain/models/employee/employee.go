package employee

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/brianmorais/go-user-importation/domain/utils"
)

type Employee struct {
	EmployeeId         sql.NullInt64
	FirstName          sql.NullString
	LastName           sql.NullString
	DateOfBirth        sql.NullTime
	Gender             sql.NullString
	RegistrationNumber sql.NullString
	Email              sql.NullString
	AdmissionDate      sql.NullTime
	Active             sql.NullBool
	Location           sql.NullString
	CreatedDate        sql.NullTime
	ModifiedDate       sql.NullTime
	ModifiedUser       sql.NullString
	RoleId             sql.NullInt32
	EmployeeTypeId     sql.NullInt32
	AbsenceId          sql.NullInt32
	TerminationDate    sql.NullTime
	ExceptionReason    sql.NullString
	ExceptionBeginDate sql.NullTime
	ExceptionEndDate   sql.NullTime
	Cpf                sql.NullString
}

type Employees []Employee

func GetExceptRegistrationNumber(employees []string, employeesView []int64) []string {
	var registrationNumbers []string
	var contains bool

	for i := range employees {
		contains = false
		for j := range employeesView {
			if employees[i] == strconv.FormatInt(employeesView[j], 10) {
				contains = true
				break
			}
		}

		if !contains {
			registrationNumbers = append(registrationNumbers, employees[i])
		}
	}

	return registrationNumbers
}

func (employee Employee) Equals(employeeView EmployeeView) bool {
	genderAreEquals := utils.CleanString(employeeView.GetSexoSigla()) == utils.CleanString(employee.GetGender())
	activeAreEquals := employee.GetActive() == (employeeView.GetSituacao() == "A" && (!employeeView.DataRescisao.Valid || employeeView.DataRescisao == sql.NullTime{}))
	firstNameAreEquals := utils.CleanString(employee.GetFirstName()) == utils.CleanString(employeeView.GetPrimeiroNome())
	lastNameAreEquals := utils.CleanString(employee.GetLastName()) == utils.CleanString(employeeView.GetSobrenome())
	dateOfBirthAreEquals := employee.GetDateOfBirth() == employeeView.GetDataNascimento()
	admissionDateAreEquals := employee.GetAdmissionDate() == employeeView.GetDataAdmissao()
	locationAreEquals := employee.GetLocation() == employeeView.GetEstado()
	cpfAreEquals := employee.GetCpf() == employeeView.GetCpf()

	if genderAreEquals && activeAreEquals && firstNameAreEquals && lastNameAreEquals && dateOfBirthAreEquals && admissionDateAreEquals && locationAreEquals && cpfAreEquals {
		return true
	}

	return false
}

func (employees *Employees) FindEmployeeById(id int64) Employee {
	for i := range *employees {
		if (*employees)[i].GetEmployeeId() == id {
			return (*employees)[i]
		}
	}

	return Employee{}
}

func (employees *Employees) FindEmployeeByRegistrationNumber(registrationNumber string) Employee {
	for i := range *employees {
		if (*employees)[i].RegistrationNumber.String == registrationNumber {
			return (*employees)[i]
		}
	}

	return Employee{}
}

func (employees *Employees) GetNotExceptionAndActiveEmployeeRegistrationNumbers() []string {
	var registrationNumbers []string

	for i := range *employees {
		if (*employees)[i].GetRoleId() != 7 && (*employees)[i].GetActive() {
			registrationNumbers = append(registrationNumbers, (*employees)[i].GetRegistrationNumber())
		}
	}

	return registrationNumbers
}

func (employees *Employees) GetEmployeesToDeactivate(registrationNumbers []string) *Employees {
	var toDeactivate Employees

	for i := range *employees {
		for j := range registrationNumbers {
			if (*employees)[i].GetRegistrationNumber() == registrationNumbers[j] {
				toDeactivate = append(toDeactivate, (*employees)[i])
				break
			}
		}
	}

	return &toDeactivate
}

func (employee *Employee) IsValid() bool {
	return employee.EmployeeId.Valid && (employee != &Employee{})
}

func (employee Employee) RegistrationNumberToInt64() int64 {
	re, _ := strconv.ParseInt(employee.GetRegistrationNumber(), 10, 64)
	return re
}

func (employee Employee) IsTemporaryEmployee() bool {
	return employee.GetRoleId() == 7
}

func (employee *Employee) HasExceptionDate() bool {
	return employee.ExceptionEndDate.Valid
}

func (employee *Employee) BlockUserException() bool {
	return employee.HasExceptionDate() && (time.Now().After(employee.GetExceptionEndDate()) || employee.GetActive() || (employee.GetActive() && employee.GetExceptionEndDate() == time.Time{}))
}

func (employee *Employee) IsUnlimitedDateException() bool {
	return employee.GetExceptionEndDate() == time.Time{}
}

func (employee *Employee) SetFirstName(firstName string) {
	employee.FirstName = sql.NullString{String: firstName, Valid: true}
}

func (employee *Employee) SetLastName(lastName string) {
	employee.LastName = sql.NullString{String: lastName, Valid: true}
}

func (employee *Employee) SetDateOfBirth(dateOfBirth time.Time) {
	employee.DateOfBirth = sql.NullTime{Time: dateOfBirth, Valid: true}
}

func (employee *Employee) SetGender(gender string) {
	employee.Gender = sql.NullString{String: gender, Valid: true}
}

func (employee *Employee) SetRegistrationNumber(registrationNumber string) {
	employee.RegistrationNumber = sql.NullString{String: registrationNumber, Valid: true}
}

func (employee *Employee) SetEmail(email string) {
	employee.Email = sql.NullString{String: email, Valid: true}
}

func (employee *Employee) SetAdmissionDate(admissionDate time.Time) {
	employee.AdmissionDate = sql.NullTime{Time: admissionDate, Valid: true}
}

func (employee *Employee) SetActive(active bool) {
	employee.Active = sql.NullBool{Bool: active, Valid: true}
}

func (employee *Employee) SetLocation(location string) {
	employee.Location = sql.NullString{String: location, Valid: true}
}

func (employee *Employee) SetCreatedDate(createdDate time.Time) {
	employee.CreatedDate = sql.NullTime{Time: createdDate, Valid: true}
}

func (employee *Employee) SetModifiedDate(modifiedDate time.Time) {
	employee.ModifiedDate = sql.NullTime{Time: modifiedDate, Valid: true}
}

func (employee *Employee) SetModifiedUser(modifiedUser string) {
	employee.ModifiedUser = sql.NullString{String: modifiedUser, Valid: true}
}

func (employee *Employee) SetRoleId(roleId int32) {
	employee.RoleId = sql.NullInt32{Int32: roleId, Valid: true}
}

func (employee *Employee) SetEmployeeTypeId(employeeTypeId int32) {
	employee.EmployeeTypeId = sql.NullInt32{Int32: employeeTypeId, Valid: true}
}

func (employee *Employee) SetAbsenceId(absenceId int32) {
	employee.AbsenceId = sql.NullInt32{Int32: absenceId, Valid: true}
}

func (employee *Employee) SetTerminationDate(terminationDate time.Time) {
	employee.TerminationDate = sql.NullTime{Time: terminationDate, Valid: true}
}

func (employee *Employee) SetExceptionReason(exceptionReason string) {
	employee.ExceptionReason = sql.NullString{String: exceptionReason, Valid: true}
}

func (employee *Employee) SetExceptionBenginDate(exceptionBeginDate time.Time) {
	employee.ExceptionBeginDate = sql.NullTime{Time: exceptionBeginDate, Valid: true}
}

func (employee *Employee) SetExceptionEndDate(exceptionEndDate time.Time) {
	employee.ExceptionEndDate = sql.NullTime{Time: exceptionEndDate, Valid: true}
}

func (employee *Employee) SetCpf(cpf string) {
	employee.Cpf = sql.NullString{String: cpf, Valid: true}
}

func (employee *Employee) GetRegistrationNumber() string {
	return employee.RegistrationNumber.String
}

func (employee *Employee) GetEmployeeId() int64 {
	return employee.EmployeeId.Int64
}

func (employee *Employee) GetFirstName() string {
	return employee.FirstName.String
}

func (employee *Employee) GetLastName() string {
	return employee.LastName.String
}

func (employee *Employee) GetEmail() string {
	return employee.Email.String
}

func (employee *Employee) GetRoleId() int32 {
	return employee.RoleId.Int32
}

func (employee *Employee) GetExceptionEndDate() time.Time {
	return employee.ExceptionEndDate.Time
}

func (employee *Employee) GetActive() bool {
	return employee.Active.Bool
}

func (employee *Employee) GetDateOfBirth() time.Time {
	return employee.DateOfBirth.Time
}

func (employee *Employee) GetAdmissionDate() time.Time {
	return employee.AdmissionDate.Time
}

func (employee *Employee) GetLocation() string {
	return employee.Location.String
}

func (employee *Employee) GetCpf() string {
	return employee.Cpf.String
}

func (employee *Employee) GetGender() string {
	return employee.Gender.String
}
