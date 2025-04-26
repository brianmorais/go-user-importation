package repositories

import (
	dependentModel "github.com/brianmorais/go-user-importation/domain/models/dependent"
)

type DependentRepository struct{}

func (DependentRepository) GetDependentsView() (*dependentModel.DependentsView, error) {
	conn := getReadConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			ID_DEPENDENTE,
			RE,
			Primeiro_nome,
			Sobrenome,
			DEP_DssNomeCompleto,
			DEP_CdiLigacaoPessoa,
			DT_NASCIMENTO,
			SEXO,
			DEP_CdiSituacaoDependente,
			CPF
		FROM
			Vw_vdependentes`

	rows, err := conn.Query(queryText)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var dependents dependentModel.DependentsView

	for rows.Next() {
		var dependent dependentModel.DependentView

		err := rows.Scan(
			&dependent.IdDependent,
			&dependent.Re,
			&dependent.PrimeirioNome,
			&dependent.Sobrenome,
			&dependent.NomeCompleto,
			&dependent.CdiLigacaoPessoa,
			&dependent.DataNascimento,
			&dependent.Sexo,
			&dependent.CdiSituacaoDependente,
			&dependent.Cpf,
		)

		if err != nil {
			return nil, err
		}

		dependents = append(dependents, dependent)
	}

	return &dependents, nil
}

func (DependentRepository) GetDependents() (*dependentModel.Dependents, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT
			DependentId,
			FirstName,
			LastName,
			DateOfBirth,
			Gender,
			Active,
			CreatedDate,
			ModifiedDate,
			ModifiedUser,
			EmployeeId,
			Kinship,
			Cpf
		FROM 
			[Dependent]`

	rows, err := conn.Query(queryText)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var dependents dependentModel.Dependents

	for rows.Next() {
		var dependent dependentModel.Dependent

		err := rows.Scan(
			&dependent.DependentId,
			&dependent.FirstName,
			&dependent.LastName,
			&dependent.DateOfBirth,
			&dependent.Gender,
			&dependent.Active,
			&dependent.CreatedDate,
			&dependent.ModifiedDate,
			&dependent.ModifiedUser,
			&dependent.EmployeeId,
			&dependent.Kinship,
			&dependent.Cpf,
		)

		if err != nil {
			return nil, err
		}

		dependents = append(dependents, dependent)
	}

	return &dependents, nil
}

func (DependentRepository) UpdateDependent(dependent *dependentModel.Dependent) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		UPDATE [Dependent] SET
			FirstName = ?,
			LastName = ?,
			DateOfBirth = ?,
			Gender = ?,
			Active = ?,
			ModifiedDate = ?,
			ModifiedUser = ?,
			EmployeeId = ?,
			Kinship = ?,
			Cpf = ?
		WHERE DependentId = ?`

	res, err := conn.Exec(
		queryText,
		dependent.FirstName,
		dependent.LastName,
		dependent.DateOfBirth,
		dependent.Gender,
		dependent.Active,
		dependent.ModifiedDate,
		dependent.ModifiedUser,
		dependent.EmployeeId,
		dependent.Kinship,
		dependent.Cpf,
		dependent.DependentId,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (DependentRepository) CreateDependent(dependent *dependentModel.Dependent) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		INSERT INTO [Dependent] (
			FirstName,
			LastName,
			DateOfBirth,
			Gender,
			Kinship,
			EmployeeId,
			Active,
			CreatedDate,
			ModifiedDate,
			ModifiedUser,
			Cpf
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := conn.Exec(
		queryText,
		dependent.FirstName,
		dependent.LastName,
		dependent.DateOfBirth,
		dependent.Gender,
		dependent.Kinship,
		dependent.EmployeeId,
		dependent.Active,
		dependent.CreatedDate,
		dependent.ModifiedDate,
		dependent.ModifiedUser,
		dependent.Cpf,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
