package interfaces

import benefitModel "github.com/brianmorais/go-user-importation/domain/models/benefit"

type IBenefitRepository interface {
	GetBenefitByEmployeIdAndCycleId(employeeId int64, cycleId int64) (benefitModel.EmployeeBenefit, error)
	GetActiveCycle() (benefitModel.Cycle, error)
	CreateEmployeeBenefit(employeeBenefit benefitModel.EmployeeBenefit) (int64, error)
	UpdateEmployeeBenefit(employeeBenefit benefitModel.EmployeeBenefit) (int64, error)
}
