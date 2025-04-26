package ioc

import (
	"github.com/golobby/container/v3"
	"github.com/brianmorais/go-user-importation/configuration"
	"github.com/brianmorais/go-user-importation/domain/interfaces"
	"github.com/brianmorais/go-user-importation/repositories"
)

func SetDependencies() {
	container.Singleton(func() *configuration.Settings {
		config := configuration.GetSettings()
		return config
	})

	container.Transient(func() interfaces.IBenefitRepository {
		return &repositories.BenefitRepository{}
	})

	container.Transient(func() interfaces.IDependentRepository {
		return &repositories.DependentRepository{}
	})

	container.Transient(func() interfaces.IEmployeeRepository {
		return &repositories.EmployeeRepository{}
	})

	container.Transient(func() interfaces.ILogRepository {
		return &repositories.LogRepository{}
	})

	container.Transient(func() interfaces.IUserRepository {
		return &repositories.UserRepository{}
	})
}
