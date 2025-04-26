package employeeModel

import (
	"database/sql"
	"strconv"
	"time"
)

type EmployeeView struct {
	Re             sql.NullInt64
	PrimeiroNome   sql.NullString
	Sobrenome      sql.NullString
	NomeCompleto   sql.NullString
	DataNascimento sql.NullTime
	DataAdmissao   sql.NullTime
	IdSituacao     sql.NullString
	Situacao       sql.NullString
	DataRescisao   sql.NullTime
	Endereco       sql.NullString
	Numero         sql.NullString
	Cidade         sql.NullString
	Estado         sql.NullString
	Cep            sql.NullString
	Pais           sql.NullString
	Usuario        sql.NullString
	GpoFuncional   sql.NullString
	SexoSigla      sql.NullString
	Cpf            sql.NullString
}

type EmployeesView []EmployeeView

func (employees *EmployeesView) GetEmployeeViewRegistrationNumbers() []int64 {
	var registrationNumbers []int64

	for i := range *employees {
		registrationNumbers = append(registrationNumbers, (*employees)[i].Re.Int64)
	}

	return registrationNumbers
}

func (employee EmployeeView) RegistrationNumberToString() string {
	return strconv.FormatInt(employee.Re.Int64, 10)
}

func (employee *EmployeeView) SetPrimeiroNome(nome string) {
	employee.PrimeiroNome = sql.NullString{String: nome, Valid: true}
}

func (employee *EmployeeView) SetSobrenome(nome string) {
	employee.Sobrenome = sql.NullString{String: nome, Valid: true}
}

func (employee *EmployeeView) GetPrimeiroNome() string {
	return employee.PrimeiroNome.String
}

func (employee *EmployeeView) GetSobrenome() string {
	return employee.Sobrenome.String
}

func (employee *EmployeeView) GetSituacao() string {
	return employee.Situacao.String
}

func (employee *EmployeeView) GetSexoSigla() string {
	return employee.SexoSigla.String
}

func (employee *EmployeeView) GetUsuario() string {
	return employee.Usuario.String
}

func (employee *EmployeeView) GetGpoFuncional() string {
	return employee.GpoFuncional.String
}

func (employee *EmployeeView) GetEstado() string {
	return employee.Estado.String
}

func (employee *EmployeeView) GetDataRescisao() time.Time {
	return employee.DataRescisao.Time
}

func (employee *EmployeeView) GetDataNascimento() time.Time {
	return employee.DataNascimento.Time
}

func (employee *EmployeeView) GetDataAdmissao() time.Time {
	return employee.DataAdmissao.Time
}

func (employee *EmployeeView) GetCpf() string {
	return employee.Cpf.String
}
