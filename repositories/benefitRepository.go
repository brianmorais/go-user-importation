package repositories

import (
	benefitModel "github.com/brianmorais/go-user-importation/domain/models/benefit"
)

type BenefitRepository struct{}

func (BenefitRepository) GetBenefitByEmployeIdAndCycleId(employeeId int64, cycleId int64) (benefitModel.EmployeeBenefit, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			EmployeeId,
			CycleId,
			RoleId,
			Benefit,
			CreatedDate,
			ModifiedDate,
			ModifiedUser
		FROM 
			EmployeeBenefit
		WHERE
			EmployeeId = ?
			AND CycleId = ?`

	rows, err := conn.Query(queryText, employeeId, cycleId)

	if err != nil {
		return benefitModel.EmployeeBenefit{}, err
	}

	defer rows.Close()

	var benefit benefitModel.EmployeeBenefit

	for rows.Next() {
		err = rows.Scan(
			&benefit.EmployeeId,
			&benefit.CycleId,
			&benefit.RoleId,
			&benefit.Benefit,
			&benefit.CreatedDate,
			&benefit.ModifiedDate,
			&benefit.ModifiedUser,
		)
	}

	if err != nil {
		return benefitModel.EmployeeBenefit{}, err
	}

	return benefit, nil
}

func (BenefitRepository) GetActiveCycle() (benefitModel.Cycle, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		SELECT 
			CycleId,
			StartPeriod,
			Active,
			CreatedDate,
			ModifiedDate,
			ModifiedUser,
			BenefitSetting,
			Title,
			Visible
		FROM 
			Cycle
		WHERE
			Active = 1`

	rows, err := conn.Query(queryText)

	if err != nil {
		return benefitModel.Cycle{}, err
	}

	defer rows.Close()

	var cycle benefitModel.Cycle

	for rows.Next() {
		err = rows.Scan(
			&cycle.CycleId,
			&cycle.StartPeriod,
			&cycle.Active,
			&cycle.CreatedDate,
			&cycle.ModifiedDate,
			&cycle.ModifiedUser,
			&cycle.BenefitSetting,
			&cycle.Title,
			&cycle.Visible,
		)
	}

	if err != nil {
		return benefitModel.Cycle{}, err
	}

	return cycle, nil
}

func (BenefitRepository) CreateEmployeeBenefit(employeeBenefit benefitModel.EmployeeBenefit) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		INSERT INTO EmployeeBenefit (
			EmployeeId,
			CycleId,
			RoleId,
			Benefit,
			CreatedDate,
			ModifiedDate,
			ModifiedUser
		) VALUES (?, ?, ?, ?, ?, ?, ?)`

	res, err := conn.Exec(
		queryText,
		employeeBenefit.EmployeeId,
		employeeBenefit.CycleId,
		employeeBenefit.RoleId,
		employeeBenefit.Benefit,
		employeeBenefit.CreatedDate,
		employeeBenefit.ModifiedDate,
		employeeBenefit.ModifiedUser,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (BenefitRepository) UpdateEmployeeBenefit(employeeBenefit benefitModel.EmployeeBenefit) (int64, error) {
	conn := getWriteConnection()

	defer conn.Close()

	queryText := `
		UPDATE EmployeeBenefit SET
			RoleId = ?,
			Benefit = ?,
			ModifiedDate = ?,
			ModifiedUser = ?
		WHERE 
			EmployeeId = ?
			AND CycleId = ?`

	res, err := conn.Exec(
		queryText,
		employeeBenefit.RoleId,
		employeeBenefit.Benefit,
		employeeBenefit.ModifiedDate,
		employeeBenefit.ModifiedUser,
		employeeBenefit.EmployeeId,
		employeeBenefit.CycleId,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
