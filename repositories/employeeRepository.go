package repositories

import (
	employeeModel "github.com/brianmorais/go-user-importation/domain/models/employee"
)

type EmployeeRepository struct{}

func (EmployeeRepository) GetEmployees() (*employeeModel.Employees, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			EmployeeId,
			FirstName,
			LastName,
			DateOfBirth,
			Gender,
			RegistrationNumber,
			Email,
			AdmissionDate,
			Active,
			[Location],
			CreatedDate,
			ModifiedDate,
			ModifiedUser,
			RoleId,
			EmployeeTypeId,
			AbsenceId,
			TerminationDate,
			ExceptionReason,
			ExceptionBeginDate,
			ExceptionEndDate,
			Cpf
		FROM 
			Employee`

	rows, err := conn.Query(queryText)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees employeeModel.Employees

	for rows.Next() {
		var employee employeeModel.Employee

		err := rows.Scan(
			&employee.EmployeeId,
			&employee.FirstName,
			&employee.LastName,
			&employee.DateOfBirth,
			&employee.Gender,
			&employee.RegistrationNumber,
			&employee.Email,
			&employee.AdmissionDate,
			&employee.Active,
			&employee.Location,
			&employee.CreatedDate,
			&employee.ModifiedDate,
			&employee.ModifiedUser,
			&employee.RoleId,
			&employee.EmployeeTypeId,
			&employee.AbsenceId,
			&employee.TerminationDate,
			&employee.ExceptionReason,
			&employee.ExceptionBeginDate,
			&employee.ExceptionEndDate,
			&employee.Cpf,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

func (EmployeeRepository) GetEmployeesView() (*employeeModel.EmployeesView, error) {
	conn := getReadConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			RE,
			Primeiro_nome,
			Sobrenome,
			NOME_COMPLETO,
			DT_NASCIMENTO,
			DT_ADMISSAO,
			ID_situacao,
			Situacao,
			DT_RESCISAO,
			ENDERECO,
			NUMERO,
			CIDADE,
			ESTADO,
			CEP,
			PAIS,
			USUARIO,
			GPO_FUNCIONAL,
			sex_d1ssexosigla,
			CPF
		FROM 
			Vw_vemployee`

	rows, err := conn.Query(queryText)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees employeeModel.EmployeesView

	for rows.Next() {
		var employee employeeModel.EmployeeView

		err := rows.Scan(
			&employee.Re,
			&employee.PrimeiroNome,
			&employee.Sobrenome,
			&employee.NomeCompleto,
			&employee.DataNascimento,
			&employee.DataAdmissao,
			&employee.IdSituacao,
			&employee.Situacao,
			&employee.DataRescisao,
			&employee.Endereco,
			&employee.Numero,
			&employee.Cidade,
			&employee.Estado,
			&employee.Cep,
			&employee.Pais,
			&employee.Usuario,
			&employee.GpoFuncional,
			&employee.SexoSigla,
			&employee.Cpf,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

func (EmployeeRepository) UpdateEmployee(employee *employeeModel.Employee) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		UPDATE [Employee] SET 
			FirstName = ?,
			LastName = ?,
			DateOfBirth = ?,
			Gender = ?,
			RegistrationNumber = ?,
			Email = ?,
			AdmissionDate = ?,
			Active = ?,
			[Location] = ?,
			ModifiedDate = ?,
			ModifiedUser = ?,
			RoleId = ?,
			EmployeeTypeId = ?,
			AbsenceId = ?,
			TerminationDate = ?,
			ExceptionReason = ?,
			ExceptionBeginDate = ?,
			ExceptionEndDate = ?,
			Cpf = ?
		WHERE 
			EmployeeId = ?`

	res, err := conn.Exec(
		queryText,
		employee.FirstName,
		employee.LastName,
		employee.DateOfBirth,
		employee.Gender,
		employee.RegistrationNumber,
		employee.Email,
		employee.AdmissionDate,
		employee.Active,
		employee.Location,
		employee.ModifiedDate,
		employee.ModifiedUser,
		employee.RoleId,
		employee.EmployeeTypeId,
		employee.AbsenceId,
		employee.TerminationDate,
		employee.ExceptionReason,
		employee.ExceptionBeginDate,
		employee.ExceptionEndDate,
		employee.Cpf,
		employee.EmployeeId,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (EmployeeRepository) CreateEmployee(employee *employeeModel.Employee) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		INSERT INTO Employee (
			FirstName,
			LastName,
			DateOfBirth,
			Gender,
			RegistrationNumber,
			Email,
			AdmissionDate,
			Active,
			[Location],
			CreatedDate,
			ModifiedDate,
			ModifiedUser,
			RoleId,
			EmployeeTypeId,
			AbsenceId,
			TerminationDate,
			ExceptionReason,
			ExceptionBeginDate,
			ExceptionEndDate,
			Cpf
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := conn.Exec(
		queryText,
		employee.FirstName,
		employee.LastName,
		employee.DateOfBirth,
		employee.Gender,
		employee.RegistrationNumber,
		employee.Email,
		employee.AdmissionDate,
		employee.Active,
		employee.Location,
		employee.CreatedDate,
		employee.ModifiedDate,
		employee.ModifiedUser,
		employee.RoleId,
		employee.EmployeeTypeId,
		employee.AbsenceId,
		employee.TerminationDate,
		employee.ExceptionReason,
		employee.ExceptionBeginDate,
		employee.ExceptionEndDate,
		employee.Cpf,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (EmployeeRepository) GetEmployeeRoles() (*employeeModel.EmployeeRoles, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			RoleId,
			Role,
			Hierarchy
		FROM
			EmployeeRole`

	rows, err := conn.Query(queryText)

	if err != nil {
		return &employeeModel.EmployeeRoles{}, err
	}

	defer rows.Close()

	var roles employeeModel.EmployeeRoles

	for rows.Next() {
		var role employeeModel.EmployeeRole

		err := rows.Scan(
			&role.RoleId,
			&role.Role,
			&role.Hierarchy,
		)

		if err != nil {
			return &employeeModel.EmployeeRoles{}, err
		}

		roles = append(roles, role)
	}

	return &roles, nil
}

func (EmployeeRepository) GetEmployeeByRegistrationNumber(registrationNumber string) (*employeeModel.Employee, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			EmployeeId,
			FirstName,
			LastName,
			DateOfBirth,
			Gender,
			RegistrationNumber,
			Email,
			AdmissionDate,
			Active,
			[Location],
			CreatedDate,
			ModifiedDate,
			ModifiedUser,
			RoleId,
			EmployeeTypeId,
			AbsenceId,
			TerminationDate,
			ExceptionReason,
			ExceptionBeginDate,
			ExceptionEndDate,
			Cpf
		FROM 
			Employee
		WHERE
			RegistrationNumber = ?`

	rows, err := conn.Query(queryText, registrationNumber)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employee employeeModel.Employee

	for rows.Next() {
		err := rows.Scan(
			&employee.EmployeeId,
			&employee.FirstName,
			&employee.LastName,
			&employee.DateOfBirth,
			&employee.Gender,
			&employee.RegistrationNumber,
			&employee.Email,
			&employee.AdmissionDate,
			&employee.Active,
			&employee.Location,
			&employee.CreatedDate,
			&employee.ModifiedDate,
			&employee.ModifiedUser,
			&employee.RoleId,
			&employee.EmployeeTypeId,
			&employee.AbsenceId,
			&employee.TerminationDate,
			&employee.ExceptionReason,
			&employee.ExceptionBeginDate,
			&employee.ExceptionEndDate,
			&employee.Cpf,
		)

		if err != nil {
			return nil, err
		}
	}

	return &employee, nil
}
