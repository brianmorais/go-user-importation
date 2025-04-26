package importation

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/brianmorais/go-user-importation/domain/enums/user_access_type"
	"github.com/brianmorais/go-user-importation/domain/enums/user_type"
	benefitModel "github.com/brianmorais/go-user-importation/domain/models/benefit"
	employeeModel "github.com/brianmorais/go-user-importation/domain/models/employee"
	userModel "github.com/brianmorais/go-user-importation/domain/models/user"
	"github.com/brianmorais/go-user-importation/domain/utils"
)

var employeeRoles *employeeModel.EmployeeRoles

func readEmployeeViewData() {
	saveMessageLog(fmt.Sprintf("Início da importação de employees - %v", *importationNumber))

	var err error
	var employeesView *employeeModel.EmployeesView
	var employees *employeeModel.Employees

	if employeesView, err = employeeRepository.GetEmployeesView(); err != nil {
		saveErrorLog("Erro ao obter employees da view", err, "")
		return
	}

	if employees, err = employeeRepository.GetEmployees(); err != nil {
		saveErrorLog("Erro ao obter employees do banco", err, "")
		return
	}

	if employeeRoles, err = employeeRepository.GetEmployeeRoles(); err != nil {
		saveErrorLog("Erro ao obter EmployeeRoles", err, "")
		return
	}

	deactivateEmployees(employees, employeesView)
	updateEmployees(employees, employeesView)
	saveMessageLog(fmt.Sprintf("Fim da importação de employees - %v", *importationNumber))
}

func deactivateEmployees(employees *employeeModel.Employees, employeesView *employeeModel.EmployeesView) {
	employeeRegistrationNumbers := employees.GetNotExceptionAndActiveEmployeeRegistrationNumbers()
	employeeViewRegistrationNumbers := employeesView.GetEmployeeViewRegistrationNumbers()
	deactivateListResult := employeeModel.GetExceptRegistrationNumber(employeeRegistrationNumbers, employeeViewRegistrationNumbers)
	employeesToDeactivate := employees.GetEmployeesToDeactivate(deactivateListResult)

	if settings.GoRoutines.UseGoRoutines {
		var wg sync.WaitGroup

		employeesChannel := make(chan employeeModel.Employee)

		wg.Add(settings.GoRoutines.GoRoutinesCount)

		for i := 1; i <= settings.GoRoutines.GoRoutinesCount; i++ {
			go func() {
				defer wg.Done()
				for currentEmployee := range employeesChannel {
					executeEmployeeDeactivateProcess(currentEmployee)
				}
			}()
		}

		for j := range *employeesToDeactivate {
			employeesChannel <- (*employeesToDeactivate)[j]
		}

		close(employeesChannel)
		wg.Wait()
	} else {
		for i := range *employeesToDeactivate {
			currentEmployee := (*employeesToDeactivate)[i]
			executeEmployeeDeactivateProcess(currentEmployee)
		}
	}
}

func executeEmployeeDeactivateProcess(currentEmployee employeeModel.Employee) {
	currentEmployee.SetActive(false)
	res, err := employeeRepository.UpdateEmployee(&currentEmployee)

	if err != nil || res != 1 {
		saveErrorLog(fmt.Sprintf("Erro ao tentar desativar o employee %v %v", currentEmployee.GetFirstName(), currentEmployee.GetLastName()), err, currentEmployee.GetRegistrationNumber())
	}
}

func updateEmployees(employees *employeeModel.Employees, employeesView *employeeModel.EmployeesView) {
	if settings.GoRoutines.UseGoRoutines {
		var wg sync.WaitGroup
		employeesChannel := make(chan employeeModel.EmployeeView)

		wg.Add(settings.GoRoutines.GoRoutinesCount)

		for i := 1; i <= settings.GoRoutines.GoRoutinesCount; i++ {
			go func() {
				defer wg.Done()
				for currentEmployee := range employeesChannel {
					executeEmployeeUpdateProcess(currentEmployee, employees)
				}
			}()
		}

		for j := range *employeesView {
			employeesChannel <- (*employeesView)[j]
		}

		close(employeesChannel)
		wg.Wait()
	} else {
		for i := range *employeesView {
			currentEmployee := (*employeesView)[i]
			executeEmployeeUpdateProcess(currentEmployee, employees)
		}
	}
}

func executeEmployeeUpdateProcess(currentEmployee employeeModel.EmployeeView, employees *employeeModel.Employees) {
	currentEmployee.SetPrimeiroNome(utils.CleanString(currentEmployee.GetPrimeiroNome()))
	currentEmployee.SetSobrenome(utils.CleanString(currentEmployee.GetSobrenome()))

	employeeToUpdate := employees.FindEmployeeByRegistrationNumber(currentEmployee.RegistrationNumberToString())

	if (employeeToUpdate != employeeModel.Employee{}) {
		updateEmployee(&employeeToUpdate, &currentEmployee)
	} else {
		createEmployee(&currentEmployee)
	}
}

func updateEmployee(employee *employeeModel.Employee, employeeView *employeeModel.EmployeeView) {
	roleId, err := matchRole(*employeeView)

	if err != nil {
		saveErrorLog(fmt.Sprintf("Erro ao tentar obter o roleId do employee %v %v", employeeView.GetPrimeiroNome(), employeeView.GetSobrenome()), err, employee.GetRegistrationNumber())
		return
	}

	now := time.Now()
	gender := utils.CleanString(employeeView.GetSexoSigla())
	email := fmt.Sprintf("%v%v", strings.ToLower(employeeView.GetUsuario()), settings.EmailDomain)
	changeRole := employee.GetRoleId() != roleId
	active := employeeView.GetSituacao() == "A" && (!employeeView.DataRescisao.Valid || employeeView.DataRescisao == sql.NullTime{})

	if strings.ToLower(employee.GetEmail()) != email || employee.GetRoleId() != roleId || !employee.Equals(*employeeView) {
		employee.SetActive(active)
		employee.AdmissionDate = employeeView.DataAdmissao
		employee.DateOfBirth = employeeView.DataNascimento
		employee.SetEmail(email)
		employee.SetFirstName(utils.CleanString(employeeView.GetPrimeiroNome()))
		employee.SetLastName(utils.CleanString(employeeView.GetSobrenome()))
		employee.SetGender(gender)
		employee.Location = employeeView.Estado
		employee.SetModifiedDate(now)
		employee.SetModifiedUser(settings.ModifiedUser)
		employee.TerminationDate = employeeView.DataRescisao
		employee.Cpf = employeeView.Cpf

		upgrade, err := isUpgradeRole(*employee, roleId)

		if err != nil {
			saveErrorLog("Erro ao obter hierarquia do employee", err, employee.GetRegistrationNumber())
		}

		if changeRole && upgrade {
			employee.SetRoleId(roleId)
			ajustEmployeeBenefitByRole(employee)
		}

		res, err := employeeRepository.UpdateEmployee(employee)

		if err != nil || res != 1 {
			saveErrorLog(fmt.Sprintf("Erro ao tentar atualizar os dados do employee %v %v", employeeView.GetPrimeiroNome(), employeeView.GetSobrenome()), err, employee.GetRegistrationNumber())
			return
		}
	}

	res, err := hasBenefitEmployee(employee)

	if err != nil {
		saveErrorLog(fmt.Sprintf("Erro ao verificar se o employee %v %v possui benefícios", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
	}

	if !res && err == nil {
		linkCycleToNewEmployee(employee)
	}

	verifyEmployeeUserForUpdate(employee, utils.CleanString(employeeView.GetUsuario()))
	updateExceptionUserPerExceptionDate(employee)
	updateEmployeeStatusPerExceptionUserStatus(employee)
}

func createEmployee(employeeView *employeeModel.EmployeeView) {
	roleId, err := matchRole(*employeeView)

	if err != nil {
		saveErrorLog(fmt.Sprintf("Erro ao tentar obter o roleId do employee %v %v", employeeView.GetPrimeiroNome(), employeeView.GetSobrenome()), err, employeeView.RegistrationNumberToString())
		return
	}

	now := time.Now()
	email := fmt.Sprintf("%v%v", strings.ToLower(employeeView.GetUsuario()), settings.EmailDomain)
	active := employeeView.GetSituacao() == "A" && (!employeeView.DataRescisao.Valid || employeeView.DataRescisao == sql.NullTime{})
	firstName := utils.CleanString(employeeView.GetPrimeiroNome())
	lastName := utils.CleanString(employeeView.GetSobrenome())
	gender := utils.CleanString(strings.ToUpper(employeeView.GetSexoSigla()))

	employee := employeeModel.Employee{
		AdmissionDate:   employeeView.DataAdmissao,
		DateOfBirth:     employeeView.DataNascimento,
		Location:        employeeView.Estado,
		TerminationDate: employeeView.DataRescisao,
		Cpf:             employeeView.Cep,
	}

	employee.SetActive(active)
	employee.SetCreatedDate(now)
	employee.SetEmail(email)
	employee.SetEmployeeTypeId(1)
	employee.SetFirstName(firstName)
	employee.SetLastName(lastName)
	employee.SetGender(gender)
	employee.SetModifiedDate(now)
	employee.SetModifiedUser(settings.ModifiedUser)
	employee.SetRegistrationNumber(employeeView.RegistrationNumberToString())
	employee.SetRoleId(roleId)

	res, err := employeeRepository.CreateEmployee(&employee)

	if err != nil || res != 1 {
		saveErrorLog(fmt.Sprintf("Erro ao tentar criar o employee %v %v", employeeView.GetPrimeiroNome(), employeeView.GetSobrenome()), err, employee.GetRegistrationNumber())
		return
	}

	insertedEmployee, err := employeeRepository.GetEmployeeByRegistrationNumber(employeeView.RegistrationNumberToString())

	if err != nil || !insertedEmployee.IsValid() {
		saveErrorLog(fmt.Sprintf("Erro ao tentar obter o employee %v %v", employeeView.GetPrimeiroNome(), employeeView.GetSobrenome()), err, employee.GetRegistrationNumber())
		return
	}

	createNewUser(*insertedEmployee, employeeView.GetUsuario())
	linkCycleToNewEmployee(insertedEmployee)
}

func verifyEmployeeUserForUpdate(employee *employeeModel.Employee, user string) {
	users, err := userRepository.FindUsersByEmployeeId(employee.GetEmployeeId())

	if err != nil {
		saveErrorLog(fmt.Sprintf("Erro ao tentar obter os users do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		return
	}

	if len(users) == 0 {
		createNewUser(*employee, user)
		currentCycle, err := benefitRepository.GetActiveCycle()

		if err != nil {
			saveErrorLog(fmt.Sprintf("Erro ao tentar obter o ciclo atual do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		}

		if currentCycle.IsValid() {
			employeeBenefit, err := benefitRepository.GetBenefitByEmployeIdAndCycleId(employee.GetEmployeeId(), currentCycle.GetCycleId())

			if err != nil {
				saveErrorLog(fmt.Sprintf("Erro ao obter os benefícios do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			}

			if employeeBenefit.IsValid() {
				linkCycleToNewEmployee(employee)
			}
		}
	} else if len(users) > 1 {
		userAd := users.FindActiveDirectoryUser()

		if (userAd != userModel.User{}) {
			userAd.SetLocked(!employee.GetActive())
			userAd.SetModifiedDate(time.Now())
			userAd.SetModifiedUser(settings.ModifiedUser)

			res, err := userRepository.UpdateUser(userAd)

			if err != nil || res != 1 {
				saveErrorLog(fmt.Sprintf("Erro ao atualizar o user do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			}
		}

		userDb := users.FindDatabaseUser()

		if (userDb != userModel.User{}) {
			userDb.SetLocked(employee.BlockUserException())
			userDb.SetModifiedDate(time.Now())
			userDb.SetModifiedUser(settings.ModifiedUser)

			res, err := userRepository.UpdateUser(userDb)

			if err != nil || res != 1 {
				saveErrorLog(fmt.Sprintf("Erro ao atualizar o user do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			}
		}

		if employee.IsUnlimitedDateException() && employee.GetActive() {
			employee.SetModifiedDate(time.Now())
			employee.SetModifiedUser(settings.ModifiedUser)
			employee.SetExceptionEndDate(time.Now())

			res, err := employeeRepository.UpdateEmployee(employee)

			if err != nil || res != 1 {
				saveErrorLog(fmt.Sprintf("Erro ao atualizar o user do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			}
		}
	} else if len(users) == 1 {
		userAdUpdate := users[0]

		userAdUpdate.SetLocked(!employee.GetActive())
		userAdUpdate.SetModifiedDate(time.Now())
		userAdUpdate.SetModifiedUser(settings.ModifiedUser)

		res, err := userRepository.UpdateUser(userAdUpdate)

		if err != nil || res != 1 {
			saveErrorLog(fmt.Sprintf("Erro ao atualizar o user do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		}
	}

	userAd := users.FindActiveDirectoryUser()

	if userAd.IsValid() && utils.CleanString(userAd.GetUserId()) != utils.CleanString(user) {
		res, err := userRepository.DeleteUser(userAd.GetUserId())

		if res != 1 || err != nil {
			saveErrorLog(fmt.Sprintf("Erro ao deletar o user do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			return
		}

		userAd.SetUserId(user)
		userAd.SetModifiedDate(time.Now())
		userAd.SetModifiedUser(settings.ModifiedUser)

		userRepository.CreateUser(userAd)
	}
}

func updateEmployeeStatusPerExceptionUserStatus(employee *employeeModel.Employee) {
	users, err := userRepository.FindUsersByEmployeeId(employee.GetEmployeeId())

	if err != nil {
		saveErrorLog(fmt.Sprintf("Erro ao tentar obter os users do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		return
	}

	if len(users) > 0 {
		databaseUser := users.FindDatabaseUser()

		if !employee.GetActive() && databaseUser.IsValid() && !databaseUser.GetLocked() && employee.GetExceptionEndDate().After(time.Now()) {
			employee.SetActive(true)
			employee.SetModifiedDate(time.Now())
			employee.SetModifiedUser(settings.ModifiedUser)

			res, err := employeeRepository.UpdateEmployee(employee)

			if res != 1 || err != nil {
				saveErrorLog(fmt.Sprintf("Erro ao tentar atualizar o employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			}
		}
	}
}

func updateExceptionUserPerExceptionDate(employee *employeeModel.Employee) {
	users, err := userRepository.FindUsersByEmployeeId(employee.GetEmployeeId())

	if err != nil {
		saveErrorLog(fmt.Sprintf("Erro ao tentar obter os users do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		return
	}

	if len(users) > 0 {
		databaseUser := users.FindDatabaseUser()

		if databaseUser.IsValid() && databaseUser.GetLocked() && time.Now().Before(employee.GetExceptionEndDate()) && !employee.GetActive() {
			databaseUser.SetLocked(false)
			databaseUser.SetModifiedDate(time.Now())
			databaseUser.SetModifiedUser(settings.ModifiedUser)
			databaseUser.SetAccessTypeId(user_access_type.Employee)

			res, err := userRepository.UpdateUser(databaseUser)

			if err != nil || res != 1 {
				saveErrorLog(fmt.Sprintf("Erro ao tentar atualizar o user do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			}
		}
	}
}

func hasBenefitEmployee(employee *employeeModel.Employee) (bool, error) {
	currentCycle, err := benefitRepository.GetActiveCycle()

	if err != nil {
		return false, err
	}

	employeeBenefit, err := benefitRepository.GetBenefitByEmployeIdAndCycleId(employee.GetEmployeeId(), currentCycle.GetCycleId())

	if err != nil {
		return false, err
	}

	if employeeBenefit.IsValid() {
		return true, nil
	}

	return false, nil
}

func ajustEmployeeBenefitByRole(employee *employeeModel.Employee) {
	currentCycle, err := benefitRepository.GetActiveCycle()

	if err != nil {
		saveErrorLog("Erro ao obter o cycle ativo", err, employee.GetRegistrationNumber())
		return
	}

	if currentCycle.IsValid() && employee.IsValid() {
		employeeBenefit, err := benefitRepository.GetBenefitByEmployeIdAndCycleId(employee.GetEmployeeId(), currentCycle.GetCycleId())

		if err != nil {
			saveErrorLog(fmt.Sprintf("Erro ao obter os benefícios do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			return
		}

		benefitCycle, err := benefitModel.DeserializeBenefit(currentCycle.GetBenefitSetting())

		if err != nil {
			saveErrorLog(fmt.Sprintf("Erro ao descerializar json de benefícios do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			return
		}

		if len(benefitCycle) > 0 {
			benefits := benefitCycle.GetBenefitsByRoleId(int64(employee.GetRoleId()))
			jsonBenefits, err := benefits.SerializeBenefit()

			if err != nil {
				saveErrorLog(fmt.Sprintf("Erro ao serializar json de benefícios do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
				return
			}

			employeeBenefit.SetBenefit(jsonBenefits)
			employeeBenefit.SetRoleId(employee.GetRoleId())
			employeeBenefit.SetModifiedDate(time.Now())
			employeeBenefit.SetModifiedUser(settings.ModifiedUser)

			res, err := benefitRepository.UpdateEmployeeBenefit(employeeBenefit)

			if err != nil || res != 1 {
				saveErrorLog(fmt.Sprintf("Erro ao salvar benefícios do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
			}
		}
	}
}

func isUpgradeRole(employee employeeModel.Employee, roleIdView int32) (bool, error) {
	roles, err := employeeRepository.GetEmployeeRoles()

	if err != nil {
		return false, err
	}

	employeeHierarchy := roles.FindHierarchByRoleId(employee.GetRoleId())
	viewEmployeeHierarchy := roles.FindHierarchByRoleId(roleIdView)

	if employeeHierarchy != 0 && viewEmployeeHierarchy != 0 && viewEmployeeHierarchy > employeeHierarchy {
		return true, nil
	}

	return false, nil
}

func matchRole(employee employeeModel.EmployeeView) (int32, error) {
	gpoFuncional := strings.ToLower(employee.GetGpoFuncional())
	var res employeeModel.EmployeeRole

	switch gpoFuncional {
	case "outros":
		res = employeeRoles.FindRoleByName("Tripulante")
	case "presidente", "vice-presidente":
		res = employeeRoles.FindRoleByName("Diretoria")
	default:
		res = employeeRoles.FindRoleByName(gpoFuncional)
	}

	if (res == employeeModel.EmployeeRole{}) {
		return 0, errors.New("grupo funcional não localizado")
	}

	return int32(res.GetRoleId()), nil
}

func isUserDatabaseType(employee employeeModel.Employee) (bool, error) {
	users, err := userRepository.GetUsersByTypeAndEmployeeId(employee.GetEmployeeId(), user_type.Database)

	if err != nil {
		return false, err
	}

	if len(users) > 0 {
		return true, nil
	}

	return false, nil
}

func createNewUser(employee employeeModel.Employee, user string) {
	userData, err := userRepository.FindUserById(user)

	if err != nil {
		saveErrorLog(fmt.Sprintf("Erro ao tentar localizar o usuário do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		return
	}

	if !userData.IsValid() {
		userData = userModel.User{
			EmployeeId: employee.EmployeeId,
		}

		now := time.Now()
		userData.SetCreatedDate(now)
		userData.SetLocked(false)
		userData.SetModifiedDate(now)
		userData.SetModifiedUser(settings.ModifiedUser)
		userData.SetPasswordModifiedDate(now)
		userData.SetPasswordReset(false)
		userData.SetUserId(strings.ToLower(user))
		userData.SetUserTypeId(user_type.ActiveDirectory)
		userData.SetAccessTypeId(user_access_type.Employee)
		userData.SetPassword("")

		res, err := userRepository.CreateUser(userData)

		if err != nil || res != 1 {
			saveErrorLog(fmt.Sprintf("Erro ao tentar criar o usuário do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		}

	} else if userData.EmployeeId != employee.EmployeeId {
		userData.EmployeeId = employee.EmployeeId
		userData.SetModifiedDate(time.Now())
		userData.SetModifiedUser(settings.ModifiedUser)
		res, err := userRepository.UpdateUser(userData)

		if err != nil || res != 1 {
			saveErrorLog(fmt.Sprintf("Erro ao tentar atualizar o usuário do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		}
	}
}

func linkCycleToNewEmployee(employee *employeeModel.Employee) {
	activeCycle, err := benefitRepository.GetActiveCycle()

	if err != nil {
		saveErrorLog(fmt.Sprintf("Erro ao obter o ciclo para criar o benefício do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
		return
	}

	now := time.Now()

	benefit := benefitModel.EmployeeBenefit{
		Benefit:    activeCycle.BenefitSetting,
		CycleId:    activeCycle.CycleId,
		RoleId:     employee.RoleId,
		EmployeeId: employee.EmployeeId,
	}
	benefit.SetCreatedDate(now)
	benefit.SetModifiedDate(now)
	benefit.SetModifiedUser(settings.ModifiedUser)

	res, err := benefitRepository.CreateEmployeeBenefit(benefit)

	if err != nil || res != 1 {
		saveErrorLog(fmt.Sprintf("Erro ao salvar o benefício do employee %v %v", employee.GetFirstName(), employee.GetLastName()), err, employee.GetRegistrationNumber())
	}
}
