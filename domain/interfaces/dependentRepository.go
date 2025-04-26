package interfaces

import dependentModel "github.com/brianmorais/go-user-importation/domain/models/dependent"

type IDependentRepository interface {
	GetDependentsView() (*dependentModel.DependentsView, error)
	GetDependents() (*dependentModel.Dependents, error)
	UpdateDependent(dependent *dependentModel.Dependent) (int64, error)
	CreateDependent(dependent *dependentModel.Dependent) (int64, error)
}
