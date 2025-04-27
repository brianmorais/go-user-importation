package application

import (
	"fmt"
	"sync"
	"time"

	dependentModel "github.com/brianmorais/go-user-importation/domain/models/dependent"
	employeeModel "github.com/brianmorais/go-user-importation/domain/models/employee"
	"github.com/brianmorais/go-user-importation/domain/utils"
)

func readDependentViewData() {
	saveMessageLog(fmt.Sprintf("Início da importação de dependentes - %v", *importationNumber))

	var err error
	var dependentsView *dependentModel.DependentsView
	var dependents *dependentModel.Dependents
	var employees *employeeModel.Employees

	if dependentsView, err = dependentRepository.GetDependentsView(); err != nil {
		saveErrorLog("Erro ao obter dependentes da view", err, "")
		return
	}

	if dependents, err = dependentRepository.GetDependents(); err != nil {
		saveErrorLog("Erro ao obter dependentes do banco", err, "")
		return
	}

	if employees, err = employeeRepository.GetEmployees(); err != nil {
		saveErrorLog("Erro ao obter employees do banco na importação de dependentes", err, "")
		return
	}

	deactivateDependents(dependents, employees, dependentsView)
	updateDependents(dependents, employees, dependentsView)
	saveMessageLog(fmt.Sprintf("Fim da importação de dependentes - %v", *importationNumber))
}

func deactivateDependents(dependents *dependentModel.Dependents, employees *employeeModel.Employees, dependentsView *dependentModel.DependentsView) {
	if settings.GoRoutines.UseGoRoutines {
		var wg sync.WaitGroup
		dependentsChannel := make(chan dependentModel.Dependent)

		wg.Add(settings.GoRoutines.GoRoutinesCount)

		for i := 1; i <= settings.GoRoutines.GoRoutinesCount; i++ {
			go func() {
				defer wg.Done()
				for currentDependent := range dependentsChannel {
					executeDependentDeactivateProcess(currentDependent, employees, dependentsView)
				}
			}()
		}

		for j := range *dependents {
			dependentsChannel <- (*dependents)[j]
		}

		close(dependentsChannel)
		wg.Wait()
	} else {
		for i := range *dependents {
			currentDependent := (*dependents)[i]
			executeDependentDeactivateProcess(currentDependent, employees, dependentsView)
		}
	}
}

func executeDependentDeactivateProcess(currentDependent dependentModel.Dependent, employees *employeeModel.Employees, dependentsView *dependentModel.DependentsView) {
	employee := employees.FindEmployeeById(currentDependent.GetEmployeeId())

	if (employee != employeeModel.Employee{}) {
		userIsDatabase, err := isUserDatabaseType(employee)

		if err != nil {
			saveErrorLog(fmt.Sprintf("Erro ao tentar desativar os dados do dependente %v", currentDependent.GetFirstName()), err, employee.GetRegistrationNumber())
			return
		}

		if !employee.IsTemporaryEmployee() && !userIsDatabase {
			fullName := fmt.Sprintf("%v %v", utils.CleanString(currentDependent.GetFirstName()), utils.CleanString(currentDependent.GetLastName()))
			res := dependentsView.FindDependentByRegistrationNumberAndName(employee.RegistrationNumberToInt64(), fullName)

			if (res == dependentModel.DependentView{}) {
				currentDependent.SetActive(false)
				res, err := dependentRepository.UpdateDependent(&currentDependent)
				if err != nil || res != 1 {
					saveErrorLog(fmt.Sprintf("Erro ao tentar desativar os dados do dependente %v", currentDependent.GetFirstName()), err, employee.GetRegistrationNumber())
				}
			}
		}
	}
}

func updateDependents(dependents *dependentModel.Dependents, employees *employeeModel.Employees, dependentsView *dependentModel.DependentsView) {
	if settings.GoRoutines.UseGoRoutines {
		var wg sync.WaitGroup
		dependentsChannel := make(chan dependentModel.DependentView)

		wg.Add(settings.GoRoutines.GoRoutinesCount)

		for i := 1; i <= settings.GoRoutines.GoRoutinesCount; i++ {
			go func() {
				defer wg.Done()
				for currentDependent := range dependentsChannel {
					executeDependentUpdateProcess(currentDependent, employees, dependents)
				}
			}()
		}

		for j := range *dependentsView {
			dependentsChannel <- (*dependentsView)[j]
		}

		close(dependentsChannel)
		wg.Wait()
	} else {
		for i := range *dependentsView {
			currentDependent := (*dependentsView)[i]
			executeDependentUpdateProcess(currentDependent, employees, dependents)
		}
	}
}

func executeDependentUpdateProcess(currentDependent dependentModel.DependentView, employees *employeeModel.Employees, dependents *dependentModel.Dependents) {
	employee := employees.FindEmployeeByRegistrationNumber(currentDependent.RegistrationNumberToString())

	if (employee != employeeModel.Employee{}) {
		fullName := fmt.Sprintf("%v %v", utils.CleanString(currentDependent.GetPrimeiroNome()), utils.CleanString(currentDependent.GetSobrenome()))

		dependent := dependents.FindDependentByEmployeeIdAndName(employee.GetEmployeeId(), fullName)

		if (dependent != dependentModel.Dependent{}) {
			updateDependent(&employee, &dependent, &currentDependent)
		} else {
			createDependent(&employee, &currentDependent)
		}
	}
}

func updateDependent(employee *employeeModel.Employee, dependent *dependentModel.Dependent, dependentView *dependentModel.DependentView) {
	kindship := dependentView.GetViewDataKindship()

	if dependent.GetKinship() != kindship || !dependent.Equals(*dependentView) {
		dependent.SetActive(dependentView.GetCdiSituacaoDependente() == 1)
		dependent.SetFirstName(utils.CleanString(dependentView.GetPrimeiroNome()))
		dependent.SetLastName(utils.CleanString(dependentView.GetSobrenome()))
		dependent.DateOfBirth = dependentView.DataNascimento
		dependent.Gender = dependentView.Sexo
		dependent.SetKinship(kindship)
		dependent.Cpf = dependentView.Cpf
		dependent.SetModifiedDate(time.Now())
		dependent.SetModifiedUser(settings.ModifiedUser)

		res, err := dependentRepository.UpdateDependent(dependent)

		if err != nil || res != 1 {
			saveErrorLog(fmt.Sprintf("Erro ao tentar atualizar os dados do dependente %v", dependentView.GetNomeCompleto()), err, employee.GetRegistrationNumber())
		}
	}
}

func createDependent(employee *employeeModel.Employee, dependentView *dependentModel.DependentView) {
	if dependentView.NameIsValid() {
		now := time.Now()

		dependent := dependentModel.Dependent{
			EmployeeId:  employee.EmployeeId,
			Gender:      dependentView.Sexo,
			DateOfBirth: dependentView.DataNascimento,
			Cpf:         dependentView.Cpf,
		}

		dependent.SetFirstName(utils.CleanString(dependentView.GetPrimeiroNome()))
		dependent.SetLastName(utils.CleanString(dependentView.GetSobrenome()))
		dependent.SetKinship(dependentView.GetViewDataKindship())
		dependent.SetCreatedDate(now)
		dependent.SetModifiedDate(now)
		dependent.SetModifiedUser(settings.ModifiedUser)
		dependent.SetActive(true)

		res, err := dependentRepository.CreateDependent(&dependent)

		if err != nil || res != 1 {
			saveErrorLog(fmt.Sprintf("Erro ao tentar inserir o dependente %v", dependentView.GetNomeCompleto()), err, employee.GetRegistrationNumber())
		}
	}
}
